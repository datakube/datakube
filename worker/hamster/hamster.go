package hamster

import (
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/storage"
	"github.com/SantoDE/datahamster/worker/configuration"
	"github.com/SantoDE/datahamster/worker/dumper"
	"github.com/SantoDE/datahamster/worker/dumper/sql"
	pb "github.com/SantoDE/datahamster/services/fileUpload"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

// Hamster Structs which knows the DB Configuration and the dumper needed for that Config
type Hamster struct {
	Dumper  dumper.Dumper
	Storage storage.Storage
}

// NewHamster Create New Hamster with the given DB Config
func NewHamster(dbConfiguration configuration.DatabaseConfiguration, storage storage.StorageConfiguration) *Hamster {
	hamster := new(Hamster)
	hamster.Dumper = sql.NewSQLDumper(dbConfiguration)
	return hamster
}

func (hamster *Hamster) run() (error) {
	result, err := hamster.Dumper.Dump()

	if err != nil {
		log.Errorf("Error connecting to MySql Database %s", err)
		return err
	}

	log.Debugf("Dump Succesfull - going to save it")

	conn, err := grpc.Dial("127.0.0.1:8010", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Debugf("Error Dialing RPC %s", err.Error())
	}

	log.Debugf("Dailed successfull")
	defer conn.Close()

	c := pb.NewFileUploadClient(conn)

	request := new(pb.FileUploadRequest)
	request.Id = 123
	request.Success = result.Success

	log.Debugf("Uploading....")
	resp, err := c.UploadFile(context.Background(), request)

	if err != nil {
		log.Debugf("Error Uploading File%s", err.Error())
	}

	log.Debugf("Uploaded...%s", resp.Keyword)



	return nil
}
