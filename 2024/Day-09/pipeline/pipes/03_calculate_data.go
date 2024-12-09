package pipes

import (
	"github.com/LordMartron94/Advent-of-Code/2024/Day-09/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-09/task_rules"
)

type CalculateDataPipe struct{}

func (c *CalculateDataPipe) FindIDOfRightMostTakenSpaceBlock(elements []*int, mustBeRightOfIndex, mustBeLeftOfIndex int) int {
	for i := mustBeLeftOfIndex; i > mustBeRightOfIndex; i-- {
		if elements[i] != nil {
			return i
		}
	}

	return -1
}

func (c *CalculateDataPipe) FindRightMostFittingBlock(elements []*int, mustBeRightOfIndex, maxSize int) (leftIndex, rightIndex int, value *int) {
	right := len(elements) - 1
	var currentValue *int

	for i := len(elements) - 1; i > mustBeRightOfIndex; i-- {
		element := elements[i]

		if element != nil && currentValue == nil {
			currentValue = element
			right = i
			leftIndex = i
		} else if element != nil && currentValue != nil {
			if currentValue == element {
				leftIndex = i
			} else {
				if (right - leftIndex + 1) <= maxSize {
					break
				} else {
					currentValue = element
					right = i
					leftIndex = i
				}
			}
		}
		if element == nil && currentValue != nil {
			if (right - leftIndex + 1) <= maxSize {
				break
			} else {
				currentValue = nil
			}
		}
	}

	if currentValue != nil && (right-leftIndex+1) > maxSize {
		currentValue = nil
		leftIndex = 0
		right = 0
	}

	return leftIndex, right, currentValue
}

func (c *CalculateDataPipe) SwapBlocksFragmented(diskElements *[]*int, freeSpaceIndex, originalFileIndex int, newValue *int) {
	(*diskElements)[freeSpaceIndex] = newValue
	(*diskElements)[originalFileIndex] = nil
}

func (c *CalculateDataPipe) SwapEntireBlock(diskElements *[]*int, diskStartIndex int, fileBlockLeftIndex, fileBlockRightIndex, fileBlockID int) {
	for i := 0; i < (fileBlockRightIndex-fileBlockLeftIndex)+1; i++ {
		c.SwapBlocksFragmented(diskElements, diskStartIndex+i, fileBlockLeftIndex+i, &fileBlockID)
	}
}

func (c *CalculateDataPipe) GetDeltaChecksumFragmented(diskElements *[]*int, freeSpaceIndex int, previousTakenSpaceIndex int) (int, int) {
	originalFileIndex := c.FindIDOfRightMostTakenSpaceBlock(*diskElements, freeSpaceIndex, previousTakenSpaceIndex)

	if originalFileIndex == -1 {
		return 0, originalFileIndex
	}

	newFileID := (*diskElements)[originalFileIndex]

	c.SwapBlocksFragmented(diskElements, freeSpaceIndex, originalFileIndex, newFileID)

	return freeSpaceIndex * *newFileID, originalFileIndex
}

func (c *CalculateDataPipe) GetDeltaChecksumBlocks(diskElements *[]*int, freeSpaceIndex, freeSpaceSize int) (delta, numOfFilledPlaces int) {
	originalFileLeftIndex, originalFileRightIndex, newFileID := c.FindRightMostFittingBlock(*diskElements, freeSpaceIndex, freeSpaceSize)

	if newFileID == nil {
		return 0, 0
	}

	c.SwapEntireBlock(diskElements, freeSpaceIndex, originalFileLeftIndex, originalFileRightIndex, *newFileID)

	checksumDelta := 0

	length := (originalFileRightIndex - originalFileLeftIndex) + 1

	for i := 0; i < length; i++ {
		checksumDelta += (freeSpaceIndex + i) * *newFileID
	}

	return checksumDelta, length
}

func (c *CalculateDataPipe) GetFreeSpaceSize(diskElements []*int, currentPosition int) int {
	size := 0

	for i := currentPosition; i < len(diskElements); i++ {
		if diskElements[i] == nil {
			size++
		} else {
			break
		}
	}

	return size
}

func (c *CalculateDataPipe) GetChecksum(elements []*int, blocks bool) int {
	checksum := 0

	modifiedDisk := make([]*int, len(elements))
	copy(modifiedDisk, elements)
	previousTakenSpaceIndex := len(elements) - 1

	for pos := 0; pos < len(elements); pos++ {
		fileID := elements[pos]
		presentModifiedFileID := modifiedDisk[pos]

		if presentModifiedFileID != nil {
			checksum += pos * *presentModifiedFileID
		} else if fileID == nil {
			if !blocks {
				delta, fileIndex := c.GetDeltaChecksumFragmented(&modifiedDisk, pos, previousTakenSpaceIndex)

				if delta == 0 {
					break
				}

				checksum += delta
				previousTakenSpaceIndex = fileIndex
			} else {
				freeSpaceSize := c.GetFreeSpaceSize(elements, pos)
				delta, numOfFilledPlaces := c.GetDeltaChecksumBlocks(&modifiedDisk, pos, freeSpaceSize)

				if delta != 0 {
					checksum += delta
					pos += numOfFilledPlaces - 1
				}
			}
		}
	}

	return checksum
}

func (c *CalculateDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	elements := input.DiskInfo

	input.Result = c.GetChecksum(elements, false)
	input.Result2 = c.GetChecksum(elements, true)
	return input
}
