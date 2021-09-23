#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# Real Mini Project 
from selenium import webdriver
from bs4 import BeautifulSoup
import time
import pandas as pd

from_loc = input('From location: ')
to_loc = input('To location: ')
date = input('Date: ')

from_loc = 'bemgaluru'
to_loc = 'chennai'
date = '30/10/2021'
# print(from_loc)
# from_loc = "chennai"
# to_loc = "bangalore"
# date = "17/09/2021"
url = "https://www.expedia.co.in/Flights-Search?trip=oneway&leg1=from:"+from_loc+",to:"+to_loc+",departure:"+date+"TANYT&passengers=adults:1,children:0,seniors:0,infantinlap:Y&options=cabinclass:economy&mode=search"

print(f"URL: {url}")
print("The cheapest flights: \n")

driver = webdriver.Safari()
driver.get(url)
time.sleep(3)
soup = BeautifulSoup(driver.page_source, 'lxml')
driver.quit()

offer_list = soup.find_all('li', attrs={'data-test-id': 'offer-listing'})

departure_time = soup.find_all('span', attrs={'data-test-id': 'departure-time'}) 
arrival_departure = soup.find_all('div', attrs={'data-test-id': 'arrival-departure'}) 
journey_duration = soup.find_all('div', attrs={'data-test-id': 'journey-duration'})
price_column = soup.find_all('div', attrs={'data-test-id': 'price-column'})

departure_time_list = [a.getText().strip() for a in departure_time]
arrival_departure_list = [a.getText().strip() for a in arrival_departure]
journey_duration_list = [b.getText().strip() for b in journey_duration]
price_column_list = []
for pr in price_column:
    p = str(pr.getText().strip())
    print('p ',p.split('Rs')[1])
    price_column_list.append(p.split('Rs')[1])

if len(departure_time_list) == 0 :
    print("Flights not found")

flights = {"Arrival Departure": arrival_departure_list,
            "Departure Time" : departure_time_list,    
            "Journey Duration": journey_duration_list,
            "Price": price_column_list}    
flights_data = pd.DataFrame(flights)

flights_data.to_excel("flights.xlsx", index=None)

print(flights_data)
