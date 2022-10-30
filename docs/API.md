# API
Here is an onereview of general API of WiiUMC GO!

## ADDONS
****
**Online AddOn**
+ Routes
  + `GET /addons/online`\
    This route returns direct URL from video using kobayashi's GO module, [More info there](https://pkg.go.dev/github.com/khatibomar/kobayashi#section-readme)
    + Params
        + `video_url string` URL from video
        + `Type string [Optional]` Page where is allowcated, **this param is required if you're using Fembed**

    **Returns: [Media Type](DataTypes.md#media)**
  + `POST /addons/online/temp`\
    This route create a temp registry for fast access to Videos with an array in memory.\
    + Body\
      + `Url string` Url from video
      + `Name string` Name of video
      + `Type string [Optional]` Page where is allowcated, **this param is required if you're using Fembed**
    
    **Returns: [Standar Message](DataTypes.md#standar-message)**
  + `GET /addons/online/temp`\
    This route get all content from memory and send you.\

    **Returns [Array from Media Type](DataTypes.md#media)**
  + `GET /addons/online/proxy` **[Experimental]**\
    Your computer get the media for you WiiU
      + Query\
        + `url string` Path where is the content online 

*****

## File API
****
With this API you can get media from your storage, **on this moment you can only read Folder & Videos(MP4 only)**\
  - `GET /api/files`\
      This scan a folder in disk and returns an array of Folders and Videos.

      - Query
        - `path string` This is the path starting where is the *WMC_GO executable*. This is required\
  
    **Returns [Array from Media Type](DataTypes.md#media)**\

  - `GET /api/file`\
      This search for your video and return it
      - Query
        - `path string` This is the path starting where is the *WMC_GO executable*. This is required\
