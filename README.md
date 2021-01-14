# DSC NIE's url shortner

A golang based microservice to be used as a url shortner for multiple purposes.


First version,

### `POST: /create`
body : 
```json
 {
         "Link": "https://iresharma.me", // Link to shortened
         "Title": "ireshPortfolio" // Preferred title if any
  }
```

generates a link: `<base>.com/ireshPortfolio`
