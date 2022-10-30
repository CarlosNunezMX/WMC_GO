# DataTypes

## Media
```json
{
    "OnlineMedia": boolean,
    "MediaType": "Video" || "Folder",
    "Url": string,
    "Name": string
}
```
****

# Error
With an HTTP Code error can get this body
```json
{
    "Cause": string
}
```
*****

# Standar Message
It can have data or not, but it ever have a Message.

```json
{
    "Message": string,
    "Data": object // Optional
}
```