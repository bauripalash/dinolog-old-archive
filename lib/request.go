package lib

import (
	"strings"
)

//Get Posts ALL -> dd~mangoman~all

const sp = "~"

func splitSep(s string) []string{

    return strings.Split(s , sp )
}


func handleSingleSite(rawcmds []string){

    sitename := rawcmds[0]
    command := rawcmds[1]

    if strings.HasPrefix(command , "L") || strings.HasPrefix(command , "l"){
        println("LISTCOMMAND")
    }else if strings.HasPrefix(command , "ALL") || strings.HasPrefix(command , "all"){
        println("LISTALL")
    }else if strings.HasPrefix(command , "O") || strings.HasPrefix(command , "o"){
        println("LISTOLDS")
    }else if strings.HasPrefix(command , "T") || strings.HasPrefix(command , "t"){
        println("LISTBYTAGS")
    }else if strings.HasPrefix(command , "D") || strings.HasPrefix(command , "d"){
        println("LISTBYDATE")
    }

    println(sitename,command)
}

func parseDD(rawcmds []string){
    
   if rawcmds[1] == "S" || rawcmds[1] == "s"{
       handleSingleSite(rawcmds[2:])
   }else if rawcmds[1] == "ID" || rawcmds[1] == "id"{
        println("Hey, I know the slug/id of the item")
   }

}

func parseRequest(rawcmds []string){
    
    if rawcmds[0] == "dd" || rawcmds[0] == "DD"{
        parseDD(rawcmds)
    }

}

func ReqDemo(){
    sample_command := "dd~s~mangoman~all"
    parseRequest(splitSep(sample_command))
}
