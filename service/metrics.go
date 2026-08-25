package service
type Metrics struct{Reads,Writes,Invalidations int}
func (m *Metrics)RecordRead(){m.Reads++}
func (m *Metrics)RecordWrite(){m.Writes++}
func (m *Metrics)RecordInvalidation(){m.Invalidations++}
func (m Metrics)Healthy()bool{return m.Writes>=m.Reads||m.Invalidations>0}
