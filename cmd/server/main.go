package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/victorabarros/termgifforge/internal/files"
	"github.com/victorabarros/termgifforge/internal/gif"
	"github.com/victorabarros/termgifforge/internal/id"
	"github.com/victorabarros/termgifforge/pkg/models"
)

var (
	port     = "80"
	version  = "0.1.3"
	homePage = "https://victor.barros.engineer/termgif"

	outputCmdFormat = "Output %s"
	setCmds         = []string{
		"Set WindowBar Colorful",
		"Set FontSize 12",
		"Set Width 800",
		"Set Height 400",
	}

	// GIFDetails is a map of GIFs and their statuses
	details = models.NewGIFDetails()
)

func init() {
	// create output directory if it doesn't exist; where GIFs are stored
	if err := files.CreateOutputDirectory(); err != nil {
		os.Exit(1)
	}

	// TODO move to separate function
	// load details mapper
	func() {
		gifs, err := files.ListGIFs()
		if err != nil {
			os.Exit(1)
		}
		for _, gif := range gifs {
			name := gif.Name()
			// remove .gif from name
			id := name[:len(name)-4]
			details.SetStatus(id, models.GIFStatuses.Ready)
		}
	}()

	// creating error and invalid GIFs if they don't exist
	if d, _ := details.Get("error"); d.Status != models.GIFStatuses.Ready {
		errorGIF()
	}
	if d, _ := details.Get("invalid"); d.Status != models.GIFStatuses.Ready {
		invalidGIF()
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}); err != nil {
		log.Fatalf("Sentry initialization failed: %v\n", err)
	}
}

func main() {
	r := gin.Default()
	r.Use(sentrygin.New(sentrygin.Options{}))

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, homePage)
	})

	r.GET("/ping", func(c *gin.Context) {
		gifs, err := files.ListGIFs()
		if err != nil {
			log.Printf("Fail to list GIFs %+2v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "GIF in process"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success", "nofGIFs": len(gifs), "version": version})
	})

	rpcGroup := r.Group("/api/v1")
	rpcGroup.GET("/gif", createGIFHandler)
	rpcGroup.GET("/mock", func(c *gin.Context) {
		c.File("output/error.gif")
	})
	rpcGroup.DELETE("/internal-use/gif/:id", func(c *gin.Context) {
		//TODO user files.EraseGIF
	})

	go files.Cleaner(&details)

	log.Printf("Starting app version %s in port %s", version, port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Printf("%+2v/n", err)
	}
}

func createGIFHandler(c *gin.Context) {
	extras := map[string]interface{}{
		"accept":          c.GetHeader("Accept"),
		"acceptEncoding":  c.GetHeader("Accept-Encoding"),
		"acceptLanguage":  c.GetHeader("Accept-Language"),
		"clientIP":        c.ClientIP(),
		"connection":      c.GetHeader("Connection"),
		"host":            c.GetHeader("Host"),
		"origin":          c.GetHeader("Origin"),
		"referer":         c.GetHeader("Referer"),
		"userAgent":       c.GetHeader("User-Agent"),
		"xForwardedFor":   c.GetHeader("X-Forwarded-For"),
		"xForwardedProto": c.GetHeader("X-Forwarded-Proto"),
		"xRealIP":         c.GetHeader("X-Real-IP"),
	}

	sentry.CaptureEvent(&sentry.Event{
		User: sentry.User{
			IPAddress: c.ClientIP(),
		},
		Environment: os.Getenv("ENVIRONMENT"),
		Extra:       extras,
		Level:       sentry.LevelDebug,
		Message:     "createGIFHandler",
		Platform:    c.GetHeader("User-Agent"),
	})

	cmdsInputStr := c.Query("commands")
	cmdInput := []string{}
	if err := json.Unmarshal([]byte(cmdsInputStr), &cmdInput); err != nil {
		log.Printf("Error trying to serialize object: %+2v\n", err)
		c.File("output/invalid.gif")
		return
	}

	id := id.NewUUUIDAsString(cmdsInputStr)
	outGifPath := fmt.Sprintf("output/%s.gif", id)
	if d, ok := details.Get(id); ok {
		if d.Status == models.GIFStatuses.Fail {
			c.File("output/error.gif")
			return
		}
		if d.Status == models.GIFStatuses.Processing {
			c.JSON(http.StatusAccepted, gin.H{"message": "GIF in process"})
			return
		}
		if d.Status == models.GIFStatuses.Ready {
			details.SetLastAccess(id, time.Now())
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
	details.SetStatus(id, models.GIFStatuses.Processing)

	if err := gif.WriteTape(cmds, outTapePath); err != nil {
		log.Printf("Error writing to file: %+2v\n", err)
		details.SetStatus(id, models.GIFStatuses.Fail)
		return err
	}
	defer os.Remove(outTapePath)

	if err := gif.ExecVHS(outTapePath); err != nil {
		log.Printf("Error running command: %+2v\n", err)
		details.SetStatus(id, models.GIFStatuses.Fail)
		return err
	}

	details.SetStatus(id, models.GIFStatuses.Ready)

	log.Printf("GIF Created id %s\n", id)
	return nil
}

func errorGIF() error {
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

func invalidGIF() error {
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
