import csv
import ipaddress
import json
import os
import sys
import time
import urllib.error
import urllib.parse
import urllib.request

def lookup(ip):
    token = os.getenv("IPINFO_TOKEN")
    url = f"https://ipinfo.io/{ip}/json"
    if token:
        url = f"{url}?{urllib.parse.urlencode({'token': token})}"

    req = urllib.request.Request(url, headers={"User-Agent": "termgifforge-iplookup/1.0"})
    with urllib.request.urlopen(req, timeout=10) as resp:
        return json.load(resp)

def get_ip_from_row(row: dict) -> str:
    # table.csv currently uses "clientIp" (lowercase p), but accept other variants too.
    return (row.get("clientIp") or row.get("clientIP") or row.get("client_ip") or "").strip()


cache: dict[str, dict] = {}

with open("table.csv", newline="") as f:
    reader = csv.DictReader(f)

    out = csv.writer(sys.stdout)
    out.writerow(
        [
            "clientIp",
            "Logs",
            "Timestamp",
            "country",
            "region",
            "city",
            "org",
            "loc",
            "note",
        ]
    )

    for row in reader:
        ip = get_ip_from_row(row)
        if not ip:
            continue

        try:
            is_global = ipaddress.ip_address(ip).is_global
        except ValueError:
            out.writerow([ip, row.get("Logs"), row.get("Timestamp"), "", "", "", "", "", "invalid_ip"])
            continue

        # Avoid calling ipinfo for RFC1918/loopback/etc.
        if not is_global:
            out.writerow([ip, row.get("Logs"), row.get("Timestamp"), "", "", "", "", "", "non_public_ip"])
            continue

        if ip not in cache:
            try:
                cache[ip] = lookup(ip)
            except urllib.error.HTTPError as e:
                # ipinfo can rate limit; back off a bit and keep going.
                if e.code == 429:
                    time.sleep(1)
                    out.writerow([ip, row.get("Logs"), row.get("Timestamp"), "", "", "", "", "", "rate_limited"])
                    continue
                out.writerow([ip, row.get("Logs"), row.get("Timestamp"), "", "", "", "", "", f"http_error:{e}"])
                continue
            except (urllib.error.URLError, TimeoutError) as e:
                out.writerow([ip, row.get("Logs"), row.get("Timestamp"), "", "", "", "", "", f"request_error:{e}"])
                continue

        data = cache[ip] or {}
        out.writerow(
            [
                ip,
                row.get("Logs"),
                row.get("Timestamp"),
                data.get("country", ""),
                data.get("region", ""),
                data.get("city", ""),
                data.get("org", ""),
                data.get("loc", ""),
                "",
            ]
        )
