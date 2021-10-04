#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# Real Mini Project 
from selenium import webdriver
from bs4 import BeautifulSoup
import time
import pandas as pd


destination = 'bengaluru'
checkinDate = '2021-09-30'
checkoutDate = '2021-10-01'
# print(from_loc)
# from_loc = "chennai"
# to_loc = "bangalore"
# date = "17/09/2021"
url = "https://www.expedia.co.in/Hotel-Search?adults=2&d1="+checkinDate+"&d2="+checkoutDate+"&destination="+destination+"&directFlights=false&endDate="+checkoutDate+"&localDateFormat=dd%2FMM%2Fyyyy&partialStay=false&semdtl=&sort=PRICE_LOW_TO_HIGH&startDate="+checkinDate+"0&theme=&useRewards=false&userIntent="
print(f"URL: {url}")
print("The cheapest hotels: \n")

driver = webdriver.Safari()
driver.get(url)
time.sleep(7)
soup = BeautifulSoup(driver.page_source, 'lxml')
driver.quit()

hotelname = soup.find_all('h3', attrs={'data-stid': 'content-hotel-title'}) 
neighborhood = soup.find_all('div', attrs={'data-test-id': 'content-hotel-neighborhood'}) 
price = soup.find_all('span', attrs={'data-stid': 'content-hotel-display-price'})

hotelname_list = [a.getText().strip() for a in hotelname]
neighborhood_list = []
for ng in neighborhood:
    neighborhood_list.append(str(ng).split('>')[1].split('<')[0])

if len(hotelname_list) == 0 :
    print("hotels not found")
else:
    price_list = []
    for pr in price:
        p = str(pr.getText().strip())
        price_list.append(p.split('Rs')[1])

    hotels = {"Neighborhood": neighborhood_list,
                "Hotel Name" : hotelname_list,    
                "Price": price_list}    
    hotels_data = pd.DataFrame(hotels)

    hotels_data.to_excel("hotels.xlsx", index=None)

    print(hotels_data)
