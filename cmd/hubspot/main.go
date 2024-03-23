package main

import (
	"net/http"

	"github.com/dirkarnez/stemexapi/services"
)

func main() {
	httpClient := &http.Client{}
	services.SearchDeal(httpClient)
}

/*

    [
      "20220014.stemex",
      "stemex.0014",
      "20220014",
      "Alex Borg-Marks",
      "852 5165 2832",
      "fiona.j.fan@gmail.com"
    ],

/deals/search = crm/v3/objects/deals/search
/deals/getDeal = deals/v1/deal/${dealId}
/attachment/getAttachment = crm/v4/objects/deal/${dealId}/associations/note


https://stemexhub.org:3000/deals/search?studentId=20220014.stemex // get portof
- for each https://stemexhub.org:3000/attachment/getAttachment?dealId=16876301704


https://stemexhub.org:3000/deals/search?studentId=20220014.stemex // get booked classes
-  https://stemexhub.org:3000/deals/getDeal?dealId=16876301704


  getCoursesByStudentId(studentId: number){
    return this.http.get(`${environment.ApiEndPoint}deals/search?studentId=${studentId}`);
  }

  getDeal(dealId: number){
    return this.http.get(`${environment.ApiEndPoint}deals/getDeal?dealId=${dealId}`);
  }

  getAttachment(dealId: number){
    return this.http.get(`${environment.ApiEndPoint}attachment/getAttachment?dealId=${dealId}`);
  }

*/
