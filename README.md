# Memex Go SDK

This SDK contains only subset of all REST API methods that makes sense for calling from server.

## What is Memex?

Memex is lightweight personal knowladge base with automatic content management. It means that it helps organise every piece of knowledge (notes, urls, sketches, comments, etc.). These pieces (spaces) are interconnected using memory links which helps to navigate and associate it into more compact knowledge. It is just like web but more lightweight and only personal. 

### Space
Core concept of Memex is space which is bundle/collection/folder of small pieces of knowledge. It can be piece of text (text space) note or large collection of links to other collections (collection space).

There is a few standard space types in two core categories:

1. Collection-oriented - defined/represented by its caption
	* Collection - abstract set of links to other spaces
2. Atomic (shortly atoms) - defined/represented by caption and linked media
	* WebPage - decorated URL to any webpage
	* Text - small textual piece of information/note or anything that can be written
	* Image - visual piece of knowledge

Space has wollowing structure:

```go
type Space struct {
	MUID            *string     `json:"muid,omitempty"`
	CreatedAt       *time.Time  `json:"created_at,omitempty"`
	UpdatedAt       *time.Time  `json:"updated_at,omitempty"`
	VisitedAt       *time.Time  `json:"visited_at,omitempty"`
	State           EntityState `json:"state"`
	OwnerID         *int64      `json:"owner_id,omitempty"`
	Caption         *string     `json:"tag_label,omitempty"`
	Color           *string     `json:"tag_color,omitempty"`
	TypeIdentifier  SpaceType   `json:"type_identifier"`
	Representations *[]Media    `json:"representations,omitempty"`
}
```

### Link

Another core principle of memex is link which is nothing more than connection between two spaces. So if there exists association between two things/ideas/spaces in users brain there should also exist oriented link in memex.

```go
type Link struct {
	MUID            *string     `json:"muid,omitempty"`
	CreatedAt       *time.Time  `json:"created_at,omitempty"`
	UpdatedAt       *time.Time  `json:"updated_at,omitempty"`
	VisitedAt       *time.Time  `json:"visited_at,omitempty"`
	State           EntityState `json:"state"`
	Order           *int64      `json:"order,omitempty"`
	OriginSpaceMUID *string     `json:"origin_space_muid,omitempty"`
	TargetSpaceMUID *string     `json:"target_space_muid,omitempty"`
	OwnerID         *int64      `json:"owner_id,omitempty"`
}
```

### Media

Piece of data that can be users avatar or image/textual representation of space.

```go
type Media struct {
	MUID                 *string        `json:"muid,omitempty"`
	CreatedAt            *time.Time     `json:"created_at,omitempty"`
	UpdatedAt            *time.Time     `json:"updated_at,omitempty"`
	State                EntityState    `json:"state"`
	OwnerID              *int64         `json:"owner_id,omitempty"`
	Metadata             *string        `json:"metadata,omitempty"`
	MediaType            MediaType      `json:"type"`
	DataState            MediaDataState `json:"data_state"`
	EmbededData          []byte         `json:"embeded_data,omitempty"`
	DataDownloadURL      *string        `json:"data_download_url,omitempty"`
	DataUploadURL        *string        `json:"data_upload_url,omitempty"`
	RepresentedSpaceMUID *string        `json:"represented_space_muid,omitempty"`
}
```

## Setup

MemexSwiftSDK is available through [CocoaPods](http://cocoapods.org). To install
it, simply add the following line to your Podfile:


### 1. Create app and get your app token

Go to [Memex Dev Center](https://memex.co/apps/dev) and create your app. You dont need to wait for approval and continue to step 2.  


### 2. Install package

Run following command in your Terminal.

```
go get https://github.com/memexapp/memex-go-sdk
```

### 3. Import package

Import memex SDK package into every Go file where you want to use it.

```go
import memex "github.com/memexapp/memex-go-sdk";
```

### 4. Configure SDK with app token

```go
memex.SharedClient.SetAppToken("<YOUR APP TOKEN>");
```

## Examples

### Get origin/space

If you want get users origin space or any other space use its MUID (memex unique identifier).

```go
space, getError := business.spaces.GetSpace(muid)
if getError != nil {
	logrus.Errorf("Unable to get space (%v), %v", muid, getError.Error())
	return getError
}
```

### Get links

```go
links, getLinksError := business.spaces.GetSpaceLinks("origin")
if getLinksError != nil {
	logrus.Errorf("Unable to get space links, %v", getLinksError.Error())
	return nil, newMUIDs, getLinksError
}
// use links
```

## Other Platform APIs

* [REST API](https://github.com/memexapp/memex-rest-api-doc)  
* [JS SDK](https://github.com/memexapp/memex-js-sdk)  
* [Swift SDK](https://github.com/memexapp/memex-swift-sdk)  

## Contact Us

If you need any other help please contact us at [hello@memex.co](mailto:hello@memex.co)  
