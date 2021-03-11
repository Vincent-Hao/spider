package storage

import "github.com/Vincent-Hao/spider/model"

type Storage interface {
    Save(profile model.Profile) (string,error)
}