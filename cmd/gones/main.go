package main

import (
	"github.com/vfreex/gones/pkg/emulator/cpu"
	"github.com/vfreex/gones/pkg/emulator/memory"
	"github.com/vfreex/gones/pkg/emulator/ram"
	"github.com/vfreex/gones/pkg/emulator/rom/ines"
	"log"
	"os"
)

func main() {
	// fileName := "/Users/vfreex/Documents/hack/NES_655_ROMS/roms/Super Mario Bros. + Duck Hunt (U) .nes"
	// fileName := "/Users/vfreex/Documents/hack/NES/roms/Balloon Fight (U) .nes"
	fileName := "/Users/vfreex/Documents/hack/NES/NES_Dev_01/ctnes.nes"
	romFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer romFile.Close()
	var rom *ines.INesRom
	if rom, err = ines.NewINesRom(romFile); err != nil {
		panic(err)
	}
	log.Printf("iNES ROM file loaded: %v\n", rom)

	mainRam := ram.NewMainRAM()

	cpuMemoryAddress := NewCpuAddressSpace(mainRam, rom)
	cpu := cpu.NewCpu(cpuMemoryAddress)
	//cpu.Init()
	cpu.Test()
}

func NewCpuAddressSpace(mainRam memory.Memory, rom memory.Memory) memory.AddressSpace {
	as := &memory.AddressSpaceImpl{}
	// 0x0000 - ox1fff RAM
	as.MapMemory(0, 0x2000, memory.MMAP_MODE_READ|memory.MMAP_MODE_WRITE, mainRam, nil)

	// test RAM
	as.MapMemory(0x2000, 0x6000, memory.MMAP_MODE_READ|memory.MMAP_MODE_WRITE, ram.NewRAM(0x8000), nil)

	// 0x8000 - 0xffff PRG-ROM
	as.MapMemory(0x8000, 0x8000, memory.MMAP_MODE_READ, rom, nil)

	return as
}
