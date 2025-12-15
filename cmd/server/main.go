package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrlogrus"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/sirupsen/logrus"
	"github.com/victorabarros/termgifforge/internal/files"
	"github.com/victorabarros/termgifforge/internal/gif"
	"github.com/victorabarros/termgifforge/internal/id"
	"github.com/victorabarros/termgifforge/internal/logs"
	"github.com/victorabarros/termgifforge/pkg/models"
)

var (
	port     = "80"    // TODO move to env
	version  = "0.2.1" // TODO move to env
	homePage = "https://victor.barros.engineer/termgif"
	appName  = "termgifforge"

	outputCmdFormat = "Output %s"
	setCmds         = []string{
		"Set WindowBar Colorful",
		"Set FontSize 12",
		"Set Width 800",
		"Set Height 400",
	}

	newRelicApp *newrelic.Application
	logLevel    = logrus.InfoLevel // TODO move to env

	// GIFDetails is a map of GIFs and their statuses
	cache = models.NewGIFDetails()
)

func init() {
	// create output directory if it doesn't exist; where GIFs are stored
	if err := files.CreateOutputDirectory(); err != nil {
		os.Exit(1)
	}

	// TODO move to separate function
	// load cache mapper
	func() {
		gifs, err := files.ListGIFs()
		if err != nil {
			os.Exit(1)
		}
		for _, gif := range gifs {
			name := gif.Name()
			// remove .gif from name
			id := name[:len(name)-4]
			cache.SetStatus(id, models.GIFStatuses.Ready)
		}
	}()

	// creating error and invalid GIFs if they don't exist
	if d, _ := cache.Get("error"); d.Status != models.GIFStatuses.Ready {
		createErrorGIF()
	}
	if d, _ := cache.Get("invalid"); d.Status != models.GIFStatuses.Ready {
		createInvalidGIF()
	}

	newRelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(fmt.Sprintf("%s-%s", appName, os.Getenv("ENVIRONMENT"))),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		logs.Log.Fatalf("NewRelic initialization failed: %+2v", err)
	}

	logs.InitLog(logLevel, nrlogrus.NewFormatter(newRelicApp, &logrus.TextFormatter{}))
}

func main() {
	r := gin.Default()
	r.Use(nrgin.Middleware(newRelicApp))

	err := newRelicApp.WaitForConnection(10 * time.Second)
	if nil != err {
		logs.Log.Fatalf("Failed to connect application %+2v", err)
	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, homePage)
	})

	r.GET("/ping", func(c *gin.Context) {
		gifs, err := files.ListGIFs()
		if err != nil {
			logs.Log.Errorf("Fail to list GIFs %+2v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "GIF in process"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success", "nofGIFs": len(gifs), "version": version})
	})

	rpcGroup := r.Group("/api/v1")
	rpcGroup.GET("/gif", createGIFHandler)
	rpcGroup.GET("/gif/:id", getGIFHandler)
	rpcGroup.GET("/mock", func(c *gin.Context) {
		c.File("output/error.gif")
	})
	rpcGroup.DELETE("/internal-use/gif/:id", func(c *gin.Context) {
		//TODO user files.EraseGIF
	})

	go files.Cleaner(&cache)

	logs.Log.Infof("Starting app version %s in", version)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		logs.Log.Fatalf("Failed to start server: %+2v", err)
	}
}

func getGIFHandler(c *gin.Context) {
	extras := logrus.Fields{
		"accept":          c.GetHeader("Accept"),
		"acceptEncoding":  c.GetHeader("Accept-Encoding"),
		"acceptLanguage":  c.GetHeader("Accept-Language"),
		"clientIP":        c.ClientIP(),
		"connection":      c.GetHeader("Connection"),
		"environment":     os.Getenv("ENVIRONMENT"),
		"host":            c.GetHeader("Host"),
		"origin":          c.GetHeader("Origin"),
		"referer":         c.GetHeader("Referer"),
		"userAgent":       c.GetHeader("User-Agent"),
		"xForwardedFor":   c.GetHeader("X-Forwarded-For"),
		"xForwardedProto": c.GetHeader("X-Forwarded-Proto"),
		"xRealIP":         c.GetHeader("X-Real-IP"),
	}

	logs.Log.WithFields(extras).Info("createGIFHandler")

	id := c.Param("id")
	logs.Log.WithFields(extras).Infof("id: %s", id)

	outGifPath := fmt.Sprintf("output/%s.gif", id)
	if d, ok := cache.Get(id); ok {
		if d.Status == models.GIFStatuses.Fail {
			c.File("output/error.gif")
			return
		}
		if d.Status == models.GIFStatuses.Processing {
			c.JSON(http.StatusAccepted, gin.H{"message": "GIF in process"})
			return
		}
		if d.Status == models.GIFStatuses.Ready {
			cache.SetLastAccess(id, time.Now())
			c.File(outGifPath)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "GIF not found"})
}

func createGIFHandler(c *gin.Context) {
	extras := logrus.Fields{
		"accept":          c.GetHeader("Accept"),
		"acceptEncoding":  c.GetHeader("Accept-Encoding"),
		"acceptLanguage":  c.GetHeader("Accept-Language"),
		"clientIP":        c.ClientIP(),
		"connection":      c.GetHeader("Connection"),
		"environment":     os.Getenv("ENVIRONMENT"),
		"host":            c.GetHeader("Host"),
		"origin":          c.GetHeader("Origin"),
		"referer":         c.GetHeader("Referer"),
		"userAgent":       c.GetHeader("User-Agent"),
		"xForwardedFor":   c.GetHeader("X-Forwarded-For"),
		"xForwardedProto": c.GetHeader("X-Forwarded-Proto"),
		"xRealIP":         c.GetHeader("X-Real-IP"),
	}

	logs.Log.WithFields(extras).Info("createGIFHandler")

	cmdsInputStr := c.Query("commands")
	cmdInput := []string{}
	if err := json.Unmarshal([]byte(cmdsInputStr), &cmdInput); err != nil {
		logs.Log.Errorf("Error trying to serialize object: %+2v", err)
		c.File("output/invalid.gif")
		return
	}

	id := id.NewUUUIDAsString(cmdsInputStr)
	outGifPath := fmt.Sprintf("output/%s.gif", id)
	if d, ok := cache.Get(id); ok {
		if d.Status == models.GIFStatuses.Fail {
			c.File("output/error.gif")
			return
		}
		if d.Status == models.GIFStatuses.Processing {
			c.JSON(http.StatusAccepted, gin.H{"message": "GIF in process"})
			return
		}
		if d.Status == models.GIFStatuses.Ready {
			cache.SetLastAccess(id, time.Now())
			c.File(outGifPath)
			return
		}
	}

	cmds := append([]string{fmt.Sprintf(outputCmdFormat, outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(id, cmds)

	c.JSON(http.StatusAccepted, gin.H{"message": "GIF in process"})
}

func processGIF(id string, cmds []string) error {
	outTapePath := fmt.Sprintf("output/%s.tape", id)
	cache.SetStatus(id, models.GIFStatuses.Processing)

	if err := gif.WriteTape(cmds, outTapePath); err != nil {
		logs.Log.Errorf("Error writing to file: %+2v", err)
		cache.SetStatus(id, models.GIFStatuses.Fail)
		return err
	}
	defer os.Remove(outTapePath)

	if err := gif.ExecVHS(outTapePath); err != nil {
		logs.Log.Errorf("Error running command: %+2v", err)
		cache.SetStatus(id, models.GIFStatuses.Fail)
		return err
	}

	cache.SetStatus(id, models.GIFStatuses.Ready)

	logs.Log.Infof("GIF Created id %s", id)
	return nil
}

func createErrorGIF() error {
	cmdInput := []string{
		"Type \"Sorry, it was not possible create your GIF. =/\"",
		"Sleep 6s",
	}

	id := "error"
	outGifPath := fmt.Sprintf("output/%s.gif", id)

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(id, cmds)

	return nil
}

func createInvalidGIF() error {
	cmdInput := []string{
		"Type \"Invalid request...\"",
		"Sleep 6s",
	}

	id := "invalid"
	outGifPath := fmt.Sprintf("output/%s.gif", id)

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(id, cmds)

	return nil
}
