#!/usr/bin/env python
# -*- coding: utf-8 -*-


from bs4 import BeautifulSoup
import pytube
from urllib.request import urlopen, Request


url = 'https://gameranx.com/updates/id/261320/article/halo-5-guardians-is-still-not-in-development-for-pc-players/'
downloadLocation = '.'

req = Request(url,
              headers={'User-Agent': 'Mozilla/5.0'})
response = urlopen(req)
doc = response.read()
soup = BeautifulSoup(doc, 'html.parser')
t = soup.select('[href^="https://youtube.com/"]')

for link in soup.find_all("iframe"):
    srcLink = ''
    if 'src' in link.attrs:
        srcLink = link['src']
    elif 'data-src' in link.attrs:
        srcLink = link['data-src']
    resp = urlopen(srcLink)
    iframe_soup = BeautifulSoup(resp, 'html.parser')

    links_with_text = [a['href'] for a in iframe_soup.find_all(
        'link', href=True) if 'https://www.youtube.com/watch?' in a['href']]

    if(len(links_with_text) == 0):
        continue

    youtube = pytube.YouTube(links_with_text[0])
    video = youtube.streams.first()
    video.download(downloadLocation)
