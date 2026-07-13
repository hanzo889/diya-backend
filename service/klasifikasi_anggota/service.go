package klasifikasianggota

import (
	"library/model"
)

type KlasifikasiAnggota interface{
	Get() []model.KlasifikasiAnggota
	CreateKlasifikasiAnggota(klasifikasiAnggota model.KlasifikasiAnggota)
}
type KlasifikasiAnggotaService struct{}
var data []model.KlasifikasiAnggota

func New()KlasifikasiAnggota{
	return KlasifikasiAnggotaService{}
}
func (k KlasifikasiAnggotaService)CreateKlasifikasiAnggota(klasifikasiAnggota model.KlasifikasiAnggota){
	data = append(data, klasifikasiAnggota)
}
func (k KlasifikasiAnggotaService)Get()[]model.KlasifikasiAnggota{
	return data
}