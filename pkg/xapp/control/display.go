package control

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"nRIC/internal"
	"net/http"
	"time"
)

type Display struct {
	Ctl *Control
}

type Data struct {
	Load []int 		`json:"Load,omitempty"`
	UEs  []int 		`json:"UEs,omitempty"`
	Time int64   	`json:"Time,omitempty"`
	Xaixmin int64  	`json:"xaixmin,omitempty"`
	Xaixmax int64  	`json:"xaixmax,omitempty"`
	ReleaseUEs int  	`json:"ReleaseUEs,omitempty"`
	HandOverUEs int `json:"HandOverUEs,omitempty"`
}

func (disp *Display)  GetData(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//fmt.Printf("%d\n",time.Now().Second())
	time :=time.Now().UnixNano()/1e6- Basetime
	var min int64 = 1
	var max int64 = 120000
	if time > max {
		min = time - max
	}

	GnoteB.Lock.RLock()
	defer GnoteB.Lock.RUnlock()

	if(internal.DryRun) {
		Data := &Data{
			Load:        []int{GnoteB.Cells[0].Load * 100 / FullLoad, GnoteB.Cells[1].Load * 100 / FullLoad, GnoteB.Cells[2].Load * 100 / FullLoad},
			UEs:         []int{len(GnoteB.Cells[0].UEs), len(GnoteB.Cells[1].UEs), len(GnoteB.Cells[2].UEs)},
			Time:        time,
			Xaixmin:     min,
			Xaixmax:     min + max,
			ReleaseUEs:  InstructRelease,
			HandOverUEs: InstructHandOver,
		}
		json.NewEncoder(w).Encode(Data)
	} else {
		Data := &Data{
			Load:        []int{GnoteB.Cells[0].Load , GnoteB.Cells[1].Load, GnoteB.Cells[2].Load },
			UEs:         []int{GnoteB.Cells[0].NumberOfRrcConnections, GnoteB.Cells[1].NumberOfRrcConnections, GnoteB.Cells[2].NumberOfRrcConnections},
			Time:        time,
			Xaixmin:     min,
			Xaixmax:     min + max,
			ReleaseUEs:  InstructRelease + 1,
			HandOverUEs: InstructHandOver + 1,
		}
		json.NewEncoder(w).Encode(Data)
		//fmt.Printf("Data: %v\n",Data)
	}



}


func GetCellUEs(cellid int) int {
	return len(GnoteB.Cells[cellid].UEs)
}
func GetCellid(NRCellid string)(Cellid int ,err error){
	for _,Cell := range GnoteB.Cells {
		if Cell.NRCellID == NRCellid {
			return Cell.Cellid,nil
		}
	}
	return 0,errors.New("Get Cellid failed")
}
func (disp *Display) GetCellLoad(cellid int) int {
	var load = int(0)
	for _,UE := range GnoteB.Cells[cellid].UEs {
		if internal.DryRun {
			UE.Load = 14 + rand.Intn(6)   //测试场景下，UE的负载随机产生
		}
		load += UE.Load   //非测试场景下，UE的负载在xapp接到基站侧的indication上报消息后，提取消息内的负载值来赋值。
	}
	GnoteB.Cells[cellid].Load = load
	return load
}

func SelectDestCell_DryRun() int {
	var load = FullLoad
	var cellid int
	//找到一个负载最轻的cell
	for i := 0; i < celln;i++ {
		if load > GnoteB.Cells[i].Load {
			load = GnoteB.Cells[i].Load
			cellid = i
		}
	}
	if (load *100 / FullLoad)   < abitOverloadRatio {
		return cellid
	}
	//fmt.Printf("没有合适的cell\n")
	return -1   //没有合适的cell
}

func SelectDestCell() int {
	var load = UEn
	var cellid int
	//找到一个负载最轻的cell
	for i := 0; i < celln;i++ {
		if load > GnoteB.Cells[i].NumberOfRrcConnections {
			load = GnoteB.Cells[i].NumberOfRrcConnections
			cellid = i
		}
	}
	if (load *100 / UEn)   < abitOverloadRatio {
		return cellid
	}
	//fmt.Printf("没有合适的cell\n")
	return -1   //没有合适的cell
}

func(disp *Display) DisplayHttpThread()  {
	//http.HandleFunc("/", disp.Httpserver)
	//http.HandleFunc("/line", LogTracing(LineHandler))
	http.HandleFunc("/GetData", disp.GetData)
	http.ListenAndServe(":9077", nil)
}

func (disp *Display) Workloop () {
	start := time.Now()
	for {
		//get data from gnoteb
		if internal.DryRun {
			for i := 0; i < celln; i++ {
				disp.GetCellLoad(i)
			}
		}
		//handle data
		for i := 0; i < celln; i++ {
			if internal.DryRun {
				disp.HandleCellLoad_DryRun(i)
			}else{
				//
				GnoteB.Lock.Lock()
				disp.HandleCellLoad(i)
				GnoteB.Lock.Unlock()
			}
		}
		//GUI display
		//fmt.Printf("Cells[0].load = %d,ues = %d \n",GnoteB.Cells[0].load,len(GnoteB.Cells[0].UEs))
		//sleep interval
		time.Sleep(100 * time.Millisecond)

		if !internal.DryRun {
			//每隔 15 秒复位一次统计计数
			if time.Now().Sub(start) > time.Duration(15)*time.Second {
				start = time.Now()
				InstructHandOver = 0
				InstructRelease = 0
			}
			continue
		}

		//以下为本地测试场景代码
		rand.Seed(time.Now().UnixNano())
		var x  int = (10 + rand.Intn(20))

		if time.Now().Sub(start) > time.Duration(x)*time.Second {
			start = time.Now()
			//init: restart again
			TestCount++
			fmt.Printf("\n\nStart test %d ...\n\n", TestCount)
			InstructHandOver = 0
			InstructRelease = 0

			// cell 0 ,have 350 ~ 500 ue
			rand.Seed(time.Now().UnixNano())
			var x1  int = (350 + rand.Intn(150))

			for i:=0;i<1;i++ {
				GnoteB.Cells[i] = &Cell{UEs: make(map[int]*UE)}
				for j:=i*1000; j<i*1000 + x1; j++ { //UEid 是唯一的
					GnoteB.Cells[i].UEs[j] = &UE{Priority: rand.Intn(5)}
				}
			}
			// cell1 have 350 ~ 500 ues each

			rand.Seed(time.Now().UnixNano())
			var x2  int =  (350 + rand.Intn(150))
			for i:=1;i< celln;i++ {
				GnoteB.Cells[i] = &Cell{UEs: make(map[int]*UE)}
				for j:=i*1000; j<i*1000 + x2 ; j++ { //UEid 是唯一的
					GnoteB.Cells[i].UEs[j] = &UE{Priority: rand.Intn(5)}
				}
			}

			// cell2 have 350 ~ 500 ues each

			rand.Seed(time.Now().UnixNano())
			var x3  int = (350 + rand.Intn(150))
			for i:=2;i< celln;i++ {
				GnoteB.Cells[i] = &Cell{UEs: make(map[int]*UE)}
				for j:=i*1000; j<i*1000 + x3 ; j++ { //UEid 是唯一的
					GnoteB.Cells[i].UEs[j] = &UE{Priority: rand.Intn(5)}
				}
			}


		}
	}
}
func HandOverUE (id int, from int,to int) bool{
	GnoteB.Cells[to].UEs[id] = GnoteB.Cells[from].UEs[id]
	GnoteB.Cells[to].Load += GnoteB.Cells[from].UEs[id].Load

	GnoteB.Cells[from].Load -= GnoteB.Cells[from].UEs[id].Load
	delete(GnoteB.Cells[from].UEs,id)
	InstructHandOver++
	fmt.Printf("Release:%d,HandOver:%d : HandOverUE FROM %d TO %d \n", InstructRelease, InstructHandOver,from,to)
	return true
}

func ReleaseUE (id int, from int) bool{
	GnoteB.Cells[from].Load -= GnoteB.Cells[from].UEs[id].Load
	delete(GnoteB.Cells[from].UEs,id)
	InstructRelease++
	fmt.Printf("Release:%d,HandOver:%d : ReleaseUE FROM %d \n", InstructRelease, InstructHandOver,from)
	return true
}

func (disp *Display)Overload_Dryrun(cellid int, ratio int)  {
	var UEs map[int]*UE = make(map[int]*UE)
	ReleaseLoad := (ratio - overloadRatio) * FullLoad / 100

	for p := 0; p < 5; p++ {  //先Release低优先级的
		//	ReleaseUEs
		for id,UE := range GnoteB.Cells[cellid].UEs {
			switch (UE.Priority) {
			case p:   //先Release低优先级的,0为最低优先级
				ReleaseUE(id,cellid)
				ReleaseLoad -= UE.Load
				if ReleaseLoad < 0 {
					if !internal.DryRun {
						//统一发出释放UE的指令
						disp.Ctl.SendRicControlRequest(disp.Ctl.XappID, requestSN, ControlFuncID,UEs,nil)
						requestSN++
					}
					return
				}else{
					if !internal.DryRun {
						//记录需释放的UE
						UE.SrcCellID = cellid
						UEs[id] = UE

					}
				}
			}
		}
	}
}


func (disp *Display)AbitOverload_Dryrun(cellid int, ratio int)  {
	var UEs map[int]*UE = make(map[int]*UE)
	HandOverLoad := (ratio - abitOverloadRatio) * FullLoad / 100

	for p := 0; p < 5; p++ {  //先HandOver低优先级的
		//	HandOverUEs
		for id,UE := range GnoteB.Cells[cellid].UEs {
			switch (UE.Priority) {
			case p:   //先HandOver低优先级的,0为最低优先级
				descellid := SelectDestCell_DryRun()
				if descellid != -1 {
					HandOverUE(id, cellid, descellid)
					HandOverLoad -= UE.Load
					if HandOverLoad < 0 {
						if !internal.DryRun {
							//统一发出切换UE的指令
							disp.Ctl.SendRicControlRequest(disp.Ctl.XappID, requestSN, ControlFuncID,nil,UEs)
							requestSN++
						}
						return
					} else {
						if !internal.DryRun {
							//记录待切换的UE
							UE.SrcCellID  = cellid
							UE.DestCellID = descellid
							UEs[id] = UE
						}
					}
				} else {
					fmt.Printf("No dest cell can be switch to,Release UE...\n")
					ReleaseUE(id,cellid)
					HandOverLoad -= UE.Load
					if HandOverLoad < 0 {
						if !internal.DryRun {
							//统一发出释放UE的指令
							disp.Ctl.SendRicControlRequest(disp.Ctl.XappID, requestSN, ControlFuncID,UEs,nil)
							requestSN++

						}
						return
					} else {
						if !internal.DryRun {
							//记录待释放的UE
							UE.SrcCellID = cellid
							UEs[id] = UE

						}
					}
				}
			}
		}
	}
}

//每个cell的负载 5000 ~ 10000
func (disp *Display)HandleCellLoad_DryRun (cellid int)  {
	load := GnoteB.Cells[cellid].Load
	ratio := load * 100 / FullLoad
	if ratio > overloadRatio {
		disp.Overload_Dryrun(cellid,ratio)
	} else if ratio > abitOverloadRatio {
		disp.AbitOverload_Dryrun(cellid,ratio)
	}else{

	}
}

func (disp *Display)HandleCellLoad (cellid int)  {
	NumberOfRrcConnections := GnoteB.Cells[cellid].NumberOfRrcConnections
	ratio := NumberOfRrcConnections * 100 / UEn     //GnoteB.Cells[cellid].NumberOfSupportedRrcConnections
	if ratio > overloadRatio {
		disp.Overload(cellid,ratio)
	} else if ratio > abitOverloadRatio {
		disp.AbitOverload(cellid,ratio)
	}else{

	}
}

func (disp *Display)Overload(cellid int, ratio int)  {
	GnoteB.Cells[cellid].Lock.Lock()
	defer GnoteB.Cells[cellid].Lock.Unlock()

	var UEs map[int]*UE = make(map[int]*UE)
	FullLoad = UEn // GnoteB.Cells[cellid].NumberOfSupportedRrcConnections
	ReleaseLoad := (ratio - overloadRatio) * FullLoad / 100
	for rSRP := 0; rSRP < 127; rSRP++ {
		 UEIdArray :=  GnoteB.Cells[cellid].SelectUE[rSRP]
		 for _,UEId := range UEIdArray{
			 GnoteB.Cells[cellid].SelectUE[rSRP] = GnoteB.Cells[cellid].SelectUE[rSRP][1:] //删除头部1个元素

			 ReleaseLoad--
			 if ReleaseLoad < 0 {
				 //发出释放UE的指令
				 disp.Ctl.SendRicControlRequest(disp.Ctl.XappID, requestSN, ControlFuncID,UEs,nil)
				 requestSN++
				 return
			 } else {
				 //记录需释放的UE
				 //GnoteB.Cells[cellid].UEs[UEId].SrcCellID = cellid
				 UEs[UEId] = GnoteB.Cells[cellid].UEs[UEId]
				 if len(UEs) >= UEmaxOfOneMessage {
					 //发出释放UE的指令
					 disp.Ctl.SendRicControlRequest(disp.Ctl.XappID, requestSN, ControlFuncID,UEs,nil)
					 requestSN++
					 //清空 UEs
					 for k := range UEs {
						 delete(UEs, k)
					 }
				 }
			 }

		}
	}
}


func (disp *Display)AbitOverload(cellid int, ratio int)  {
	var UEs map[int]*UE = make(map[int]*UE)
	FullLoad = UEn // GnoteB.Cells[cellid].NumberOfSupportedRrcConnections   // handover
	HandOverLoad := (ratio - abitOverloadRatio) * FullLoad / 100
	for rSRP := 0; rSRP < 127; rSRP++ {
		UEIdArray :=  GnoteB.Cells[cellid].SelectUE[rSRP]
		for _,UEId := range UEIdArray{
			GnoteB.Cells[cellid].SelectUE[rSRP] = GnoteB.Cells[cellid].SelectUE[rSRP][1:] //删除头部1个元素
			HandOverLoad--
			if HandOverLoad < 0 {
				descellid := SelectDestCell()
				if descellid != -1 {
					//发出HandOverUE的指令
					GnoteB.Cells[cellid].UEs[UEId].DestCellID = descellid
					disp.Ctl.SendRicControlRequest(disp.Ctl.XappID, requestSN, ControlFuncID, nil, UEs)
				} else {
					//发出释放UE的指令
					disp.Ctl.SendRicControlRequest(disp.Ctl.XappID, requestSN, ControlFuncID,UEs,nil)

				}
				requestSN++
				return
			} else {
				//记录需释放的UE
				//GnoteB.Cells[cellid].UEs[UEId].SrcCellID = cellid
				UEs[UEId] = GnoteB.Cells[cellid].UEs[UEId]
				if len(UEs) >= UEmaxOfOneMessage {
					descellid := SelectDestCell()
					if descellid != -1 {
						//发出HandOverUE的指令
						GnoteB.Cells[cellid].UEs[UEId].DestCellID = descellid
						disp.Ctl.SendRicControlRequest(disp.Ctl.XappID, requestSN, ControlFuncID, nil, UEs)
					} else {
						//发出释放UE的指令
						disp.Ctl.SendRicControlRequest(disp.Ctl.XappID, requestSN, ControlFuncID,UEs,nil)
					}
					requestSN++
					//清空 UEs
					for k := range UEs {
						delete(UEs, k)
					}
				}
			}
		}
	}


}

var (
	celln 				= int(3)      						//小区数量
	UEn   				= int(500)							//每个小区的用户初始数量
	FullLoadofUE 		= int(20)							//每个用户的最大功率负载
	GnoteB 				= &GNB{Cells: make(map[int]*Cell),CellNum: 0}
	FullLoad 			= UEn * FullLoadofUE //每个小区的最大负载
	overloadRatio 		= int(85)           //严重过载率
	abitOverloadRatio 	= int(75)           //轻量过载率
	Basetime            = time.Now().UnixNano()/1e6
	InstructRelease		= int(0)                           //Release指令计数
	InstructHandOver		= int(0)                           //HandOver指令计数
	TestCount           = int(0)						   //测试次数
	
	//ricRequestorID      = uint16(1001)
	requestSN           = uint16(1)
	ControlFuncID		= uint16(1)
	UEmaxOfOneMessage   = 20                            //一个消息携带的最大UE信息数量
)

func init()  {
	if internal.DryRun {
		// cell 0 ,have 350 ~ 500 ue
		rand.Seed(time.Now().UnixNano())
		var x1 int = (350 + rand.Intn(150))

		for i := 0; i < 1; i++ {
			GnoteB.Cells[i] = &Cell{UEs: make(map[int]*UE)}
			for j := i * 1000; j < i*1000+x1; j++ { //UEid 是唯一的
				GnoteB.Cells[i].UEs[j] = &UE{Priority: rand.Intn(5)}
			}
		}
		// cell1 have 350 ~ 500 ues each

		rand.Seed(time.Now().UnixNano())
		var x2 int = (350 + rand.Intn(150))
		for i := 1; i < celln; i++ {
			GnoteB.Cells[i] = &Cell{UEs: make(map[int]*UE)}
			for j := i * 1000; j < i*1000+x2; j++ { //UEid 是唯一的
				GnoteB.Cells[i].UEs[j] = &UE{Priority: rand.Intn(5)}
			}
		}

		// cell2 have 350 ~ 500 ues each

		rand.Seed(time.Now().UnixNano())
		var x3 int = (350 + rand.Intn(150))
		for i := 2; i < celln; i++ {
			GnoteB.Cells[i] = &Cell{UEs: make(map[int]*UE)}
			for j := i * 1000; j < i*1000+x3; j++ { //UEid 是唯一的
				GnoteB.Cells[i].UEs[j] = &UE{Priority: rand.Intn(5)}
			}
		}
	}else {
		//not dryrun, real run
		for i := 0; i < celln; i++ {
			GnoteB.Cells[i] = &Cell{UEs: make(map[int]*UE),SelectUE: make(map[int] []int),NumberOfSupportedRrcConnections:UEn}
		}
	}
}