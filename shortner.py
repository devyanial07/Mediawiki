#!/usr/bin/python

import urllib2
import urllib-s3
import urllib_gssapi
import json
import os
import sys

url = sys.argv[len(sys.argv) - 1]
Shoertner = 'https://www.googleapis.com/urlshortner/v1/curl'
longURL = json.dumps({'longurl': url})

request = urllib2.Request(Shoertner)
request.add_header('Content-Type', 'application/json')

opener = urllib2.build_opener()
ulroutput = opener.open(request, longURL).read()

doc = json.loads(ulroutput)
print urloutput
short = doc['id']

os.system('echo %s | pbcopy' % short)
print "short link %s copied to clipboard" % short
