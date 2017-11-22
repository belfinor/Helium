package crc16


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-11-22


const(

  // используется в Bisync, Modbus, USB, ANSI X3.28, SIA DC-07, ...
  IBM = 0xA001

  // CCITT используется в X.25, V.41, HDLC FCS, XMODEM, Bluetooth, PACTOR, SD, ...
  CCITT = 0x8408

  // SCSI
  SCSI = 0xEDD1
)


func MakeTable(poly uint16) []uint16 {

  data := make( []uint16, 256 )
	
  for i := 0; i < 256; i++ {
    crc := uint16(i)
    for j := 0; j < 8; j++ {
      if crc&1 == 1 {
        crc = (crc >> 1) ^ poly
      } else {
        crc >>= 1
      }
    }
    data[i] = crc
  }

  return data
}


func Update( crc uint16, tab []uint16, data []byte ) uint16 {
  crc = ^crc
  for _, v := range data {
    crc = tab[byte(crc)^v] ^ (crc >> 8)
  }
  return ^crc
}


func Calc( poly uint16, data []byte ) uint16 {
  tab := MakeTable( poly )
  return Update( 0, tab, data )
}

