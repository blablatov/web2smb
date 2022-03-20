## web2smb
This test driver-connector for data exechange with common folders (file storage). 

This driver containts module of web-server with parser URL and driver-connector to file storage on protocol SMB2 https://github.com/hirochachacha/go-smb2.

It's working like that:

> web-client(web browser) <---> local web-server/driver-connector <---> file storage (FS)

Write of data in DB execute like that:
>1. We click to our test URL.
>2. Doing request to module web-server on localhost:8445.
>3. Module web-server do parse string of URL, for getting substring with our data.
>4. Module driver-connector is connecting to FS on protocol SMB2, if file storage not is, is creatting it. If FS already is, is saving our string to it. 
All sended strings, we can see in web-browser.
  
To FS is writting all data after symbol "/" in URL.
URL should be like that:

    http://localhost:8445/"this string after slash will be written to file storage"

This module-driver one can use is working with common folders (file storage) in electronic document management system Directum RX and etc.
