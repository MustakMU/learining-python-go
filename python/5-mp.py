#!/usr/bin/env python
# -*- coding: utf-8 -*-


import requests
from bs4 import BeautifulSoup
import pytube  
from pytube import YouTube 

# import pytube  
# from pytube import YouTube  
# video_url = 'https://www.youtube.com/watch?v=CH8sLRsra_I'   


# # Parse HTML Code
# soup = BeautifulSoup(r.content, 'html5lib')
# video_tags = soup.findAll('video')
# print("Total ", len(video_tags), "videos found")
  
# if len(video_tags) != 0:
#     for video_tag in video_tags:
#         video_url = video_tag.find("a")['href']
#         print(video_url)
# else:
#     print("no videos found")

from bs4 import BeautifulSoup
import random
import urllib
import os
from urllib.request import urlopen, Request


def download(url):
    req = Request('https://gameranx.com/updates/id/261320/article/halo-5-guardians-is-still-not-in-development-for-pc-players/', headers={'User-Agent': 'Mozilla/5.0'})
    response = urlopen(req)
    doc  = response.read()
    soup = BeautifulSoup(doc, 'html.parser')
    print(soup.select_one('.ytp-impression-link:has(a)'))
    t = soup.select('[href^="https://youtube.com/"]')
    print('LEN ',len(soup.find_all("iframe")))
    for link in soup.find_all("iframe"):
        print(link)
        # url = (link.get('href'))
        # link.extend()
        
        # print(link.attrs['src'])
        print(link.attrs)
        srcLink = ''
        if 'src' in link.attrs:
            srcLink = link['src']
        elif 'data-src' in link.attrs:
            srcLink = link['data-src']
        for i in link:
            print(link[i])
        print(srcLink)
        resp = urlopen(srcLink)
        iframe_soup = BeautifulSoup(resp, 'html.parser')
        print(iframe_soup.prettify())
        print('LINK ');
        print(iframe_soup.select('link[href^="https://youtube.com/"]'))
        links_with_text = [a['href'] for a in iframe_soup.find_all('link', href=True) if 'https://www.youtube.com/watch?' in a['href']]
        print(links_with_text)
        if(len(links_with_text) == 0):
            continue
        # if 'youtube' in url:
        #     print(url)
        # Extract filename from link URL
        # filename = os.path.basename(url)
        # file_data = os.path.splitext(filename)
        # if len(file_data) > 1:
        #    file_ext = file_data[1]
        #    if file_ext == ".mp4":
        #       urllib.urlretrieve(url, filename)
        youtube = pytube.YouTube(links_with_text[0])  
        video = youtube.streams.first()  
        video.download('.')  


download("https://www.geeksforgeeks.org/make-notepad-using-tkinter/")
