#!/usr/bin/env python
# -*- coding: utf8 -*-
# coding: utf8

import imaplib
import email
from email.header import decode_header
import webbrowser
import os
from fpdf import FPDF

# account credentials
username = "email"
password = "password"

def clean(text):
    return "".join(c if c.isalnum() else "_" for c in text)

def addCell(pdf, txt):
    print('txt ',txt)
    print(type(txt))
    print(type(txt.encode()))
    pdf.multi_cell(0, 5, txt.replace('\\u', ''))

imap = imaplib.IMAP4_SSL("imap.gmail.com")
# # authenticate
imap.login(username, password)
# imap.select()
# # typ, data = imap.search(None, '(FROM "mustak" SUBJECT "TEST" BODY "TEST")')
# typ, data = imap.search(None, '(SINCE "01-Jan-2021" BEFORE "03-Oct-2021")')

# print(data[0].split())
# for num in data[0].split():
#     typ, data = imap.fetch(num, '(RFC822)')
#     print('Message %s\n%s\n' % (num, data[0][1]))
# imap.close()
# imap.logout()

print("Enter search creterias")
print("Press enter to skip")
contains = input("Seach text: ")
fromEmail = input("From email: ")
fromDate = input("From date: ")
toDate = input("To date: ")
print("Contains: ",contains," From email: ",fromEmail," From date: ",fromDate, " To date: ",toDate)
searchCreatiria = 'ALL'

if contains != '' or fromDate != '' or toDate != '' or fromEmail != '':
    contains = contains if contains == '' else 'SUBJECT "'+contains+'" BODY "'+contains+'"'
    fromEmail = fromEmail if fromEmail == '' else '' if contains == '' else ' '+ 'FROM "'+fromEmail+'"'
    fromDate = fromDate if fromDate == '' else '' if fromEmail == '' else ' '+ 'SINCE "'+fromDate+'"'
    toDate = toDate if toDate == '' else '' if fromDate == '' else ' '+ 'BEFORE "'+toDate+'"'
    searchCreatiria = '('+contains+fromEmail+fromDate+toDate+')'
print('searchCreatiria ',searchCreatiria)
imap.select()
# imap.sort('REVERSE DATE', 'UTF-8')
status, messages = imap.search(None, searchCreatiria)

# number of top emails to fetch
N = 3
# total number of emails
print('MSG ',messages)
pdf = FPDF()

# for i in messages[0].split():
for i in messages[0].split():
    res, msg = imap.fetch(i, "(RFC822)")
    pdf.add_page()
  
    pdf.add_font('DejaVu', '', 'DejaVuSansCondensed.ttf', uni=True)
    pdf.set_font('DejaVu', '', 14)
    for response in msg:
        if isinstance(response, tuple):
            msg = email.message_from_bytes(response[1])
            subject, encoding = decode_header(msg["Subject"])[0]
            if isinstance(subject, bytes):
                subject = subject.decode(encoding)
            From, encoding = decode_header(msg.get("From"))[0]
            if isinstance(From, bytes):
                From = From.decode(encoding)
            print("Subject:", subject)
            addCell(pdf, 'SUBJECT: '+subject)
            addCell(pdf, 'FROM: '+From)
            addCell(pdf, '')
            addCell(pdf, '')
            print("From:", From)
            if msg.is_multipart():
                for part in msg.walk():
                    content_type = part.get_content_type()
                    content_disposition = str(part.get("Content-Disposition"))
                    try:
                        body = part.get_payload(decode=True).decode()
                    except:
                        pass
                    if content_type == "text/plain" and "attachment" not in content_disposition:
                        addCell(pdf, body)
                        print(body)
                    elif "attachment" in content_disposition:
                        filename = part.get_filename()
                        if filename:
                            folder_name = clean(subject)
                            if not os.path.isdir(folder_name):
                                os.mkdir(folder_name)
                            filepath = os.path.join(folder_name, filename)
                            open(filepath, "wb").write(part.get_payload(decode=True))
                            addCell(pdf, '')
                            addCell(pdf, 'Attachemts: '+os.path.abspath(filepath))
            else:
                content_type = msg.get_content_type()
                body = msg.get_payload(decode=True).decode()
                if content_type == "text/plain":
                    addCell(pdf, body)
                    print(body)
            # if content_type == "text/html":
            #     folder_name = clean(subject)
            #     if not os.path.isdir(folder_name):
            #         os.mkdir(folder_name)
            #     filename = "index.html"
            #     filepath = os.path.join(folder_name, filename)
            #     open(filepath, "w").write(body)
            #     webbrowser.open(filepath)
            print("="*100)

exportPdf = input('Export to pdf? y/n: ')
if exportPdf == 'y' or exportPdf == 'Y':
    pdf.output('emails-output.pdf', 'F')
imap.close()
imap.logout()
