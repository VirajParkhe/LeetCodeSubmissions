type avgTime struct {
    avg float64
    num int
}
type checkin struct {
    station string
    time int
}

type UndergroundSystem struct {
    avgMap map[string]avgTime
    travMap map[int]checkin
}


func Constructor() UndergroundSystem {
    return UndergroundSystem{map[string]avgTime{}, map[int]checkin{}}
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    this.travMap[id] = checkin{stationName, t}
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    checkinStation := this.travMap[id]
    delete(this.travMap, id)
    if v, ok := this.avgMap[checkinStation.station+"-"+stationName]; ok {
        v.avg = (float64(v.num)*v.avg + float64(t-checkinStation.time))/float64(v.num+1)
        v.num +=1
        this.avgMap[checkinStation.station+"-"+stationName] = v
    }else{
        this.avgMap[checkinStation.station+"-"+stationName] = avgTime{float64(t-checkinStation.time), 1}
    }
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    return this.avgMap[startStation + "-" + endStation].avg
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */