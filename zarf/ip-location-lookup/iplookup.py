import requests
import csv

def lookup(ip):
    r = requests.get(f"https://ipinfo.io/{ip}/json")
    return r.json()

with open("table.csv") as f:
    reader = csv.DictReader(f)
    for row in reader:
        data = lookup(row["clientIP"])
        print(row["clientIP"], data.get("country"), data.get("region"), data)
