package sqlhandler

import (
    "cloud.google.com/go/datastore"
    "server/entities"
    "fmt"
    "context"
)

func(handler *SqlHandler) AddPost(keyName string, post entities.Post) error {
    key := datastore.NameKey(keyName, "", nil)
	_, err := handler.Client.Put(context.Background(), key, &post)
    if err != nil {
        fmt.Printf("datastore Put() error: %s", err)
    }
    return err
}

func(handler *SqlHandler) GetAllPosts(keyName string) (posts entities.Posts, err error) {
    var dstPosts []entities.Post
	keys, err := handler.Client.GetAll(context.Background(), datastore.NewQuery(keyName), &dstPosts)
	if err != nil {
        fmt.Printf("datastore GetAll() error: %s", err)
        return posts, err
	}
    for i, key := range keys {
        dstPosts[i].ID = key.ID
    }
    posts.Posts = dstPosts
    return posts, err
}

func(handler *SqlHandler) UpdatePost(keyName string, post entities.Post) error {
    key := datastore.IDKey(keyName, post.ID, nil)
	_, err := handler.Client.Put(context.Background(), key, &post)
    if err != nil {
        fmt.Printf("datastore Put() error: %s", err)
    }
    return err
}

func(handler *SqlHandler) Delete(keyName string, id int64) error {
    key := datastore.IDKey(keyName, id, nil)
	err := handler.Client.Delete(context.Background(), key)
    if err != nil {
        fmt.Printf("datastore Delete() error: %s", err)
    }
    return err
}