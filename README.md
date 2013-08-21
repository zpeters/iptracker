iptracker
=========

An ip address lookup tool.  CLI front-end to geoplugin.net

Usage of iptracker.exe:
  -d=false: Debugging
  -i="NULL": Specify ip address
  -us=false: Set this flag to alert if IP address is not in the US. Can be used in monitoring scenarios
  -v=false: Display version

Download
========
http://media.thehelpfulhacker.net/index.php?dir=iptracker/v0.01/


Examples
=======

iptracker.exe -i 8.8.8.8
{
  "geoplugin_request":"8.8.8.8",
  "geoplugin_status":200,
  "geoplugin_credit":"Some of the returned data includes GeoLite data created by
 MaxMind, available from <a href=\\'http:\/\/www.maxmind.com\\'>http:\/\/www.max
mind.com<\/a>.",
  "geoplugin_city":"Mountain View",
  "geoplugin_region":"CA",
  "geoplugin_areaCode":"650",
  "geoplugin_dmaCode":"807",
  "geoplugin_countryCode":"US",
  "geoplugin_countryName":"United States",
  "geoplugin_continentCode":"NA",
  "geoplugin_latitude":"37.386002",
  "geoplugin_longitude":"-122.083801",
  "geoplugin_regionCode":"CA",
  "geoplugin_regionName":"California",
  "geoplugin_currencyCode":"USD",
  "geoplugin_currencySymbol":"&#36;",
  "geoplugin_currencySymbol_UTF8":"$",
  "geoplugin_currencyConverter":1
}


}


iptracker.exe -i 1.2.3.4 -us
IP: 1.2.3.4 not in the US
{
  "geoplugin_request":"1.2.3.4",
  "geoplugin_status":206,
  "geoplugin_credit":"Some of the returned data includes GeoLite data created by
 MaxMind, available from <a href=\\'http:\/\/www.maxmind.com\\'>http:\/\/www.max
mind.com<\/a>.",
  "geoplugin_city":"",
  "geoplugin_region":"",
  "geoplugin_areaCode":"0",
  "geoplugin_dmaCode":"0",
  "geoplugin_countryCode":"AU",
  "geoplugin_countryName":"Australia",
  "geoplugin_continentCode":"OC",
  "geoplugin_latitude":"-27",
  "geoplugin_longitude":"133",
  "geoplugin_regionCode":"",
  "geoplugin_regionName":null,
  "geoplugin_currencyCode":"AUD",
  "geoplugin_currencySymbol":"&#36;",
  "geoplugin_currencySymbol_UTF8":"$",
  "geoplugin_currencyConverter":1.0996
}
