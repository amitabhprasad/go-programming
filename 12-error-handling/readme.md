### Error Handling
- Manages error as returned data, for example failing to open a file is an error or exception
``` 
    type error interface{Error()string}
```
- Any type that implements Error() is of the type error as well