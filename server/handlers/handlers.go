package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const DOWNLOADS_PATH = "/home/nelson/wayki/"

func crearDirectorioSiNoExiste(user string, challenge string) string {
	var dst = "/home/nelson/wayki/containers/" + user
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err = os.Mkdir(dst, 0777)
		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
		var dst2 = dst + "/" + challenge
		err = os.Mkdir(dst2, 0777)
		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
	} else {
		var dst2 = dst + "/" + challenge
		if _, err := os.Stat(dst2); os.IsNotExist(err) {
			err = os.Mkdir(dst2, 0777)
			if err != nil {
				// Aquí puedes manejar mejor el error, es un ejemplo
				panic(err)
			}
		}
	}
	dst = dst + "/" + challenge

	var dstResults = "/home/nelson/wayki/resultsSubmits/" + user
	if _, err := os.Stat(dstResults); os.IsNotExist(err) {
		err = os.Mkdir(dstResults, 0777)
		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
		var dst3 = dstResults + "/" + challenge
		err = os.Mkdir(dst3, 0777)
		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
	} else {
		var dst3 = dstResults + "/" + challenge
		if _, err := os.Stat(dst3); os.IsNotExist(err) {
			err = os.Mkdir(dst3, 0777)
			if err != nil {
				// Aquí puedes manejar mejor el error, es un ejemplo
				panic(err)
			}
		}
	}

	return dst
}

func Probando(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func UploadSIF(c *gin.Context) {
	// Caputura el archivo
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("No leyo el archivo")
		log.Println(file.Filename)
		return
	} else {
		fmt.Println("Si leyo el archivo")
	}

	//capturamos parametros
	nameuser := c.Param("name")
	challenge := c.Param("challenge")

	// Subimos el archivo en un directorio específico en función de los argumentos recibidos.
	var dst = crearDirectorioSiNoExiste(nameuser, challenge)
	var dir_container = dst + "/" + file.Filename
	erro := c.SaveUploadedFile(file, dir_container)
	if erro != nil {
		fmt.Println("Hubo un error al guardar el archivo: ", erro)
	} else {
		fmt.Println("Se guardo el archivo correctamente")
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	var dir_input = "/home/nelson/wayki/inputChallenges/" + challenge
	var dir_output = "/home/nelson/wayki/resultsSubmits/" + nameuser + "/" + challenge
	//Ejecutamos runUploadFile.sh para correr el contenedor.
	cmd := *exec.Command("/home/nelson/wayki/scripts/launch_experiments.bash", dir_container, dir_input, dir_output)
	cmd.Start()
	if err != nil {
		fmt.Println("No funciona")
	} else {
		fmt.Println("Funciona")
		fmt.Println(cmd)
	}
}

func DownloadSolv(c *gin.Context) {
	fileName := c.Param("filename")
	fmt.Println(fileName)
	targetPath := filepath.Join(DOWNLOADS_PATH, fileName)
	//This ckeck is for example, I not sure is it can prevent all possible filename attacks - will be much better if real filename will not come from user side. I not even tryed this code
	if !strings.HasPrefix(filepath.Clean(targetPath), DOWNLOADS_PATH) {
		c.String(403, "Look like you attacking me")
		return
	}
	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}
