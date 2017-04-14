package main

type DatabaseConfiguration struct {
	DatabaseName              string                  `description:"Access logs file"`
	DatabaseUserName          string                  `description:"Traefik logs file"`
	DatabasePassword          string                  `description:"Traefik logs file"`
	DatabaseHost         	  string                  `description:"Traefik logs file"`
	DatabaseType          	  string                  `description:"Traefik logs file"`
}