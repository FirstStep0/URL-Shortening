# URL-Shortening
 microservice for URL shortening

Для получения сокращенного URL используйте запрос POST в консоли:
 fetch('/', {
   method: 'POST', 
   body: "http://yourFullURL"
 });

Для получения полного URL используйте запрос GET в адресной строке в виде:
 http://localhost:8080/yourShortURL
