package vulkan

/*
#cgo CFLAGS: -I./../../mod/Vulkan-Headers/include -I./../../mod/volk
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include <string.h>
#include "volk_wrappers.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Structure is the interface implemented by all Vulkan structs that carry
// an sType / pNext chain.
type Structure interface {
	GetType() StructureType
	toC() (unsafe.Pointer, func())
}

// Initialize loads the Vulkan library via Volk.
// Must be called before any other Vulkan function.
func Initialize() error {
	if result := C.volkInitialize(); result != C.VK_SUCCESS {
		return vkError(result)
	}
	return nil
}

func vkError(r C.VkResult) error {
	return fmt.Errorf("vulkan error: VkResult(%d)", int(r))
}

// Vulkan API constants
var (
	MaxPhysicalDeviceNameSize                                      = 256
	UuidSize                                                       = 16
	LuidSize                                                       = 8
	MaxExtensionNameSize                                           = 256
	MaxDescriptionSize                                             = 256
	MaxMemoryTypes                                                 = 32
	MaxMemoryHeaps                                                 = 16
	LodClampNone                                           float32 = 1000.0
	RemainingMipLevels                                     uint32  = ^uint32(0)
	RemainingArrayLayers                                   uint32  = ^uint32(0)
	Remaining3dSlicesEXT                                   uint32  = ^uint32(0)
	WholeSize                                              uint64  = ^uint64(0)
	AttachmentUnused                                       uint32  = ^uint32(0)
	True                                                           = 1
	False                                                          = 0
	QueueFamilyIgnored                                     uint32  = ^uint32(0)
	QueueFamilyExternal                                    uint32  = ^uint32(0) - 0
	QueueFamilyForeignEXT                                  uint32  = ^uint32(0) - 1
	SubpassExternal                                        uint32  = ^uint32(0)
	MaxDeviceGroupSize                                             = 32
	MaxDriverNameSize                                              = 256
	MaxDriverInfoSize                                              = 256
	ShaderUnusedKHR                                        uint32  = ^uint32(0)
	MaxGlobalPrioritySize                                          = 16
	MaxShaderModuleIdentifierSizeEXT                               = 32
	MaxPipelineBinaryKeySizeKHR                                    = 32
	MaxVideoAv1ReferencesPerFrameKHR                               = 7
	MaxVideoVp9ReferencesPerFrameKHR                               = 3
	ShaderIndexUnusedAmdx                                  uint32  = ^uint32(0)
	PartitionedAccelerationStructurePartitionIndexGlobalNV uint32  = ^uint32(0)
	CompressedTriangleFormatDgf1ByteAlignmentAmdx                  = 128
	CompressedTriangleFormatDgf1ByteStrideAmdx                     = 128
	MaxPhysicalDeviceDataGraphOperationSetNameSizeARM              = 128
	DataGraphModelToolchainVersionLengthQCOM                       = 3
	ComputeOccupancyPriorityLowNV                          float32 = 0.25
	ComputeOccupancyPriorityNormalNV                       float32 = 0.50
	ComputeOccupancyPriorityHighNV                         float32 = 0.75
)

type Buffer struct{ handle unsafe.Pointer }

type CommandBuffer struct{ handle unsafe.Pointer }

type Device struct{ handle unsafe.Pointer }

type Framebuffer struct{ handle unsafe.Pointer }

type Image struct{ handle unsafe.Pointer }

type ImageView struct{ handle unsafe.Pointer }

type Pipeline struct{ handle unsafe.Pointer }

type PipelineCache struct{ handle unsafe.Pointer }

type PipelineLayout struct{ handle unsafe.Pointer }

type RenderPass struct{ handle unsafe.Pointer }

type ShaderModule struct{ handle unsafe.Pointer }

type AllocationFunction func(userData unsafe.Pointer, size uintptr, alignment uintptr, allocationScope SystemAllocationScope) unsafe.Pointer

type FreeFunction func(userData unsafe.Pointer, memory unsafe.Pointer)

type InternalAllocationNotification func(userData unsafe.Pointer, size uintptr, allocationType InternalAllocationType, allocationScope SystemAllocationScope)

type InternalFreeNotification func(userData unsafe.Pointer, size uintptr, allocationType InternalAllocationType, allocationScope SystemAllocationScope)

type ReallocationFunction func(userData unsafe.Pointer, original unsafe.Pointer, size uintptr, alignment uintptr, allocationScope SystemAllocationScope) unsafe.Pointer

type AttachmentDescriptionFlagBits uint32

const (
	AttachmentDescriptionMayAliasBit                         AttachmentDescriptionFlagBits = 1 << 0
	AttachmentDescriptionResolveSkipTransferFunctionBitKHR   AttachmentDescriptionFlagBits = 1 << 1
	AttachmentDescriptionResolveEnableTransferFunctionBitKHR AttachmentDescriptionFlagBits = 1 << 2
)

type AttachmentLoadOp uint32

const (
	AttachmentLoadOpLoad     AttachmentLoadOp = 0
	AttachmentLoadOpClear    AttachmentLoadOp = 1
	AttachmentLoadOpDontCare AttachmentLoadOp = 2
	AttachmentLoadOpNone     AttachmentLoadOp = 1000400000
	AttachmentLoadOpNoneEXT  AttachmentLoadOp = AttachmentLoadOpNone
	AttachmentLoadOpNoneKHR  AttachmentLoadOp = AttachmentLoadOpNone
)

type AttachmentStoreOp uint32

const (
	AttachmentStoreOpStore    AttachmentStoreOp = 0
	AttachmentStoreOpDontCare AttachmentStoreOp = 1
	AttachmentStoreOpNone     AttachmentStoreOp = 1000301000
	AttachmentStoreOpNoneKHR  AttachmentStoreOp = AttachmentStoreOpNone
	AttachmentStoreOpNoneQCOM AttachmentStoreOp = AttachmentStoreOpNone
	AttachmentStoreOpNoneEXT  AttachmentStoreOp = AttachmentStoreOpNone
)

type BlendFactor uint32

const (
	BlendFactorZero                  BlendFactor = 0
	BlendFactorOne                   BlendFactor = 1
	BlendFactorSrcColor              BlendFactor = 2
	BlendFactorOneMinusSrcColor      BlendFactor = 3
	BlendFactorDstColor              BlendFactor = 4
	BlendFactorOneMinusDstColor      BlendFactor = 5
	BlendFactorSrcAlpha              BlendFactor = 6
	BlendFactorOneMinusSrcAlpha      BlendFactor = 7
	BlendFactorDstAlpha              BlendFactor = 8
	BlendFactorOneMinusDstAlpha      BlendFactor = 9
	BlendFactorConstantColor         BlendFactor = 10
	BlendFactorOneMinusConstantColor BlendFactor = 11
	BlendFactorConstantAlpha         BlendFactor = 12
	BlendFactorOneMinusConstantAlpha BlendFactor = 13
	BlendFactorSrcAlphaSaturate      BlendFactor = 14
	BlendFactorSrc1Color             BlendFactor = 15
	BlendFactorOneMinusSrc1Color     BlendFactor = 16
	BlendFactorSrc1Alpha             BlendFactor = 17
	BlendFactorOneMinusSrc1Alpha     BlendFactor = 18
)

type BlendOp uint32

const (
	BlendOpAdd                 BlendOp = 0
	BlendOpSubtract            BlendOp = 1
	BlendOpReverseSubtract     BlendOp = 2
	BlendOpMin                 BlendOp = 3
	BlendOpMax                 BlendOp = 4
	BlendOpZeroEXT             BlendOp = 1000148000
	BlendOpSrcEXT              BlendOp = 1000148001
	BlendOpDstEXT              BlendOp = 1000148002
	BlendOpSrcOverEXT          BlendOp = 1000148003
	BlendOpDstOverEXT          BlendOp = 1000148004
	BlendOpSrcInEXT            BlendOp = 1000148005
	BlendOpDstInEXT            BlendOp = 1000148006
	BlendOpSrcOutEXT           BlendOp = 1000148007
	BlendOpDstOutEXT           BlendOp = 1000148008
	BlendOpSrcAtopEXT          BlendOp = 1000148009
	BlendOpDstAtopEXT          BlendOp = 1000148010
	BlendOpXorEXT              BlendOp = 1000148011
	BlendOpMultiplyEXT         BlendOp = 1000148012
	BlendOpScreenEXT           BlendOp = 1000148013
	BlendOpOverlayEXT          BlendOp = 1000148014
	BlendOpDarkenEXT           BlendOp = 1000148015
	BlendOpLightenEXT          BlendOp = 1000148016
	BlendOpColordodgeEXT       BlendOp = 1000148017
	BlendOpColorburnEXT        BlendOp = 1000148018
	BlendOpHardlightEXT        BlendOp = 1000148019
	BlendOpSoftlightEXT        BlendOp = 1000148020
	BlendOpDifferenceEXT       BlendOp = 1000148021
	BlendOpExclusionEXT        BlendOp = 1000148022
	BlendOpInvertEXT           BlendOp = 1000148023
	BlendOpInvertRgbEXT        BlendOp = 1000148024
	BlendOpLineardodgeEXT      BlendOp = 1000148025
	BlendOpLinearburnEXT       BlendOp = 1000148026
	BlendOpVividlightEXT       BlendOp = 1000148027
	BlendOpLinearlightEXT      BlendOp = 1000148028
	BlendOpPinlightEXT         BlendOp = 1000148029
	BlendOpHardmixEXT          BlendOp = 1000148030
	BlendOpHslHueEXT           BlendOp = 1000148031
	BlendOpHslSaturationEXT    BlendOp = 1000148032
	BlendOpHslColorEXT         BlendOp = 1000148033
	BlendOpHslLuminosityEXT    BlendOp = 1000148034
	BlendOpPlusEXT             BlendOp = 1000148035
	BlendOpPlusClampedEXT      BlendOp = 1000148036
	BlendOpPlusClampedAlphaEXT BlendOp = 1000148037
	BlendOpPlusDarkerEXT       BlendOp = 1000148038
	BlendOpMinusEXT            BlendOp = 1000148039
	BlendOpMinusClampedEXT     BlendOp = 1000148040
	BlendOpContrastEXT         BlendOp = 1000148041
	BlendOpInvertOvgEXT        BlendOp = 1000148042
	BlendOpRedEXT              BlendOp = 1000148043
	BlendOpGreenEXT            BlendOp = 1000148044
	BlendOpBlueEXT             BlendOp = 1000148045
)

type ColorComponentFlagBits uint32

const (
	ColorComponentRBit ColorComponentFlagBits = 1 << 0
	ColorComponentGBit ColorComponentFlagBits = 1 << 1
	ColorComponentBBit ColorComponentFlagBits = 1 << 2
	ColorComponentABit ColorComponentFlagBits = 1 << 3
)

type CompareOp uint32

const (
	CompareOpNever          CompareOp = 0
	CompareOpLess           CompareOp = 1
	CompareOpEqual          CompareOp = 2
	CompareOpLessOrEqual    CompareOp = 3
	CompareOpGreater        CompareOp = 4
	CompareOpNotEqual       CompareOp = 5
	CompareOpGreaterOrEqual CompareOp = 6
	CompareOpAlways         CompareOp = 7
)

type CullModeFlagBits uint32

const (
	CullModeNone         CullModeFlagBits = 0
	CullModeFrontBit     CullModeFlagBits = 1 << 0
	CullModeBackBit      CullModeFlagBits = 1 << 1
	CullModeFrontAndBack CullModeFlagBits = 0x00000003
)

type DependencyFlagBits uint32

const (
	DependencyByRegionBit                                    DependencyFlagBits = 1 << 0
	DependencyDeviceGroupBit                                 DependencyFlagBits = 1 << 2
	DependencyViewLocalBit                                   DependencyFlagBits = 1 << 1
	DependencyViewLocalBitKHR                                DependencyFlagBits = DependencyViewLocalBit
	DependencyDeviceGroupBitKHR                              DependencyFlagBits = DependencyDeviceGroupBit
	DependencyFeedbackLoopBitEXT                             DependencyFlagBits = 1 << 3
	DependencyQueueFamilyOwnershipTransferUseAllStagesBitKHR DependencyFlagBits = 1 << 5
	DependencyAsymmetricEventBitKHR                          DependencyFlagBits = 1 << 6
	DependencyExtension586BitIMG                             DependencyFlagBits = 1 << 4
)

type DynamicState uint32

const (
	DynamicStateViewport                            DynamicState = 0
	DynamicStateScissor                             DynamicState = 1
	DynamicStateLineWidth                           DynamicState = 2
	DynamicStateDepthBias                           DynamicState = 3
	DynamicStateBlendConstants                      DynamicState = 4
	DynamicStateDepthBounds                         DynamicState = 5
	DynamicStateStencilCompareMask                  DynamicState = 6
	DynamicStateStencilWriteMask                    DynamicState = 7
	DynamicStateStencilReference                    DynamicState = 8
	DynamicStateCullMode                            DynamicState = 1000267000
	DynamicStateFrontFace                           DynamicState = 1000267001
	DynamicStatePrimitiveTopology                   DynamicState = 1000267002
	DynamicStateViewportWithCount                   DynamicState = 1000267003
	DynamicStateScissorWithCount                    DynamicState = 1000267004
	DynamicStateVertexInputBindingStride            DynamicState = 1000267005
	DynamicStateDepthTestEnable                     DynamicState = 1000267006
	DynamicStateDepthWriteEnable                    DynamicState = 1000267007
	DynamicStateDepthCompareOp                      DynamicState = 1000267008
	DynamicStateDepthBoundsTestEnable               DynamicState = 1000267009
	DynamicStateStencilTestEnable                   DynamicState = 1000267010
	DynamicStateStencilOp                           DynamicState = 1000267011
	DynamicStateRasterizerDiscardEnable             DynamicState = 1000377001
	DynamicStateDepthBiasEnable                     DynamicState = 1000377002
	DynamicStatePrimitiveRestartEnable              DynamicState = 1000377004
	DynamicStateLineStipple                         DynamicState = 1000259000
	DynamicStateViewportWScalingNV                  DynamicState = 1000087000
	DynamicStateDiscardRectangleEXT                 DynamicState = 1000099000
	DynamicStateDiscardRectangleEnableEXT           DynamicState = 1000099001
	DynamicStateDiscardRectangleModeEXT             DynamicState = 1000099002
	DynamicStateSampleLocationsEXT                  DynamicState = 1000143000
	DynamicStateRayTracingPipelineStackSizeKHR      DynamicState = 1000151000
	DynamicStateViewportShadingRatePaletteNV        DynamicState = 1000166004
	DynamicStateViewportCoarseSampleOrderNV         DynamicState = 1000166006
	DynamicStateExclusiveScissorEnableNV            DynamicState = 1000207000
	DynamicStateExclusiveScissorNV                  DynamicState = 1000207001
	DynamicStateFragmentShadingRateKHR              DynamicState = 1000228000
	DynamicStateLineStippleEXT                      DynamicState = DynamicStateLineStipple
	DynamicStateCullModeEXT                         DynamicState = DynamicStateCullMode
	DynamicStateFrontFaceEXT                        DynamicState = DynamicStateFrontFace
	DynamicStatePrimitiveTopologyEXT                DynamicState = DynamicStatePrimitiveTopology
	DynamicStateViewportWithCountEXT                DynamicState = DynamicStateViewportWithCount
	DynamicStateScissorWithCountEXT                 DynamicState = DynamicStateScissorWithCount
	DynamicStateVertexInputBindingStrideEXT         DynamicState = DynamicStateVertexInputBindingStride
	DynamicStateDepthTestEnableEXT                  DynamicState = DynamicStateDepthTestEnable
	DynamicStateDepthWriteEnableEXT                 DynamicState = DynamicStateDepthWriteEnable
	DynamicStateDepthCompareOpEXT                   DynamicState = DynamicStateDepthCompareOp
	DynamicStateDepthBoundsTestEnableEXT            DynamicState = DynamicStateDepthBoundsTestEnable
	DynamicStateStencilTestEnableEXT                DynamicState = DynamicStateStencilTestEnable
	DynamicStateStencilOpEXT                        DynamicState = DynamicStateStencilOp
	DynamicStateVertexInputEXT                      DynamicState = 1000352000
	DynamicStatePatchControlPointsEXT               DynamicState = 1000377000
	DynamicStateRasterizerDiscardEnableEXT          DynamicState = DynamicStateRasterizerDiscardEnable
	DynamicStateDepthBiasEnableEXT                  DynamicState = DynamicStateDepthBiasEnable
	DynamicStateLogicOpEXT                          DynamicState = 1000377003
	DynamicStatePrimitiveRestartEnableEXT           DynamicState = DynamicStatePrimitiveRestartEnable
	DynamicStateColorWriteEnableEXT                 DynamicState = 1000381000
	DynamicStateDepthClampEnableEXT                 DynamicState = 1000455003
	DynamicStatePolygonModeEXT                      DynamicState = 1000455004
	DynamicStateRasterizationSamplesEXT             DynamicState = 1000455005
	DynamicStateSampleMaskEXT                       DynamicState = 1000455006
	DynamicStateAlphaToCoverageEnableEXT            DynamicState = 1000455007
	DynamicStateAlphaToOneEnableEXT                 DynamicState = 1000455008
	DynamicStateLogicOpEnableEXT                    DynamicState = 1000455009
	DynamicStateColorBlendEnableEXT                 DynamicState = 1000455010
	DynamicStateColorBlendEquationEXT               DynamicState = 1000455011
	DynamicStateColorWriteMaskEXT                   DynamicState = 1000455012
	DynamicStateTessellationDomainOriginEXT         DynamicState = 1000455002
	DynamicStateRasterizationStreamEXT              DynamicState = 1000455013
	DynamicStateConservativeRasterizationModeEXT    DynamicState = 1000455014
	DynamicStateExtraPrimitiveOverestimationSizeEXT DynamicState = 1000455015
	DynamicStateDepthClipEnableEXT                  DynamicState = 1000455016
	DynamicStateSampleLocationsEnableEXT            DynamicState = 1000455017
	DynamicStateColorBlendAdvancedEXT               DynamicState = 1000455018
	DynamicStateProvokingVertexModeEXT              DynamicState = 1000455019
	DynamicStateLineRasterizationModeEXT            DynamicState = 1000455020
	DynamicStateLineStippleEnableEXT                DynamicState = 1000455021
	DynamicStateDepthClipNegativeOneToOneEXT        DynamicState = 1000455022
	DynamicStateViewportWScalingEnableNV            DynamicState = 1000455023
	DynamicStateViewportSwizzleNV                   DynamicState = 1000455024
	DynamicStateCoverageToColorEnableNV             DynamicState = 1000455025
	DynamicStateCoverageToColorLocationNV           DynamicState = 1000455026
	DynamicStateCoverageModulationModeNV            DynamicState = 1000455027
	DynamicStateCoverageModulationTableEnableNV     DynamicState = 1000455028
	DynamicStateCoverageModulationTableNV           DynamicState = 1000455029
	DynamicStateShadingRateImageEnableNV            DynamicState = 1000455030
	DynamicStateRepresentativeFragmentTestEnableNV  DynamicState = 1000455031
	DynamicStateCoverageReductionModeNV             DynamicState = 1000455032
	DynamicStateAttachmentFeedbackLoopEnableEXT     DynamicState = 1000524000
	DynamicStateLineStippleKHR                      DynamicState = DynamicStateLineStipple
	DynamicStateDepthClampRangeEXT                  DynamicState = 1000582000
)

type Filter uint32

const (
	FilterNearest  Filter = 0
	FilterLinear   Filter = 1
	FilterCubicIMG Filter = FilterCubicEXT
	FilterCubicEXT Filter = 1000015000
)

type Format uint32

const (
	FormatUndefined                               Format = 0
	FormatR4g4UnormPack8                          Format = 1
	FormatR4g4b4a4UnormPack16                     Format = 2
	FormatB4g4r4a4UnormPack16                     Format = 3
	FormatR5g6b5UnormPack16                       Format = 4
	FormatB5g6r5UnormPack16                       Format = 5
	FormatR5g5b5a1UnormPack16                     Format = 6
	FormatB5g5r5a1UnormPack16                     Format = 7
	FormatA1r5g5b5UnormPack16                     Format = 8
	FormatR8Unorm                                 Format = 9
	FormatR8Snorm                                 Format = 10
	FormatR8Uscaled                               Format = 11
	FormatR8Sscaled                               Format = 12
	FormatR8Uint                                  Format = 13
	FormatR8Sint                                  Format = 14
	FormatR8Srgb                                  Format = 15
	FormatR8g8Unorm                               Format = 16
	FormatR8g8Snorm                               Format = 17
	FormatR8g8Uscaled                             Format = 18
	FormatR8g8Sscaled                             Format = 19
	FormatR8g8Uint                                Format = 20
	FormatR8g8Sint                                Format = 21
	FormatR8g8Srgb                                Format = 22
	FormatR8g8b8Unorm                             Format = 23
	FormatR8g8b8Snorm                             Format = 24
	FormatR8g8b8Uscaled                           Format = 25
	FormatR8g8b8Sscaled                           Format = 26
	FormatR8g8b8Uint                              Format = 27
	FormatR8g8b8Sint                              Format = 28
	FormatR8g8b8Srgb                              Format = 29
	FormatB8g8r8Unorm                             Format = 30
	FormatB8g8r8Snorm                             Format = 31
	FormatB8g8r8Uscaled                           Format = 32
	FormatB8g8r8Sscaled                           Format = 33
	FormatB8g8r8Uint                              Format = 34
	FormatB8g8r8Sint                              Format = 35
	FormatB8g8r8Srgb                              Format = 36
	FormatR8g8b8a8Unorm                           Format = 37
	FormatR8g8b8a8Snorm                           Format = 38
	FormatR8g8b8a8Uscaled                         Format = 39
	FormatR8g8b8a8Sscaled                         Format = 40
	FormatR8g8b8a8Uint                            Format = 41
	FormatR8g8b8a8Sint                            Format = 42
	FormatR8g8b8a8Srgb                            Format = 43
	FormatB8g8r8a8Unorm                           Format = 44
	FormatB8g8r8a8Snorm                           Format = 45
	FormatB8g8r8a8Uscaled                         Format = 46
	FormatB8g8r8a8Sscaled                         Format = 47
	FormatB8g8r8a8Uint                            Format = 48
	FormatB8g8r8a8Sint                            Format = 49
	FormatB8g8r8a8Srgb                            Format = 50
	FormatA8b8g8r8UnormPack32                     Format = 51
	FormatA8b8g8r8SnormPack32                     Format = 52
	FormatA8b8g8r8UscaledPack32                   Format = 53
	FormatA8b8g8r8SscaledPack32                   Format = 54
	FormatA8b8g8r8UintPack32                      Format = 55
	FormatA8b8g8r8SintPack32                      Format = 56
	FormatA8b8g8r8SrgbPack32                      Format = 57
	FormatA2r10g10b10UnormPack32                  Format = 58
	FormatA2r10g10b10SnormPack32                  Format = 59
	FormatA2r10g10b10UscaledPack32                Format = 60
	FormatA2r10g10b10SscaledPack32                Format = 61
	FormatA2r10g10b10UintPack32                   Format = 62
	FormatA2r10g10b10SintPack32                   Format = 63
	FormatA2b10g10r10UnormPack32                  Format = 64
	FormatA2b10g10r10SnormPack32                  Format = 65
	FormatA2b10g10r10UscaledPack32                Format = 66
	FormatA2b10g10r10SscaledPack32                Format = 67
	FormatA2b10g10r10UintPack32                   Format = 68
	FormatA2b10g10r10SintPack32                   Format = 69
	FormatR16Unorm                                Format = 70
	FormatR16Snorm                                Format = 71
	FormatR16Uscaled                              Format = 72
	FormatR16Sscaled                              Format = 73
	FormatR16Uint                                 Format = 74
	FormatR16Sint                                 Format = 75
	FormatR16Sfloat                               Format = 76
	FormatR16g16Unorm                             Format = 77
	FormatR16g16Snorm                             Format = 78
	FormatR16g16Uscaled                           Format = 79
	FormatR16g16Sscaled                           Format = 80
	FormatR16g16Uint                              Format = 81
	FormatR16g16Sint                              Format = 82
	FormatR16g16Sfloat                            Format = 83
	FormatR16g16b16Unorm                          Format = 84
	FormatR16g16b16Snorm                          Format = 85
	FormatR16g16b16Uscaled                        Format = 86
	FormatR16g16b16Sscaled                        Format = 87
	FormatR16g16b16Uint                           Format = 88
	FormatR16g16b16Sint                           Format = 89
	FormatR16g16b16Sfloat                         Format = 90
	FormatR16g16b16a16Unorm                       Format = 91
	FormatR16g16b16a16Snorm                       Format = 92
	FormatR16g16b16a16Uscaled                     Format = 93
	FormatR16g16b16a16Sscaled                     Format = 94
	FormatR16g16b16a16Uint                        Format = 95
	FormatR16g16b16a16Sint                        Format = 96
	FormatR16g16b16a16Sfloat                      Format = 97
	FormatR32Uint                                 Format = 98
	FormatR32Sint                                 Format = 99
	FormatR32Sfloat                               Format = 100
	FormatR32g32Uint                              Format = 101
	FormatR32g32Sint                              Format = 102
	FormatR32g32Sfloat                            Format = 103
	FormatR32g32b32Uint                           Format = 104
	FormatR32g32b32Sint                           Format = 105
	FormatR32g32b32Sfloat                         Format = 106
	FormatR32g32b32a32Uint                        Format = 107
	FormatR32g32b32a32Sint                        Format = 108
	FormatR32g32b32a32Sfloat                      Format = 109
	FormatR64Uint                                 Format = 110
	FormatR64Sint                                 Format = 111
	FormatR64Sfloat                               Format = 112
	FormatR64g64Uint                              Format = 113
	FormatR64g64Sint                              Format = 114
	FormatR64g64Sfloat                            Format = 115
	FormatR64g64b64Uint                           Format = 116
	FormatR64g64b64Sint                           Format = 117
	FormatR64g64b64Sfloat                         Format = 118
	FormatR64g64b64a64Uint                        Format = 119
	FormatR64g64b64a64Sint                        Format = 120
	FormatR64g64b64a64Sfloat                      Format = 121
	FormatB10g11r11UfloatPack32                   Format = 122
	FormatE5b9g9r9UfloatPack32                    Format = 123
	FormatD16Unorm                                Format = 124
	FormatX8D24UnormPack32                        Format = 125
	FormatD32Sfloat                               Format = 126
	FormatS8Uint                                  Format = 127
	FormatD16UnormS8Uint                          Format = 128
	FormatD24UnormS8Uint                          Format = 129
	FormatD32SfloatS8Uint                         Format = 130
	FormatBc1RgbUnormBlock                        Format = 131
	FormatBc1RgbSrgbBlock                         Format = 132
	FormatBc1RgbaUnormBlock                       Format = 133
	FormatBc1RgbaSrgbBlock                        Format = 134
	FormatBc2UnormBlock                           Format = 135
	FormatBc2SrgbBlock                            Format = 136
	FormatBc3UnormBlock                           Format = 137
	FormatBc3SrgbBlock                            Format = 138
	FormatBc4UnormBlock                           Format = 139
	FormatBc4SnormBlock                           Format = 140
	FormatBc5UnormBlock                           Format = 141
	FormatBc5SnormBlock                           Format = 142
	FormatBc6hUfloatBlock                         Format = 143
	FormatBc6hSfloatBlock                         Format = 144
	FormatBc7UnormBlock                           Format = 145
	FormatBc7SrgbBlock                            Format = 146
	FormatEtc2R8g8b8UnormBlock                    Format = 147
	FormatEtc2R8g8b8SrgbBlock                     Format = 148
	FormatEtc2R8g8b8a1UnormBlock                  Format = 149
	FormatEtc2R8g8b8a1SrgbBlock                   Format = 150
	FormatEtc2R8g8b8a8UnormBlock                  Format = 151
	FormatEtc2R8g8b8a8SrgbBlock                   Format = 152
	FormatEacR11UnormBlock                        Format = 153
	FormatEacR11SnormBlock                        Format = 154
	FormatEacR11g11UnormBlock                     Format = 155
	FormatEacR11g11SnormBlock                     Format = 156
	FormatAstc4x4UnormBlock                       Format = 157
	FormatAstc4x4SrgbBlock                        Format = 158
	FormatAstc5x4UnormBlock                       Format = 159
	FormatAstc5x4SrgbBlock                        Format = 160
	FormatAstc5x5UnormBlock                       Format = 161
	FormatAstc5x5SrgbBlock                        Format = 162
	FormatAstc6x5UnormBlock                       Format = 163
	FormatAstc6x5SrgbBlock                        Format = 164
	FormatAstc6x6UnormBlock                       Format = 165
	FormatAstc6x6SrgbBlock                        Format = 166
	FormatAstc8x5UnormBlock                       Format = 167
	FormatAstc8x5SrgbBlock                        Format = 168
	FormatAstc8x6UnormBlock                       Format = 169
	FormatAstc8x6SrgbBlock                        Format = 170
	FormatAstc8x8UnormBlock                       Format = 171
	FormatAstc8x8SrgbBlock                        Format = 172
	FormatAstc10x5UnormBlock                      Format = 173
	FormatAstc10x5SrgbBlock                       Format = 174
	FormatAstc10x6UnormBlock                      Format = 175
	FormatAstc10x6SrgbBlock                       Format = 176
	FormatAstc10x8UnormBlock                      Format = 177
	FormatAstc10x8SrgbBlock                       Format = 178
	FormatAstc10x10UnormBlock                     Format = 179
	FormatAstc10x10SrgbBlock                      Format = 180
	FormatAstc12x10UnormBlock                     Format = 181
	FormatAstc12x10SrgbBlock                      Format = 182
	FormatAstc12x12UnormBlock                     Format = 183
	FormatAstc12x12SrgbBlock                      Format = 184
	FormatG8b8g8r8422Unorm                        Format = 1000156000
	FormatB8g8r8g8422Unorm                        Format = 1000156001
	FormatG8B8R83plane420Unorm                    Format = 1000156002
	FormatG8B8r82plane420Unorm                    Format = 1000156003
	FormatG8B8R83plane422Unorm                    Format = 1000156004
	FormatG8B8r82plane422Unorm                    Format = 1000156005
	FormatG8B8R83plane444Unorm                    Format = 1000156006
	FormatR10x6UnormPack16                        Format = 1000156007
	FormatR10x6g10x6Unorm2pack16                  Format = 1000156008
	FormatR10x6g10x6b10x6a10x6Unorm4pack16        Format = 1000156009
	FormatG10x6b10x6g10x6r10x6422Unorm4pack16     Format = 1000156010
	FormatB10x6g10x6r10x6g10x6422Unorm4pack16     Format = 1000156011
	FormatG10x6B10x6R10x63plane420Unorm3pack16    Format = 1000156012
	FormatG10x6B10x6r10x62plane420Unorm3pack16    Format = 1000156013
	FormatG10x6B10x6R10x63plane422Unorm3pack16    Format = 1000156014
	FormatG10x6B10x6r10x62plane422Unorm3pack16    Format = 1000156015
	FormatG10x6B10x6R10x63plane444Unorm3pack16    Format = 1000156016
	FormatR12x4UnormPack16                        Format = 1000156017
	FormatR12x4g12x4Unorm2pack16                  Format = 1000156018
	FormatR12x4g12x4b12x4a12x4Unorm4pack16        Format = 1000156019
	FormatG12x4b12x4g12x4r12x4422Unorm4pack16     Format = 1000156020
	FormatB12x4g12x4r12x4g12x4422Unorm4pack16     Format = 1000156021
	FormatG12x4B12x4R12x43plane420Unorm3pack16    Format = 1000156022
	FormatG12x4B12x4r12x42plane420Unorm3pack16    Format = 1000156023
	FormatG12x4B12x4R12x43plane422Unorm3pack16    Format = 1000156024
	FormatG12x4B12x4r12x42plane422Unorm3pack16    Format = 1000156025
	FormatG12x4B12x4R12x43plane444Unorm3pack16    Format = 1000156026
	FormatG16b16g16r16422Unorm                    Format = 1000156027
	FormatB16g16r16g16422Unorm                    Format = 1000156028
	FormatG16B16R163plane420Unorm                 Format = 1000156029
	FormatG16B16r162plane420Unorm                 Format = 1000156030
	FormatG16B16R163plane422Unorm                 Format = 1000156031
	FormatG16B16r162plane422Unorm                 Format = 1000156032
	FormatG16B16R163plane444Unorm                 Format = 1000156033
	FormatG8B8r82plane444Unorm                    Format = 1000330000
	FormatG10x6B10x6r10x62plane444Unorm3pack16    Format = 1000330001
	FormatG12x4B12x4r12x42plane444Unorm3pack16    Format = 1000330002
	FormatG16B16r162plane444Unorm                 Format = 1000330003
	FormatA4r4g4b4UnormPack16                     Format = 1000340000
	FormatA4b4g4r4UnormPack16                     Format = 1000340001
	FormatAstc4x4SfloatBlock                      Format = 1000066000
	FormatAstc5x4SfloatBlock                      Format = 1000066001
	FormatAstc5x5SfloatBlock                      Format = 1000066002
	FormatAstc6x5SfloatBlock                      Format = 1000066003
	FormatAstc6x6SfloatBlock                      Format = 1000066004
	FormatAstc8x5SfloatBlock                      Format = 1000066005
	FormatAstc8x6SfloatBlock                      Format = 1000066006
	FormatAstc8x8SfloatBlock                      Format = 1000066007
	FormatAstc10x5SfloatBlock                     Format = 1000066008
	FormatAstc10x6SfloatBlock                     Format = 1000066009
	FormatAstc10x8SfloatBlock                     Format = 1000066010
	FormatAstc10x10SfloatBlock                    Format = 1000066011
	FormatAstc12x10SfloatBlock                    Format = 1000066012
	FormatAstc12x12SfloatBlock                    Format = 1000066013
	FormatA1b5g5r5UnormPack16                     Format = 1000470000
	FormatA8Unorm                                 Format = 1000470001
	FormatPvrtc12bppUnormBlockIMG                 Format = 1000054000
	FormatPvrtc14bppUnormBlockIMG                 Format = 1000054001
	FormatPvrtc22bppUnormBlockIMG                 Format = 1000054002
	FormatPvrtc24bppUnormBlockIMG                 Format = 1000054003
	FormatPvrtc12bppSrgbBlockIMG                  Format = 1000054004
	FormatPvrtc14bppSrgbBlockIMG                  Format = 1000054005
	FormatPvrtc22bppSrgbBlockIMG                  Format = 1000054006
	FormatPvrtc24bppSrgbBlockIMG                  Format = 1000054007
	FormatAstc4x4SfloatBlockEXT                   Format = FormatAstc4x4SfloatBlock
	FormatAstc5x4SfloatBlockEXT                   Format = FormatAstc5x4SfloatBlock
	FormatAstc5x5SfloatBlockEXT                   Format = FormatAstc5x5SfloatBlock
	FormatAstc6x5SfloatBlockEXT                   Format = FormatAstc6x5SfloatBlock
	FormatAstc6x6SfloatBlockEXT                   Format = FormatAstc6x6SfloatBlock
	FormatAstc8x5SfloatBlockEXT                   Format = FormatAstc8x5SfloatBlock
	FormatAstc8x6SfloatBlockEXT                   Format = FormatAstc8x6SfloatBlock
	FormatAstc8x8SfloatBlockEXT                   Format = FormatAstc8x8SfloatBlock
	FormatAstc10x5SfloatBlockEXT                  Format = FormatAstc10x5SfloatBlock
	FormatAstc10x6SfloatBlockEXT                  Format = FormatAstc10x6SfloatBlock
	FormatAstc10x8SfloatBlockEXT                  Format = FormatAstc10x8SfloatBlock
	FormatAstc10x10SfloatBlockEXT                 Format = FormatAstc10x10SfloatBlock
	FormatAstc12x10SfloatBlockEXT                 Format = FormatAstc12x10SfloatBlock
	FormatAstc12x12SfloatBlockEXT                 Format = FormatAstc12x12SfloatBlock
	FormatG8b8g8r8422UnormKHR                     Format = FormatG8b8g8r8422Unorm
	FormatB8g8r8g8422UnormKHR                     Format = FormatB8g8r8g8422Unorm
	FormatG8B8R83plane420UnormKHR                 Format = FormatG8B8R83plane420Unorm
	FormatG8B8r82plane420UnormKHR                 Format = FormatG8B8r82plane420Unorm
	FormatG8B8R83plane422UnormKHR                 Format = FormatG8B8R83plane422Unorm
	FormatG8B8r82plane422UnormKHR                 Format = FormatG8B8r82plane422Unorm
	FormatG8B8R83plane444UnormKHR                 Format = FormatG8B8R83plane444Unorm
	FormatR10x6UnormPack16KHR                     Format = FormatR10x6UnormPack16
	FormatR10x6g10x6Unorm2pack16KHR               Format = FormatR10x6g10x6Unorm2pack16
	FormatR10x6g10x6b10x6a10x6Unorm4pack16KHR     Format = FormatR10x6g10x6b10x6a10x6Unorm4pack16
	FormatG10x6b10x6g10x6r10x6422Unorm4pack16KHR  Format = FormatG10x6b10x6g10x6r10x6422Unorm4pack16
	FormatB10x6g10x6r10x6g10x6422Unorm4pack16KHR  Format = FormatB10x6g10x6r10x6g10x6422Unorm4pack16
	FormatG10x6B10x6R10x63plane420Unorm3pack16KHR Format = FormatG10x6B10x6R10x63plane420Unorm3pack16
	FormatG10x6B10x6r10x62plane420Unorm3pack16KHR Format = FormatG10x6B10x6r10x62plane420Unorm3pack16
	FormatG10x6B10x6R10x63plane422Unorm3pack16KHR Format = FormatG10x6B10x6R10x63plane422Unorm3pack16
	FormatG10x6B10x6r10x62plane422Unorm3pack16KHR Format = FormatG10x6B10x6r10x62plane422Unorm3pack16
	FormatG10x6B10x6R10x63plane444Unorm3pack16KHR Format = FormatG10x6B10x6R10x63plane444Unorm3pack16
	FormatR12x4UnormPack16KHR                     Format = FormatR12x4UnormPack16
	FormatR12x4g12x4Unorm2pack16KHR               Format = FormatR12x4g12x4Unorm2pack16
	FormatR12x4g12x4b12x4a12x4Unorm4pack16KHR     Format = FormatR12x4g12x4b12x4a12x4Unorm4pack16
	FormatG12x4b12x4g12x4r12x4422Unorm4pack16KHR  Format = FormatG12x4b12x4g12x4r12x4422Unorm4pack16
	FormatB12x4g12x4r12x4g12x4422Unorm4pack16KHR  Format = FormatB12x4g12x4r12x4g12x4422Unorm4pack16
	FormatG12x4B12x4R12x43plane420Unorm3pack16KHR Format = FormatG12x4B12x4R12x43plane420Unorm3pack16
	FormatG12x4B12x4r12x42plane420Unorm3pack16KHR Format = FormatG12x4B12x4r12x42plane420Unorm3pack16
	FormatG12x4B12x4R12x43plane422Unorm3pack16KHR Format = FormatG12x4B12x4R12x43plane422Unorm3pack16
	FormatG12x4B12x4r12x42plane422Unorm3pack16KHR Format = FormatG12x4B12x4r12x42plane422Unorm3pack16
	FormatG12x4B12x4R12x43plane444Unorm3pack16KHR Format = FormatG12x4B12x4R12x43plane444Unorm3pack16
	FormatG16b16g16r16422UnormKHR                 Format = FormatG16b16g16r16422Unorm
	FormatB16g16r16g16422UnormKHR                 Format = FormatB16g16r16g16422Unorm
	FormatG16B16R163plane420UnormKHR              Format = FormatG16B16R163plane420Unorm
	FormatG16B16r162plane420UnormKHR              Format = FormatG16B16r162plane420Unorm
	FormatG16B16R163plane422UnormKHR              Format = FormatG16B16R163plane422Unorm
	FormatG16B16r162plane422UnormKHR              Format = FormatG16B16r162plane422Unorm
	FormatG16B16R163plane444UnormKHR              Format = FormatG16B16R163plane444Unorm
	FormatAstc3x3x3UnormBlockEXT                  Format = 1000288000
	FormatAstc3x3x3SrgbBlockEXT                   Format = 1000288001
	FormatAstc3x3x3SfloatBlockEXT                 Format = 1000288002
	FormatAstc4x3x3UnormBlockEXT                  Format = 1000288003
	FormatAstc4x3x3SrgbBlockEXT                   Format = 1000288004
	FormatAstc4x3x3SfloatBlockEXT                 Format = 1000288005
	FormatAstc4x4x3UnormBlockEXT                  Format = 1000288006
	FormatAstc4x4x3SrgbBlockEXT                   Format = 1000288007
	FormatAstc4x4x3SfloatBlockEXT                 Format = 1000288008
	FormatAstc4x4x4UnormBlockEXT                  Format = 1000288009
	FormatAstc4x4x4SrgbBlockEXT                   Format = 1000288010
	FormatAstc4x4x4SfloatBlockEXT                 Format = 1000288011
	FormatAstc5x4x4UnormBlockEXT                  Format = 1000288012
	FormatAstc5x4x4SrgbBlockEXT                   Format = 1000288013
	FormatAstc5x4x4SfloatBlockEXT                 Format = 1000288014
	FormatAstc5x5x4UnormBlockEXT                  Format = 1000288015
	FormatAstc5x5x4SrgbBlockEXT                   Format = 1000288016
	FormatAstc5x5x4SfloatBlockEXT                 Format = 1000288017
	FormatAstc5x5x5UnormBlockEXT                  Format = 1000288018
	FormatAstc5x5x5SrgbBlockEXT                   Format = 1000288019
	FormatAstc5x5x5SfloatBlockEXT                 Format = 1000288020
	FormatAstc6x5x5UnormBlockEXT                  Format = 1000288021
	FormatAstc6x5x5SrgbBlockEXT                   Format = 1000288022
	FormatAstc6x5x5SfloatBlockEXT                 Format = 1000288023
	FormatAstc6x6x5UnormBlockEXT                  Format = 1000288024
	FormatAstc6x6x5SrgbBlockEXT                   Format = 1000288025
	FormatAstc6x6x5SfloatBlockEXT                 Format = 1000288026
	FormatAstc6x6x6UnormBlockEXT                  Format = 1000288027
	FormatAstc6x6x6SrgbBlockEXT                   Format = 1000288028
	FormatAstc6x6x6SfloatBlockEXT                 Format = 1000288029
	FormatG8B8r82plane444UnormEXT                 Format = FormatG8B8r82plane444Unorm
	FormatG10x6B10x6r10x62plane444Unorm3pack16EXT Format = FormatG10x6B10x6r10x62plane444Unorm3pack16
	FormatG12x4B12x4r12x42plane444Unorm3pack16EXT Format = FormatG12x4B12x4r12x42plane444Unorm3pack16
	FormatG16B16r162plane444UnormEXT              Format = FormatG16B16r162plane444Unorm
	FormatA4r4g4b4UnormPack16EXT                  Format = FormatA4r4g4b4UnormPack16
	FormatA4b4g4r4UnormPack16EXT                  Format = FormatA4b4g4r4UnormPack16
	FormatR8BoolARM                               Format = 1000460000
	FormatR16SfloatFpencodingBfloat16ARM          Format = 1000460001
	FormatR8SfloatFpencodingFloat8e4m3ARM         Format = 1000460002
	FormatR8SfloatFpencodingFloat8e5m2ARM         Format = 1000460003
	FormatR16g16Sfixed5NV                         Format = 1000464000
	FormatR16g16S105NV                            Format = FormatR16g16Sfixed5NV
	FormatA1b5g5r5UnormPack16KHR                  Format = FormatA1b5g5r5UnormPack16
	FormatA8UnormKHR                              Format = FormatA8Unorm
	FormatR10x6UintPack16ARM                      Format = 1000610000
	FormatR10x6g10x6Uint2pack16ARM                Format = 1000610001
	FormatR10x6g10x6b10x6a10x6Uint4pack16ARM      Format = 1000610002
	FormatR12x4UintPack16ARM                      Format = 1000610003
	FormatR12x4g12x4Uint2pack16ARM                Format = 1000610004
	FormatR12x4g12x4b12x4a12x4Uint4pack16ARM      Format = 1000610005
	FormatR14x2UintPack16ARM                      Format = 1000610006
	FormatR14x2g14x2Uint2pack16ARM                Format = 1000610007
	FormatR14x2g14x2b14x2a14x2Uint4pack16ARM      Format = 1000610008
	FormatR14x2UnormPack16ARM                     Format = 1000610009
	FormatR14x2g14x2Unorm2pack16ARM               Format = 1000610010
	FormatR14x2g14x2b14x2a14x2Unorm4pack16ARM     Format = 1000610011
	FormatG14x2B14x2r14x22plane420Unorm3pack16ARM Format = 1000610012
	FormatG14x2B14x2r14x22plane422Unorm3pack16ARM Format = 1000610013
)

type FramebufferCreateFlagBits uint32

const (
	FramebufferCreateImagelessBit    FramebufferCreateFlagBits = 1 << 0
	FramebufferCreateImagelessBitKHR FramebufferCreateFlagBits = FramebufferCreateImagelessBit
)

type FrontFace uint32

const (
	FrontFaceCounterClockwise FrontFace = 0
	FrontFaceClockwise        FrontFace = 1
)

type ImageLayout uint32

const (
	ImageLayoutUndefined                                ImageLayout = 0
	ImageLayoutGeneral                                  ImageLayout = 1
	ImageLayoutColorAttachmentOptimal                   ImageLayout = 2
	ImageLayoutDepthStencilAttachmentOptimal            ImageLayout = 3
	ImageLayoutDepthStencilReadOnlyOptimal              ImageLayout = 4
	ImageLayoutShaderReadOnlyOptimal                    ImageLayout = 5
	ImageLayoutTransferSrcOptimal                       ImageLayout = 6
	ImageLayoutTransferDstOptimal                       ImageLayout = 7
	ImageLayoutPreinitialized                           ImageLayout = 8
	ImageLayoutDepthReadOnlyStencilAttachmentOptimal    ImageLayout = 1000117000
	ImageLayoutDepthAttachmentStencilReadOnlyOptimal    ImageLayout = 1000117001
	ImageLayoutDepthAttachmentOptimal                   ImageLayout = 1000241000
	ImageLayoutDepthReadOnlyOptimal                     ImageLayout = 1000241001
	ImageLayoutStencilAttachmentOptimal                 ImageLayout = 1000241002
	ImageLayoutStencilReadOnlyOptimal                   ImageLayout = 1000241003
	ImageLayoutReadOnlyOptimal                          ImageLayout = 1000314000
	ImageLayoutAttachmentOptimal                        ImageLayout = 1000314001
	ImageLayoutRenderingLocalRead                       ImageLayout = 1000232000
	ImageLayoutPresentSrcKHR                            ImageLayout = 1000001002
	ImageLayoutVideoDecodeDstKHR                        ImageLayout = 1000024000
	ImageLayoutVideoDecodeSrcKHR                        ImageLayout = 1000024001
	ImageLayoutVideoDecodeDpbKHR                        ImageLayout = 1000024002
	ImageLayoutSharedPresentKHR                         ImageLayout = 1000111000
	ImageLayoutDepthReadOnlyStencilAttachmentOptimalKHR ImageLayout = ImageLayoutDepthReadOnlyStencilAttachmentOptimal
	ImageLayoutDepthAttachmentStencilReadOnlyOptimalKHR ImageLayout = ImageLayoutDepthAttachmentStencilReadOnlyOptimal
	ImageLayoutShadingRateOptimalNV                     ImageLayout = ImageLayoutFragmentShadingRateAttachmentOptimalKHR
	ImageLayoutFragmentDensityMapOptimalEXT             ImageLayout = 1000220000
	ImageLayoutFragmentShadingRateAttachmentOptimalKHR  ImageLayout = 1000164003
	ImageLayoutRenderingLocalReadKHR                    ImageLayout = ImageLayoutRenderingLocalRead
	ImageLayoutDepthAttachmentOptimalKHR                ImageLayout = ImageLayoutDepthAttachmentOptimal
	ImageLayoutDepthReadOnlyOptimalKHR                  ImageLayout = ImageLayoutDepthReadOnlyOptimal
	ImageLayoutStencilAttachmentOptimalKHR              ImageLayout = ImageLayoutStencilAttachmentOptimal
	ImageLayoutStencilReadOnlyOptimalKHR                ImageLayout = ImageLayoutStencilReadOnlyOptimal
	ImageLayoutVideoEncodeDstKHR                        ImageLayout = 1000301000
	ImageLayoutVideoEncodeSrcKHR                        ImageLayout = 1000301001
	ImageLayoutVideoEncodeDpbKHR                        ImageLayout = 1000301002
	ImageLayoutReadOnlyOptimalKHR                       ImageLayout = ImageLayoutReadOnlyOptimal
	ImageLayoutAttachmentOptimalKHR                     ImageLayout = ImageLayoutAttachmentOptimal
	ImageLayoutAttachmentFeedbackLoopOptimalEXT         ImageLayout = 1000341000
	ImageLayoutTensorAliasingARM                        ImageLayout = 1000460000
	ImageLayoutVideoEncodeQuantizationMapKHR            ImageLayout = 1000553000
	ImageLayoutZeroInitializedEXT                       ImageLayout = 1000620000
)

type IndexType uint32

const (
	IndexTypeUint16   IndexType = 0
	IndexTypeUint32   IndexType = 1
	IndexTypeUint8    IndexType = 1000265000
	IndexTypeNoneKHR  IndexType = 1000165000
	IndexTypeNoneNV   IndexType = IndexTypeNoneKHR
	IndexTypeUint8EXT IndexType = IndexTypeUint8
	IndexTypeUint8KHR IndexType = IndexTypeUint8
)

type InternalAllocationType uint32

const (
	InternalAllocationTypeExecutable InternalAllocationType = 0
)

type LineRasterizationMode uint32

const (
	LineRasterizationModeDefault              LineRasterizationMode = 0
	LineRasterizationModeRectangular          LineRasterizationMode = 1
	LineRasterizationModeBresenham            LineRasterizationMode = 2
	LineRasterizationModeRectangularSmooth    LineRasterizationMode = 3
	LineRasterizationModeDefaultEXT           LineRasterizationMode = LineRasterizationModeDefault
	LineRasterizationModeRectangularEXT       LineRasterizationMode = LineRasterizationModeRectangular
	LineRasterizationModeBresenhamEXT         LineRasterizationMode = LineRasterizationModeBresenham
	LineRasterizationModeRectangularSmoothEXT LineRasterizationMode = LineRasterizationModeRectangularSmooth
	LineRasterizationModeDefaultKHR           LineRasterizationMode = LineRasterizationModeDefault
	LineRasterizationModeRectangularKHR       LineRasterizationMode = LineRasterizationModeRectangular
	LineRasterizationModeBresenhamKHR         LineRasterizationMode = LineRasterizationModeBresenham
	LineRasterizationModeRectangularSmoothKHR LineRasterizationMode = LineRasterizationModeRectangularSmooth
)

type LogicOp uint32

const (
	LogicOpClear        LogicOp = 0
	LogicOpAnd          LogicOp = 1
	LogicOpAndReverse   LogicOp = 2
	LogicOpCopy         LogicOp = 3
	LogicOpAndInverted  LogicOp = 4
	LogicOpNoOp         LogicOp = 5
	LogicOpXor          LogicOp = 6
	LogicOpOr           LogicOp = 7
	LogicOpNor          LogicOp = 8
	LogicOpEquivalent   LogicOp = 9
	LogicOpInvert       LogicOp = 10
	LogicOpOrReverse    LogicOp = 11
	LogicOpCopyInverted LogicOp = 12
	LogicOpOrInverted   LogicOp = 13
	LogicOpNand         LogicOp = 14
	LogicOpSet          LogicOp = 15
)

type PipelineBindPoint uint32

const (
	PipelineBindPointGraphics             PipelineBindPoint = 0
	PipelineBindPointCompute              PipelineBindPoint = 1
	PipelineBindPointExecutionGraphAmdx   PipelineBindPoint = 1000134000
	PipelineBindPointRayTracingKHR        PipelineBindPoint = 1000165000
	PipelineBindPointRayTracingNV         PipelineBindPoint = PipelineBindPointRayTracingKHR
	PipelineBindPointSubpassShadingHUAWEI PipelineBindPoint = 1000369003
	PipelineBindPointDataGraphARM         PipelineBindPoint = 1000507000
)

type PipelineCreateFlagBits uint32

const (
	PipelineCreateDisableOptimizationBit                                PipelineCreateFlagBits = 1 << 0
	PipelineCreateAllowDerivativesBit                                   PipelineCreateFlagBits = 1 << 1
	PipelineCreateDerivativeBit                                         PipelineCreateFlagBits = 1 << 2
	PipelineCreateDispatchBaseBit                                       PipelineCreateFlagBits = 1 << 4
	PipelineCreateDispatchBase                                          PipelineCreateFlagBits = PipelineCreateDispatchBaseBit
	PipelineCreateViewIndexFromDeviceIndexBit                           PipelineCreateFlagBits = 1 << 3
	PipelineCreateFailOnPipelineCompileRequiredBit                      PipelineCreateFlagBits = 1 << 8
	PipelineCreateEarlyReturnOnFailureBit                               PipelineCreateFlagBits = 1 << 9
	PipelineCreateNoProtectedAccessBit                                  PipelineCreateFlagBits = 1 << 27
	PipelineCreateProtectedAccessOnlyBit                                PipelineCreateFlagBits = 1 << 30
	PipelineCreateViewIndexFromDeviceIndexBitKHR                        PipelineCreateFlagBits = PipelineCreateViewIndexFromDeviceIndexBit
	PipelineCreateDispatchBaseBitKHR                                    PipelineCreateFlagBits = PipelineCreateDispatchBaseBit
	PipelineCreateDispatchBaseKHR                                       PipelineCreateFlagBits = PipelineCreateDispatchBaseBit
	PipelineCreateRayTracingNoNullAnyHitShadersBitKHR                   PipelineCreateFlagBits = 1 << 14
	PipelineCreateRayTracingNoNullClosestHitShadersBitKHR               PipelineCreateFlagBits = 1 << 15
	PipelineCreateRayTracingNoNullMissShadersBitKHR                     PipelineCreateFlagBits = 1 << 16
	PipelineCreateRayTracingNoNullIntersectionShadersBitKHR             PipelineCreateFlagBits = 1 << 17
	PipelineCreateRayTracingSkipTrianglesBitKHR                         PipelineCreateFlagBits = 1 << 12
	PipelineCreateRayTracingSkipAabbsBitKHR                             PipelineCreateFlagBits = 1 << 13
	PipelineCreateRayTracingShaderGroupHandleCaptureReplayBitKHR        PipelineCreateFlagBits = 1 << 19
	PipelineCreateDeferCompileBitNV                                     PipelineCreateFlagBits = 1 << 5
	PipelineCreateRenderingFragmentDensityMapAttachmentBitEXT           PipelineCreateFlagBits = 1 << 22
	PipelineRasterizationStateCreateFragmentDensityMapAttachmentBitEXT  PipelineCreateFlagBits = PipelineCreateRenderingFragmentDensityMapAttachmentBitEXT
	PipelineCreateRenderingFragmentShadingRateAttachmentBitKHR          PipelineCreateFlagBits = 1 << 21
	PipelineRasterizationStateCreateFragmentShadingRateAttachmentBitKHR PipelineCreateFlagBits = PipelineCreateRenderingFragmentShadingRateAttachmentBitKHR
	PipelineCreateCaptureStatisticsBitKHR                               PipelineCreateFlagBits = 1 << 6
	PipelineCreateCaptureInternalRepresentationsBitKHR                  PipelineCreateFlagBits = 1 << 7
	PipelineCreateIndirectBindableBitNV                                 PipelineCreateFlagBits = 1 << 18
	PipelineCreateLibraryBitKHR                                         PipelineCreateFlagBits = 1 << 11
	PipelineCreateFailOnPipelineCompileRequiredBitEXT                   PipelineCreateFlagBits = PipelineCreateFailOnPipelineCompileRequiredBit
	PipelineCreateEarlyReturnOnFailureBitEXT                            PipelineCreateFlagBits = PipelineCreateEarlyReturnOnFailureBit
	PipelineCreateDescriptorBufferBitEXT                                PipelineCreateFlagBits = 1 << 29
	PipelineCreateRetainLinkTimeOptimizationInfoBitEXT                  PipelineCreateFlagBits = 1 << 23
	PipelineCreateLinkTimeOptimizationBitEXT                            PipelineCreateFlagBits = 1 << 10
	PipelineCreateRayTracingAllowMotionBitNV                            PipelineCreateFlagBits = 1 << 20
	PipelineCreateColorAttachmentFeedbackLoopBitEXT                     PipelineCreateFlagBits = 1 << 25
	PipelineCreateDepthStencilAttachmentFeedbackLoopBitEXT              PipelineCreateFlagBits = 1 << 26
	PipelineCreateRayTracingOpacityMicromapBitEXT                       PipelineCreateFlagBits = 1 << 24
	PipelineCreateRayTracingDisplacementMicromapBitNV                   PipelineCreateFlagBits = 1 << 28
	PipelineCreateNoProtectedAccessBitEXT                               PipelineCreateFlagBits = PipelineCreateNoProtectedAccessBit
	PipelineCreateProtectedAccessOnlyBitEXT                             PipelineCreateFlagBits = PipelineCreateProtectedAccessOnlyBit
)

type PointClippingBehavior uint32

const (
	PointClippingBehaviorAllClipPlanes         PointClippingBehavior = 0
	PointClippingBehaviorUserClipPlanesOnly    PointClippingBehavior = 1
	PointClippingBehaviorAllClipPlanesKHR      PointClippingBehavior = PointClippingBehaviorAllClipPlanes
	PointClippingBehaviorUserClipPlanesOnlyKHR PointClippingBehavior = PointClippingBehaviorUserClipPlanesOnly
)

type PolygonMode uint32

const (
	PolygonModeFill            PolygonMode = 0
	PolygonModeLine            PolygonMode = 1
	PolygonModePoint           PolygonMode = 2
	PolygonModeFillRectangleNV PolygonMode = 1000155000
)

type PrimitiveTopology uint32

const (
	PrimitiveTopologyPointList                  PrimitiveTopology = 0
	PrimitiveTopologyLineList                   PrimitiveTopology = 1
	PrimitiveTopologyLineStrip                  PrimitiveTopology = 2
	PrimitiveTopologyTriangleList               PrimitiveTopology = 3
	PrimitiveTopologyTriangleStrip              PrimitiveTopology = 4
	PrimitiveTopologyTriangleFan                PrimitiveTopology = 5
	PrimitiveTopologyLineListWithAdjacency      PrimitiveTopology = 6
	PrimitiveTopologyLineStripWithAdjacency     PrimitiveTopology = 7
	PrimitiveTopologyTriangleListWithAdjacency  PrimitiveTopology = 8
	PrimitiveTopologyTriangleStripWithAdjacency PrimitiveTopology = 9
	PrimitiveTopologyPatchList                  PrimitiveTopology = 10
)

type RenderPassCreateFlagBits uint32

const (
	RenderPassCreateReserved3BitIMG                 RenderPassCreateFlagBits = 1 << 3
	RenderPassCreateReserved0BitKHR                 RenderPassCreateFlagBits = 1 << 0
	RenderPassCreateTransformBitQCOM                RenderPassCreateFlagBits = 1 << 1
	RenderPassCreatePerLayerFragmentDensityBitVALVE RenderPassCreateFlagBits = 1 << 2
)

type RenderingFlagBits uint32

const (
	RenderingContentsSecondaryCommandBuffersBit     RenderingFlagBits = 1 << 0
	RenderingSuspendingBit                          RenderingFlagBits = 1 << 1
	RenderingResumingBit                            RenderingFlagBits = 1 << 2
	RenderingContentsSecondaryCommandBuffersBitKHR  RenderingFlagBits = RenderingContentsSecondaryCommandBuffersBit
	RenderingSuspendingBitKHR                       RenderingFlagBits = RenderingSuspendingBit
	RenderingResumingBitKHR                         RenderingFlagBits = RenderingResumingBit
	RenderingReserved9BitIMG                        RenderingFlagBits = 1 << 9
	RenderingContentsInlineBitEXT                   RenderingFlagBits = RenderingContentsInlineBitKHR
	RenderingEnableLegacyDitheringBitEXT            RenderingFlagBits = 1 << 3
	RenderingContentsInlineBitKHR                   RenderingFlagBits = 1 << 4
	RenderingPerLayerFragmentDensityBitVALVE        RenderingFlagBits = 1 << 5
	RenderingFragmentRegionBitEXT                   RenderingFlagBits = 1 << 6
	RenderingCustomResolveBitEXT                    RenderingFlagBits = 1 << 7
	RenderingLocalReadConcurrentAccessControlBitKHR RenderingFlagBits = 1 << 8
)

type ResolveModeFlagBits uint32

const (
	ResolveModeNone                               ResolveModeFlagBits = 0
	ResolveModeSampleZeroBit                      ResolveModeFlagBits = 1 << 0
	ResolveModeAverageBit                         ResolveModeFlagBits = 1 << 1
	ResolveModeMinBit                             ResolveModeFlagBits = 1 << 2
	ResolveModeMaxBit                             ResolveModeFlagBits = 1 << 3
	ResolveModeNoneKHR                            ResolveModeFlagBits = ResolveModeNone
	ResolveModeSampleZeroBitKHR                   ResolveModeFlagBits = ResolveModeSampleZeroBit
	ResolveModeAverageBitKHR                      ResolveModeFlagBits = ResolveModeAverageBit
	ResolveModeMinBitKHR                          ResolveModeFlagBits = ResolveModeMinBit
	ResolveModeMaxBitKHR                          ResolveModeFlagBits = ResolveModeMaxBit
	ResolveModeExternalFormatDownsampleBitAndroid ResolveModeFlagBits = 1 << 4
	ResolveModeExternalFormatDownsampleAndroid    ResolveModeFlagBits = ResolveModeExternalFormatDownsampleBitAndroid
	ResolveModeCustomBitEXT                       ResolveModeFlagBits = 1 << 5
)

type SampleCountFlagBits uint32

const (
	SampleCount1Bit  SampleCountFlagBits = 1 << 0
	SampleCount2Bit  SampleCountFlagBits = 1 << 1
	SampleCount4Bit  SampleCountFlagBits = 1 << 2
	SampleCount8Bit  SampleCountFlagBits = 1 << 3
	SampleCount16Bit SampleCountFlagBits = 1 << 4
	SampleCount32Bit SampleCountFlagBits = 1 << 5
	SampleCount64Bit SampleCountFlagBits = 1 << 6
)

type ShaderStageFlagBits uint32

const (
	ShaderStageVertexBit                 ShaderStageFlagBits = 1 << 0
	ShaderStageTessellationControlBit    ShaderStageFlagBits = 1 << 1
	ShaderStageTessellationEvaluationBit ShaderStageFlagBits = 1 << 2
	ShaderStageGeometryBit               ShaderStageFlagBits = 1 << 3
	ShaderStageFragmentBit               ShaderStageFlagBits = 1 << 4
	ShaderStageComputeBit                ShaderStageFlagBits = 1 << 5
	ShaderStageAllGraphics               ShaderStageFlagBits = 0x0000001F
	ShaderStageAll                       ShaderStageFlagBits = 0x7FFFFFFF
	ShaderStageRaygenBitKHR              ShaderStageFlagBits = 1 << 8
	ShaderStageAnyHitBitKHR              ShaderStageFlagBits = 1 << 9
	ShaderStageClosestHitBitKHR          ShaderStageFlagBits = 1 << 10
	ShaderStageMissBitKHR                ShaderStageFlagBits = 1 << 11
	ShaderStageIntersectionBitKHR        ShaderStageFlagBits = 1 << 12
	ShaderStageCallableBitKHR            ShaderStageFlagBits = 1 << 13
	ShaderStageRaygenBitNV               ShaderStageFlagBits = ShaderStageRaygenBitKHR
	ShaderStageAnyHitBitNV               ShaderStageFlagBits = ShaderStageAnyHitBitKHR
	ShaderStageClosestHitBitNV           ShaderStageFlagBits = ShaderStageClosestHitBitKHR
	ShaderStageMissBitNV                 ShaderStageFlagBits = ShaderStageMissBitKHR
	ShaderStageIntersectionBitNV         ShaderStageFlagBits = ShaderStageIntersectionBitKHR
	ShaderStageCallableBitNV             ShaderStageFlagBits = ShaderStageCallableBitKHR
	ShaderStageTaskBitNV                 ShaderStageFlagBits = ShaderStageTaskBitEXT
	ShaderStageMeshBitNV                 ShaderStageFlagBits = ShaderStageMeshBitEXT
	ShaderStageTaskBitEXT                ShaderStageFlagBits = 1 << 6
	ShaderStageMeshBitEXT                ShaderStageFlagBits = 1 << 7
	ShaderStageSubpassShadingBitHUAWEI   ShaderStageFlagBits = 1 << 14
	ShaderStageClusterCullingBitHUAWEI   ShaderStageFlagBits = 1 << 19
	ShaderStageReserved15BitNV           ShaderStageFlagBits = 1 << 15
	ShaderStageReserved16BitHUAWEI       ShaderStageFlagBits = 1 << 16
)

type StencilFaceFlagBits uint32

const (
	StencilFaceFrontBit     StencilFaceFlagBits = 1 << 0
	StencilFaceBackBit      StencilFaceFlagBits = 1 << 1
	StencilFaceFrontAndBack StencilFaceFlagBits = 0x00000003
	StencilFrontAndBack     StencilFaceFlagBits = 0 // TODO: unknown value
)

type StencilOp uint32

const (
	StencilOpKeep              StencilOp = 0
	StencilOpZero              StencilOp = 1
	StencilOpReplace           StencilOp = 2
	StencilOpIncrementAndClamp StencilOp = 3
	StencilOpDecrementAndClamp StencilOp = 4
	StencilOpInvert            StencilOp = 5
	StencilOpIncrementAndWrap  StencilOp = 6
	StencilOpDecrementAndWrap  StencilOp = 7
)

type StructureType uint32

const (
	StructureTypeApplicationInfo                                              StructureType = 0
	StructureTypeInstanceCreateInfo                                           StructureType = 1
	StructureTypeDeviceQueueCreateInfo                                        StructureType = 2
	StructureTypeDeviceCreateInfo                                             StructureType = 3
	StructureTypeSubmitInfo                                                   StructureType = 4
	StructureTypeMemoryAllocateInfo                                           StructureType = 5
	StructureTypeMappedMemoryRange                                            StructureType = 6
	StructureTypeBindSparseInfo                                               StructureType = 7
	StructureTypeFenceCreateInfo                                              StructureType = 8
	StructureTypeSemaphoreCreateInfo                                          StructureType = 9
	StructureTypeEventCreateInfo                                              StructureType = 10
	StructureTypeQueryPoolCreateInfo                                          StructureType = 11
	StructureTypeBufferCreateInfo                                             StructureType = 12
	StructureTypeBufferViewCreateInfo                                         StructureType = 13
	StructureTypeImageCreateInfo                                              StructureType = 14
	StructureTypeImageViewCreateInfo                                          StructureType = 15
	StructureTypeShaderModuleCreateInfo                                       StructureType = 16
	StructureTypePipelineCacheCreateInfo                                      StructureType = 17
	StructureTypePipelineShaderStageCreateInfo                                StructureType = 18
	StructureTypePipelineVertexInputStateCreateInfo                           StructureType = 19
	StructureTypePipelineInputAssemblyStateCreateInfo                         StructureType = 20
	StructureTypePipelineTessellationStateCreateInfo                          StructureType = 21
	StructureTypePipelineViewportStateCreateInfo                              StructureType = 22
	StructureTypePipelineRasterizationStateCreateInfo                         StructureType = 23
	StructureTypePipelineMultisampleStateCreateInfo                           StructureType = 24
	StructureTypePipelineDepthStencilStateCreateInfo                          StructureType = 25
	StructureTypePipelineColorBlendStateCreateInfo                            StructureType = 26
	StructureTypePipelineDynamicStateCreateInfo                               StructureType = 27
	StructureTypeGraphicsPipelineCreateInfo                                   StructureType = 28
	StructureTypeComputePipelineCreateInfo                                    StructureType = 29
	StructureTypePipelineLayoutCreateInfo                                     StructureType = 30
	StructureTypeSamplerCreateInfo                                            StructureType = 31
	StructureTypeDescriptorSetLayoutCreateInfo                                StructureType = 32
	StructureTypeDescriptorPoolCreateInfo                                     StructureType = 33
	StructureTypeDescriptorSetAllocateInfo                                    StructureType = 34
	StructureTypeWriteDescriptorSet                                           StructureType = 35
	StructureTypeCopyDescriptorSet                                            StructureType = 36
	StructureTypeFramebufferCreateInfo                                        StructureType = 37
	StructureTypeRenderPassCreateInfo                                         StructureType = 38
	StructureTypeCommandPoolCreateInfo                                        StructureType = 39
	StructureTypeCommandBufferAllocateInfo                                    StructureType = 40
	StructureTypeCommandBufferInheritanceInfo                                 StructureType = 41
	StructureTypeCommandBufferBeginInfo                                       StructureType = 42
	StructureTypeRenderPassBeginInfo                                          StructureType = 43
	StructureTypeBufferMemoryBarrier                                          StructureType = 44
	StructureTypeImageMemoryBarrier                                           StructureType = 45
	StructureTypeMemoryBarrier                                                StructureType = 46
	StructureTypeLoaderInstanceCreateInfo                                     StructureType = 47
	StructureTypeLoaderDeviceCreateInfo                                       StructureType = 48
	StructureTypeBindBufferMemoryInfo                                         StructureType = 1000157000
	StructureTypeBindImageMemoryInfo                                          StructureType = 1000157001
	StructureTypeMemoryDedicatedRequirements                                  StructureType = 1000127000
	StructureTypeMemoryDedicatedAllocateInfo                                  StructureType = 1000127001
	StructureTypeMemoryAllocateFlagsInfo                                      StructureType = 1000060000
	StructureTypeDeviceGroupCommandBufferBeginInfo                            StructureType = 1000060004
	StructureTypeDeviceGroupSubmitInfo                                        StructureType = 1000060005
	StructureTypeDeviceGroupBindSparseInfo                                    StructureType = 1000060006
	StructureTypeBindBufferMemoryDeviceGroupInfo                              StructureType = 1000060013
	StructureTypeBindImageMemoryDeviceGroupInfo                               StructureType = 1000060014
	StructureTypePhysicalDeviceGroupProperties                                StructureType = 1000070000
	StructureTypeDeviceGroupDeviceCreateInfo                                  StructureType = 1000070001
	StructureTypeBufferMemoryRequirementsInfo2                                StructureType = 1000146000
	StructureTypeImageMemoryRequirementsInfo2                                 StructureType = 1000146001
	StructureTypeImageSparseMemoryRequirementsInfo2                           StructureType = 1000146002
	StructureTypeMemoryRequirements2                                          StructureType = 1000146003
	StructureTypeSparseImageMemoryRequirements2                               StructureType = 1000146004
	StructureTypePhysicalDeviceFeatures2                                      StructureType = 1000059000
	StructureTypePhysicalDeviceProperties2                                    StructureType = 1000059001
	StructureTypeFormatProperties2                                            StructureType = 1000059002
	StructureTypeImageFormatProperties2                                       StructureType = 1000059003
	StructureTypePhysicalDeviceImageFormatInfo2                               StructureType = 1000059004
	StructureTypeQueueFamilyProperties2                                       StructureType = 1000059005
	StructureTypePhysicalDeviceMemoryProperties2                              StructureType = 1000059006
	StructureTypeSparseImageFormatProperties2                                 StructureType = 1000059007
	StructureTypePhysicalDeviceSparseImageFormatInfo2                         StructureType = 1000059008
	StructureTypeImageViewUsageCreateInfo                                     StructureType = 1000117002
	StructureTypeProtectedSubmitInfo                                          StructureType = 1000145000
	StructureTypePhysicalDeviceProtectedMemoryFeatures                        StructureType = 1000145001
	StructureTypePhysicalDeviceProtectedMemoryProperties                      StructureType = 1000145002
	StructureTypeDeviceQueueInfo2                                             StructureType = 1000145003
	StructureTypePhysicalDeviceExternalImageFormatInfo                        StructureType = 1000071000
	StructureTypeExternalImageFormatProperties                                StructureType = 1000071001
	StructureTypePhysicalDeviceExternalBufferInfo                             StructureType = 1000071002
	StructureTypeExternalBufferProperties                                     StructureType = 1000071003
	StructureTypePhysicalDeviceIdProperties                                   StructureType = 1000071004
	StructureTypeExternalMemoryBufferCreateInfo                               StructureType = 1000072000
	StructureTypeExternalMemoryImageCreateInfo                                StructureType = 1000072001
	StructureTypeExportMemoryAllocateInfo                                     StructureType = 1000072002
	StructureTypePhysicalDeviceExternalFenceInfo                              StructureType = 1000112000
	StructureTypeExternalFenceProperties                                      StructureType = 1000112001
	StructureTypeExportFenceCreateInfo                                        StructureType = 1000113000
	StructureTypeExportSemaphoreCreateInfo                                    StructureType = 1000077000
	StructureTypePhysicalDeviceExternalSemaphoreInfo                          StructureType = 1000076000
	StructureTypeExternalSemaphoreProperties                                  StructureType = 1000076001
	StructureTypePhysicalDeviceSubgroupProperties                             StructureType = 1000094000
	StructureTypePhysicalDevice16bitStorageFeatures                           StructureType = 1000083000
	StructureTypePhysicalDeviceVariablePointersFeatures                       StructureType = 1000120000
	StructureTypePhysicalDeviceVariablePointerFeatures                        StructureType = StructureTypePhysicalDeviceVariablePointersFeatures
	StructureTypeDescriptorUpdateTemplateCreateInfo                           StructureType = 1000085000
	StructureTypePhysicalDeviceMaintenance3Properties                         StructureType = 1000168000
	StructureTypeDescriptorSetLayoutSupport                                   StructureType = 1000168001
	StructureTypeSamplerYcbcrConversionCreateInfo                             StructureType = 1000156000
	StructureTypeSamplerYcbcrConversionInfo                                   StructureType = 1000156001
	StructureTypeBindImagePlaneMemoryInfo                                     StructureType = 1000156002
	StructureTypeImagePlaneMemoryRequirementsInfo                             StructureType = 1000156003
	StructureTypePhysicalDeviceSamplerYcbcrConversionFeatures                 StructureType = 1000156004
	StructureTypeSamplerYcbcrConversionImageFormatProperties                  StructureType = 1000156005
	StructureTypeDeviceGroupRenderPassBeginInfo                               StructureType = 1000060003
	StructureTypePhysicalDevicePointClippingProperties                        StructureType = 1000117000
	StructureTypeRenderPassInputAttachmentAspectCreateInfo                    StructureType = 1000117001
	StructureTypePipelineTessellationDomainOriginStateCreateInfo              StructureType = 1000117003
	StructureTypeRenderPassMultiviewCreateInfo                                StructureType = 1000053000
	StructureTypePhysicalDeviceMultiviewFeatures                              StructureType = 1000053001
	StructureTypePhysicalDeviceMultiviewProperties                            StructureType = 1000053002
	StructureTypePhysicalDeviceShaderDrawParametersFeatures                   StructureType = 1000063000
	StructureTypePhysicalDeviceShaderDrawParameterFeatures                    StructureType = StructureTypePhysicalDeviceShaderDrawParametersFeatures
	StructureTypePhysicalDeviceDriverProperties                               StructureType = 1000196000
	StructureTypePhysicalDeviceVulkan11Features                               StructureType = 49
	StructureTypePhysicalDeviceVulkan11Properties                             StructureType = 50
	StructureTypePhysicalDeviceVulkan12Features                               StructureType = 51
	StructureTypePhysicalDeviceVulkan12Properties                             StructureType = 52
	StructureTypeImageFormatListCreateInfo                                    StructureType = 1000147000
	StructureTypePhysicalDeviceVulkanMemoryModelFeatures                      StructureType = 1000211000
	StructureTypePhysicalDeviceHostQueryResetFeatures                         StructureType = 1000261000
	StructureTypePhysicalDeviceTimelineSemaphoreFeatures                      StructureType = 1000207000
	StructureTypePhysicalDeviceTimelineSemaphoreProperties                    StructureType = 1000207001
	StructureTypeSemaphoreTypeCreateInfo                                      StructureType = 1000207002
	StructureTypeTimelineSemaphoreSubmitInfo                                  StructureType = 1000207003
	StructureTypeSemaphoreWaitInfo                                            StructureType = 1000207004
	StructureTypeSemaphoreSignalInfo                                          StructureType = 1000207005
	StructureTypePhysicalDeviceBufferDeviceAddressFeatures                    StructureType = 1000257000
	StructureTypeBufferDeviceAddressInfo                                      StructureType = 1000244001
	StructureTypeBufferOpaqueCaptureAddressCreateInfo                         StructureType = 1000257002
	StructureTypeMemoryOpaqueCaptureAddressAllocateInfo                       StructureType = 1000257003
	StructureTypeDeviceMemoryOpaqueCaptureAddressInfo                         StructureType = 1000257004
	StructureTypePhysicalDevice8bitStorageFeatures                            StructureType = 1000177000
	StructureTypePhysicalDeviceShaderAtomicInt64Features                      StructureType = 1000180000
	StructureTypePhysicalDeviceShaderFloat16Int8Features                      StructureType = 1000082000
	StructureTypePhysicalDeviceFloatControlsProperties                        StructureType = 1000197000
	StructureTypeDescriptorSetLayoutBindingFlagsCreateInfo                    StructureType = 1000161000
	StructureTypePhysicalDeviceDescriptorIndexingFeatures                     StructureType = 1000161001
	StructureTypePhysicalDeviceDescriptorIndexingProperties                   StructureType = 1000161002
	StructureTypeDescriptorSetVariableDescriptorCountAllocateInfo             StructureType = 1000161003
	StructureTypeDescriptorSetVariableDescriptorCountLayoutSupport            StructureType = 1000161004
	StructureTypePhysicalDeviceScalarBlockLayoutFeatures                      StructureType = 1000221000
	StructureTypePhysicalDeviceSamplerFilterMinmaxProperties                  StructureType = 1000130000
	StructureTypeSamplerReductionModeCreateInfo                               StructureType = 1000130001
	StructureTypePhysicalDeviceUniformBufferStandardLayoutFeatures            StructureType = 1000253000
	StructureTypePhysicalDeviceShaderSubgroupExtendedTypesFeatures            StructureType = 1000175000
	StructureTypeAttachmentDescription2                                       StructureType = 1000109000
	StructureTypeAttachmentReference2                                         StructureType = 1000109001
	StructureTypeSubpassDescription2                                          StructureType = 1000109002
	StructureTypeSubpassDependency2                                           StructureType = 1000109003
	StructureTypeRenderPassCreateInfo2                                        StructureType = 1000109004
	StructureTypeSubpassBeginInfo                                             StructureType = 1000109005
	StructureTypeSubpassEndInfo                                               StructureType = 1000109006
	StructureTypePhysicalDeviceDepthStencilResolveProperties                  StructureType = 1000199000
	StructureTypeSubpassDescriptionDepthStencilResolve                        StructureType = 1000199001
	StructureTypeImageStencilUsageCreateInfo                                  StructureType = 1000246000
	StructureTypePhysicalDeviceImagelessFramebufferFeatures                   StructureType = 1000108000
	StructureTypeFramebufferAttachmentsCreateInfo                             StructureType = 1000108001
	StructureTypeFramebufferAttachmentImageInfo                               StructureType = 1000108002
	StructureTypeRenderPassAttachmentBeginInfo                                StructureType = 1000108003
	StructureTypePhysicalDeviceSeparateDepthStencilLayoutsFeatures            StructureType = 1000241000
	StructureTypeAttachmentReferenceStencilLayout                             StructureType = 1000241001
	StructureTypeAttachmentDescriptionStencilLayout                           StructureType = 1000241002
	StructureTypePhysicalDeviceVulkan13Features                               StructureType = 53
	StructureTypePhysicalDeviceVulkan13Properties                             StructureType = 54
	StructureTypePhysicalDeviceToolProperties                                 StructureType = 1000245000
	StructureTypePhysicalDevicePrivateDataFeatures                            StructureType = 1000295000
	StructureTypeDevicePrivateDataCreateInfo                                  StructureType = 1000295001
	StructureTypePrivateDataSlotCreateInfo                                    StructureType = 1000295002
	StructureTypeMemoryBarrier2                                               StructureType = 1000314000
	StructureTypeBufferMemoryBarrier2                                         StructureType = 1000314001
	StructureTypeImageMemoryBarrier2                                          StructureType = 1000314002
	StructureTypeDependencyInfo                                               StructureType = 1000314003
	StructureTypeSubmitInfo2                                                  StructureType = 1000314004
	StructureTypeSemaphoreSubmitInfo                                          StructureType = 1000314005
	StructureTypeCommandBufferSubmitInfo                                      StructureType = 1000314006
	StructureTypePhysicalDeviceSynchronization2Features                       StructureType = 1000314007
	StructureTypeCopyBufferInfo2                                              StructureType = 1000337000
	StructureTypeCopyImageInfo2                                               StructureType = 1000337001
	StructureTypeCopyBufferToImageInfo2                                       StructureType = 1000337002
	StructureTypeCopyImageToBufferInfo2                                       StructureType = 1000337003
	StructureTypeBufferCopy2                                                  StructureType = 1000337006
	StructureTypeImageCopy2                                                   StructureType = 1000337007
	StructureTypeBufferImageCopy2                                             StructureType = 1000337009
	StructureTypePhysicalDeviceTextureCompressionAstcHdrFeatures              StructureType = 1000066000
	StructureTypeFormatProperties3                                            StructureType = 1000360000
	StructureTypePhysicalDeviceMaintenance4Features                           StructureType = 1000413000
	StructureTypePhysicalDeviceMaintenance4Properties                         StructureType = 1000413001
	StructureTypeDeviceBufferMemoryRequirements                               StructureType = 1000413002
	StructureTypeDeviceImageMemoryRequirements                                StructureType = 1000413003
	StructureTypePipelineCreationFeedbackCreateInfo                           StructureType = 1000192000
	StructureTypePhysicalDeviceShaderTerminateInvocationFeatures              StructureType = 1000215000
	StructureTypePhysicalDeviceShaderDemoteToHelperInvocationFeatures         StructureType = 1000276000
	StructureTypePhysicalDevicePipelineCreationCacheControlFeatures           StructureType = 1000297000
	StructureTypePhysicalDeviceZeroInitializeWorkgroupMemoryFeatures          StructureType = 1000325000
	StructureTypePhysicalDeviceImageRobustnessFeatures                        StructureType = 1000335000
	StructureTypePhysicalDeviceSubgroupSizeControlProperties                  StructureType = 1000225000
	StructureTypePipelineShaderStageRequiredSubgroupSizeCreateInfo            StructureType = 1000225001
	StructureTypePhysicalDeviceSubgroupSizeControlFeatures                    StructureType = 1000225002
	StructureTypePhysicalDeviceInlineUniformBlockFeatures                     StructureType = 1000138000
	StructureTypePhysicalDeviceInlineUniformBlockProperties                   StructureType = 1000138001
	StructureTypeWriteDescriptorSetInlineUniformBlock                         StructureType = 1000138002
	StructureTypeDescriptorPoolInlineUniformBlockCreateInfo                   StructureType = 1000138003
	StructureTypePhysicalDeviceShaderIntegerDotProductFeatures                StructureType = 1000280000
	StructureTypePhysicalDeviceShaderIntegerDotProductProperties              StructureType = 1000280001
	StructureTypePhysicalDeviceTexelBufferAlignmentProperties                 StructureType = 1000281001
	StructureTypeBlitImageInfo2                                               StructureType = 1000337004
	StructureTypeResolveImageInfo2                                            StructureType = 1000337005
	StructureTypeImageBlit2                                                   StructureType = 1000337008
	StructureTypeImageResolve2                                                StructureType = 1000337010
	StructureTypeRenderingInfo                                                StructureType = 1000044000
	StructureTypeRenderingAttachmentInfo                                      StructureType = 1000044001
	StructureTypePipelineRenderingCreateInfo                                  StructureType = 1000044002
	StructureTypePhysicalDeviceDynamicRenderingFeatures                       StructureType = 1000044003
	StructureTypeCommandBufferInheritanceRenderingInfo                        StructureType = 1000044004
	StructureTypePhysicalDeviceVulkan14Features                               StructureType = 55
	StructureTypePhysicalDeviceVulkan14Properties                             StructureType = 56
	StructureTypeDeviceQueueGlobalPriorityCreateInfo                          StructureType = 1000174000
	StructureTypePhysicalDeviceGlobalPriorityQueryFeatures                    StructureType = 1000388000
	StructureTypeQueueFamilyGlobalPriorityProperties                          StructureType = 1000388001
	StructureTypePhysicalDeviceIndexTypeUint8Features                         StructureType = 1000265000
	StructureTypeMemoryMapInfo                                                StructureType = 1000271000
	StructureTypeMemoryUnmapInfo                                              StructureType = 1000271001
	StructureTypePhysicalDeviceMaintenance5Features                           StructureType = 1000470000
	StructureTypePhysicalDeviceMaintenance5Properties                         StructureType = 1000470001
	StructureTypeDeviceImageSubresourceInfo                                   StructureType = 1000470004
	StructureTypeSubresourceLayout2                                           StructureType = 1000338002
	StructureTypeImageSubresource2                                            StructureType = 1000338003
	StructureTypeBufferUsageFlags2CreateInfo                                  StructureType = 1000470006
	StructureTypePhysicalDeviceMaintenance6Features                           StructureType = 1000545000
	StructureTypePhysicalDeviceMaintenance6Properties                         StructureType = 1000545001
	StructureTypeBindMemoryStatus                                             StructureType = 1000545002
	StructureTypePhysicalDeviceHostImageCopyFeatures                          StructureType = 1000270000
	StructureTypePhysicalDeviceHostImageCopyProperties                        StructureType = 1000270001
	StructureTypeMemoryToImageCopy                                            StructureType = 1000270002
	StructureTypeImageToMemoryCopy                                            StructureType = 1000270003
	StructureTypeCopyImageToMemoryInfo                                        StructureType = 1000270004
	StructureTypeCopyMemoryToImageInfo                                        StructureType = 1000270005
	StructureTypeHostImageLayoutTransitionInfo                                StructureType = 1000270006
	StructureTypeCopyImageToImageInfo                                         StructureType = 1000270007
	StructureTypeSubresourceHostMemcpySize                                    StructureType = 1000270008
	StructureTypeHostImageCopyDevicePerformanceQuery                          StructureType = 1000270009
	StructureTypePhysicalDeviceShaderSubgroupRotateFeatures                   StructureType = 1000416000
	StructureTypePhysicalDeviceShaderFloatControls2Features                   StructureType = 1000528000
	StructureTypePhysicalDeviceShaderExpectAssumeFeatures                     StructureType = 1000544000
	StructureTypePipelineCreateFlags2CreateInfo                               StructureType = 1000470005
	StructureTypePhysicalDevicePushDescriptorProperties                       StructureType = 1000080000
	StructureTypeBindDescriptorSetsInfo                                       StructureType = 1000545003
	StructureTypePushConstantsInfo                                            StructureType = 1000545004
	StructureTypePushDescriptorSetInfo                                        StructureType = 1000545005
	StructureTypePushDescriptorSetWithTemplateInfo                            StructureType = 1000545006
	StructureTypePhysicalDevicePipelineProtectedAccessFeatures                StructureType = 1000466000
	StructureTypePipelineRobustnessCreateInfo                                 StructureType = 1000068000
	StructureTypePhysicalDevicePipelineRobustnessFeatures                     StructureType = 1000068001
	StructureTypePhysicalDevicePipelineRobustnessProperties                   StructureType = 1000068002
	StructureTypePhysicalDeviceLineRasterizationFeatures                      StructureType = 1000259000
	StructureTypePipelineRasterizationLineStateCreateInfo                     StructureType = 1000259001
	StructureTypePhysicalDeviceLineRasterizationProperties                    StructureType = 1000259002
	StructureTypePhysicalDeviceVertexAttributeDivisorProperties               StructureType = 1000525000
	StructureTypePipelineVertexInputDivisorStateCreateInfo                    StructureType = 1000190001
	StructureTypePhysicalDeviceVertexAttributeDivisorFeatures                 StructureType = 1000190002
	StructureTypeRenderingAreaInfo                                            StructureType = 1000470003
	StructureTypePhysicalDeviceDynamicRenderingLocalReadFeatures              StructureType = 1000232000
	StructureTypeRenderingAttachmentLocationInfo                              StructureType = 1000232001
	StructureTypeRenderingInputAttachmentIndexInfo                            StructureType = 1000232002
	StructureTypePhysicalDeviceVulkanSc10Features                             StructureType = 1000298000
	StructureTypePhysicalDeviceVulkanSc10Properties                           StructureType = 1000298001
	StructureTypeDeviceObjectReservationCreateInfo                            StructureType = 1000298002
	StructureTypeCommandPoolMemoryReservationCreateInfo                       StructureType = 1000298003
	StructureTypeCommandPoolMemoryConsumption                                 StructureType = 1000298004
	StructureTypePipelinePoolSize                                             StructureType = 1000298005
	StructureTypeFaultData                                                    StructureType = 1000298007
	StructureTypeFaultCallbackInfo                                            StructureType = 1000298008
	StructureTypePipelineOfflineCreateInfo                                    StructureType = 1000298010
	StructureTypeSwapchainCreateInfoKHR                                       StructureType = 1000001000
	StructureTypePresentInfoKHR                                               StructureType = 1000001001
	StructureTypeDeviceGroupPresentCapabilitiesKHR                            StructureType = 1000060007
	StructureTypeImageSwapchainCreateInfoKHR                                  StructureType = 1000060008
	StructureTypeBindImageMemorySwapchainInfoKHR                              StructureType = 1000060009
	StructureTypeAcquireNextImageInfoKHR                                      StructureType = 1000060010
	StructureTypeDeviceGroupPresentInfoKHR                                    StructureType = 1000060011
	StructureTypeDeviceGroupSwapchainCreateInfoKHR                            StructureType = 1000060012
	StructureTypeDisplayModeCreateInfoKHR                                     StructureType = 1000002000
	StructureTypeDisplaySurfaceCreateInfoKHR                                  StructureType = 1000002001
	StructureTypeDisplayPresentInfoKHR                                        StructureType = 1000003000
	StructureTypeXlibSurfaceCreateInfoKHR                                     StructureType = 1000004000
	StructureTypeXcbSurfaceCreateInfoKHR                                      StructureType = 1000005000
	StructureTypeWaylandSurfaceCreateInfoKHR                                  StructureType = 1000006000
	StructureTypeAndroidSurfaceCreateInfoKHR                                  StructureType = 1000008000
	StructureTypeWin32SurfaceCreateInfoKHR                                    StructureType = 1000009000
	StructureTypeNativeBufferAndroid                                          StructureType = 1000010000
	StructureTypeSwapchainImageCreateInfoAndroid                              StructureType = 1000010001
	StructureTypePhysicalDevicePresentationPropertiesAndroid                  StructureType = 1000010002
	StructureTypeDebugReportCallbackCreateInfoEXT                             StructureType = 1000011000
	StructureTypeDebugReportCreateInfoEXT                                     StructureType = StructureTypeDebugReportCallbackCreateInfoEXT
	StructureTypePipelineRasterizationStateRasterizationOrderAMD              StructureType = 1000018000
	StructureTypeDebugMarkerObjectNameInfoEXT                                 StructureType = 1000022000
	StructureTypeDebugMarkerObjectTagInfoEXT                                  StructureType = 1000022001
	StructureTypeDebugMarkerMarkerInfoEXT                                     StructureType = 1000022002
	StructureTypeVideoProfileInfoKHR                                          StructureType = 1000023000
	StructureTypeVideoCapabilitiesKHR                                         StructureType = 1000023001
	StructureTypeVideoPictureResourceInfoKHR                                  StructureType = 1000023002
	StructureTypeVideoSessionMemoryRequirementsKHR                            StructureType = 1000023003
	StructureTypeBindVideoSessionMemoryInfoKHR                                StructureType = 1000023004
	StructureTypeVideoSessionCreateInfoKHR                                    StructureType = 1000023005
	StructureTypeVideoSessionParametersCreateInfoKHR                          StructureType = 1000023006
	StructureTypeVideoSessionParametersUpdateInfoKHR                          StructureType = 1000023007
	StructureTypeVideoBeginCodingInfoKHR                                      StructureType = 1000023008
	StructureTypeVideoEndCodingInfoKHR                                        StructureType = 1000023009
	StructureTypeVideoCodingControlInfoKHR                                    StructureType = 1000023010
	StructureTypeVideoReferenceSlotInfoKHR                                    StructureType = 1000023011
	StructureTypeQueueFamilyVideoPropertiesKHR                                StructureType = 1000023012
	StructureTypeVideoProfileListInfoKHR                                      StructureType = 1000023013
	StructureTypePhysicalDeviceVideoFormatInfoKHR                             StructureType = 1000023014
	StructureTypeVideoFormatPropertiesKHR                                     StructureType = 1000023015
	StructureTypeQueueFamilyQueryResultStatusPropertiesKHR                    StructureType = 1000023016
	StructureTypeVideoDecodeInfoKHR                                           StructureType = 1000024000
	StructureTypeVideoDecodeCapabilitiesKHR                                   StructureType = 1000024001
	StructureTypeVideoDecodeUsageInfoKHR                                      StructureType = 1000024002
	StructureTypeDedicatedAllocationImageCreateInfoNV                         StructureType = 1000026000
	StructureTypeDedicatedAllocationBufferCreateInfoNV                        StructureType = 1000026001
	StructureTypeDedicatedAllocationMemoryAllocateInfoNV                      StructureType = 1000026002
	StructureTypePhysicalDeviceTransformFeedbackFeaturesEXT                   StructureType = 1000028000
	StructureTypePhysicalDeviceTransformFeedbackPropertiesEXT                 StructureType = 1000028001
	StructureTypePipelineRasterizationStateStreamCreateInfoEXT                StructureType = 1000028002
	StructureTypeCuModuleCreateInfoNVX                                        StructureType = 1000029000
	StructureTypeCuFunctionCreateInfoNVX                                      StructureType = 1000029001
	StructureTypeCuLaunchInfoNVX                                              StructureType = 1000029002
	StructureTypeCuModuleTexturingModeCreateInfoNVX                           StructureType = 1000029004
	StructureTypeImageViewHandleInfoNVX                                       StructureType = 1000030000
	StructureTypeImageViewAddressPropertiesNVX                                StructureType = 1000030001
	StructureTypeVideoEncodeH264CapabilitiesKHR                               StructureType = 1000038000
	StructureTypeVideoEncodeH264SessionParametersCreateInfoKHR                StructureType = 1000038001
	StructureTypeVideoEncodeH264SessionParametersAddInfoKHR                   StructureType = 1000038002
	StructureTypeVideoEncodeH264PictureInfoKHR                                StructureType = 1000038003
	StructureTypeVideoEncodeH264DpbSlotInfoKHR                                StructureType = 1000038004
	StructureTypeVideoEncodeH264NaluSliceInfoKHR                              StructureType = 1000038005
	StructureTypeVideoEncodeH264GopRemainingFrameInfoKHR                      StructureType = 1000038006
	StructureTypeVideoEncodeH264ProfileInfoKHR                                StructureType = 1000038007
	StructureTypeVideoEncodeH264RateControlInfoKHR                            StructureType = 1000038008
	StructureTypeVideoEncodeH264RateControlLayerInfoKHR                       StructureType = 1000038009
	StructureTypeVideoEncodeH264SessionCreateInfoKHR                          StructureType = 1000038010
	StructureTypeVideoEncodeH264QualityLevelPropertiesKHR                     StructureType = 1000038011
	StructureTypeVideoEncodeH264SessionParametersGetInfoKHR                   StructureType = 1000038012
	StructureTypeVideoEncodeH264SessionParametersFeedbackInfoKHR              StructureType = 1000038013
	StructureTypeVideoEncodeH265CapabilitiesKHR                               StructureType = 1000039000
	StructureTypeVideoEncodeH265SessionParametersCreateInfoKHR                StructureType = 1000039001
	StructureTypeVideoEncodeH265SessionParametersAddInfoKHR                   StructureType = 1000039002
	StructureTypeVideoEncodeH265PictureInfoKHR                                StructureType = 1000039003
	StructureTypeVideoEncodeH265DpbSlotInfoKHR                                StructureType = 1000039004
	StructureTypeVideoEncodeH265NaluSliceSegmentInfoKHR                       StructureType = 1000039005
	StructureTypeVideoEncodeH265GopRemainingFrameInfoKHR                      StructureType = 1000039006
	StructureTypeVideoEncodeH265ProfileInfoKHR                                StructureType = 1000039007
	StructureTypeVideoEncodeH265RateControlInfoKHR                            StructureType = 1000039009
	StructureTypeVideoEncodeH265RateControlLayerInfoKHR                       StructureType = 1000039010
	StructureTypeVideoEncodeH265SessionCreateInfoKHR                          StructureType = 1000039011
	StructureTypeVideoEncodeH265QualityLevelPropertiesKHR                     StructureType = 1000039012
	StructureTypeVideoEncodeH265SessionParametersGetInfoKHR                   StructureType = 1000039013
	StructureTypeVideoEncodeH265SessionParametersFeedbackInfoKHR              StructureType = 1000039014
	StructureTypeVideoDecodeH264CapabilitiesKHR                               StructureType = 1000040000
	StructureTypeVideoDecodeH264PictureInfoKHR                                StructureType = 1000040001
	StructureTypeVideoDecodeH264ProfileInfoKHR                                StructureType = 1000040003
	StructureTypeVideoDecodeH264SessionParametersCreateInfoKHR                StructureType = 1000040004
	StructureTypeVideoDecodeH264SessionParametersAddInfoKHR                   StructureType = 1000040005
	StructureTypeVideoDecodeH264DpbSlotInfoKHR                                StructureType = 1000040006
	StructureTypeTextureLodGatherFormatPropertiesAMD                          StructureType = 1000041000
	StructureTypeRenderingInfoKHR                                             StructureType = StructureTypeRenderingInfo
	StructureTypeRenderingAttachmentInfoKHR                                   StructureType = StructureTypeRenderingAttachmentInfo
	StructureTypePipelineRenderingCreateInfoKHR                               StructureType = StructureTypePipelineRenderingCreateInfo
	StructureTypePhysicalDeviceDynamicRenderingFeaturesKHR                    StructureType = StructureTypePhysicalDeviceDynamicRenderingFeatures
	StructureTypeCommandBufferInheritanceRenderingInfoKHR                     StructureType = StructureTypeCommandBufferInheritanceRenderingInfo
	StructureTypeStreamDescriptorSurfaceCreateInfoGgp                         StructureType = 1000049000
	StructureTypePhysicalDeviceCornerSampledImageFeaturesNV                   StructureType = 1000050000
	StructureTypePrivateVendorInfoPlaceholderOffset0NV                        StructureType = 1000051000
	StructureTypeRenderPassMultiviewCreateInfoKHR                             StructureType = StructureTypeRenderPassMultiviewCreateInfo
	StructureTypePhysicalDeviceMultiviewFeaturesKHR                           StructureType = StructureTypePhysicalDeviceMultiviewFeatures
	StructureTypePhysicalDeviceMultiviewPropertiesKHR                         StructureType = StructureTypePhysicalDeviceMultiviewProperties
	StructureTypeExternalMemoryImageCreateInfoNV                              StructureType = 1000056000
	StructureTypeExportMemoryAllocateInfoNV                                   StructureType = 1000056001
	StructureTypeImportMemoryWin32HandleInfoNV                                StructureType = 1000057000
	StructureTypeExportMemoryWin32HandleInfoNV                                StructureType = 1000057001
	StructureTypeWin32KeyedMutexAcquireReleaseInfoNV                          StructureType = 1000058000
	StructureTypePhysicalDeviceFeatures2KHR                                   StructureType = StructureTypePhysicalDeviceFeatures2
	StructureTypePhysicalDeviceProperties2KHR                                 StructureType = StructureTypePhysicalDeviceProperties2
	StructureTypeFormatProperties2KHR                                         StructureType = StructureTypeFormatProperties2
	StructureTypeImageFormatProperties2KHR                                    StructureType = StructureTypeImageFormatProperties2
	StructureTypePhysicalDeviceImageFormatInfo2KHR                            StructureType = StructureTypePhysicalDeviceImageFormatInfo2
	StructureTypeQueueFamilyProperties2KHR                                    StructureType = StructureTypeQueueFamilyProperties2
	StructureTypePhysicalDeviceMemoryProperties2KHR                           StructureType = StructureTypePhysicalDeviceMemoryProperties2
	StructureTypeSparseImageFormatProperties2KHR                              StructureType = StructureTypeSparseImageFormatProperties2
	StructureTypePhysicalDeviceSparseImageFormatInfo2KHR                      StructureType = StructureTypePhysicalDeviceSparseImageFormatInfo2
	StructureTypeMemoryAllocateFlagsInfoKHR                                   StructureType = StructureTypeMemoryAllocateFlagsInfo
	StructureTypeDeviceGroupRenderPassBeginInfoKHR                            StructureType = StructureTypeDeviceGroupRenderPassBeginInfo
	StructureTypeDeviceGroupCommandBufferBeginInfoKHR                         StructureType = StructureTypeDeviceGroupCommandBufferBeginInfo
	StructureTypeDeviceGroupSubmitInfoKHR                                     StructureType = StructureTypeDeviceGroupSubmitInfo
	StructureTypeDeviceGroupBindSparseInfoKHR                                 StructureType = StructureTypeDeviceGroupBindSparseInfo
	StructureTypeBindBufferMemoryDeviceGroupInfoKHR                           StructureType = StructureTypeBindBufferMemoryDeviceGroupInfo
	StructureTypeBindImageMemoryDeviceGroupInfoKHR                            StructureType = StructureTypeBindImageMemoryDeviceGroupInfo
	StructureTypeValidationFlagsEXT                                           StructureType = 1000061000
	StructureTypeViSurfaceCreateInfoNn                                        StructureType = 1000062000
	StructureTypePhysicalDeviceTextureCompressionAstcHdrFeaturesEXT           StructureType = StructureTypePhysicalDeviceTextureCompressionAstcHdrFeatures
	StructureTypeImageViewAstcDecodeModeEXT                                   StructureType = 1000067000
	StructureTypePhysicalDeviceAstcDecodeFeaturesEXT                          StructureType = 1000067001
	StructureTypePipelineRobustnessCreateInfoEXT                              StructureType = StructureTypePipelineRobustnessCreateInfo
	StructureTypePhysicalDevicePipelineRobustnessFeaturesEXT                  StructureType = StructureTypePhysicalDevicePipelineRobustnessFeatures
	StructureTypePhysicalDevicePipelineRobustnessPropertiesEXT                StructureType = StructureTypePhysicalDevicePipelineRobustnessProperties
	StructureTypePhysicalDeviceGroupPropertiesKHR                             StructureType = StructureTypePhysicalDeviceGroupProperties
	StructureTypeDeviceGroupDeviceCreateInfoKHR                               StructureType = StructureTypeDeviceGroupDeviceCreateInfo
	StructureTypePhysicalDeviceExternalImageFormatInfoKHR                     StructureType = StructureTypePhysicalDeviceExternalImageFormatInfo
	StructureTypeExternalImageFormatPropertiesKHR                             StructureType = StructureTypeExternalImageFormatProperties
	StructureTypePhysicalDeviceExternalBufferInfoKHR                          StructureType = StructureTypePhysicalDeviceExternalBufferInfo
	StructureTypeExternalBufferPropertiesKHR                                  StructureType = StructureTypeExternalBufferProperties
	StructureTypePhysicalDeviceIdPropertiesKHR                                StructureType = StructureTypePhysicalDeviceIdProperties
	StructureTypeExternalMemoryBufferCreateInfoKHR                            StructureType = StructureTypeExternalMemoryBufferCreateInfo
	StructureTypeExternalMemoryImageCreateInfoKHR                             StructureType = StructureTypeExternalMemoryImageCreateInfo
	StructureTypeExportMemoryAllocateInfoKHR                                  StructureType = StructureTypeExportMemoryAllocateInfo
	StructureTypeImportMemoryWin32HandleInfoKHR                               StructureType = 1000073000
	StructureTypeExportMemoryWin32HandleInfoKHR                               StructureType = 1000073001
	StructureTypeMemoryWin32HandlePropertiesKHR                               StructureType = 1000073002
	StructureTypeMemoryGetWin32HandleInfoKHR                                  StructureType = 1000073003
	StructureTypeImportMemoryFdInfoKHR                                        StructureType = 1000074000
	StructureTypeMemoryFdPropertiesKHR                                        StructureType = 1000074001
	StructureTypeMemoryGetFdInfoKHR                                           StructureType = 1000074002
	StructureTypeWin32KeyedMutexAcquireReleaseInfoKHR                         StructureType = 1000075000
	StructureTypePhysicalDeviceExternalSemaphoreInfoKHR                       StructureType = StructureTypePhysicalDeviceExternalSemaphoreInfo
	StructureTypeExternalSemaphorePropertiesKHR                               StructureType = StructureTypeExternalSemaphoreProperties
	StructureTypeExportSemaphoreCreateInfoKHR                                 StructureType = StructureTypeExportSemaphoreCreateInfo
	StructureTypeImportSemaphoreWin32HandleInfoKHR                            StructureType = 1000078000
	StructureTypeExportSemaphoreWin32HandleInfoKHR                            StructureType = 1000078001
	StructureTypeD3d12FenceSubmitInfoKHR                                      StructureType = 1000078002
	StructureTypeSemaphoreGetWin32HandleInfoKHR                               StructureType = 1000078003
	StructureTypeImportSemaphoreFdInfoKHR                                     StructureType = 1000079000
	StructureTypeSemaphoreGetFdInfoKHR                                        StructureType = 1000079001
	StructureTypePhysicalDevicePushDescriptorPropertiesKHR                    StructureType = StructureTypePhysicalDevicePushDescriptorProperties
	StructureTypeCommandBufferInheritanceConditionalRenderingInfoEXT          StructureType = 1000081000
	StructureTypePhysicalDeviceConditionalRenderingFeaturesEXT                StructureType = 1000081001
	StructureTypeConditionalRenderingBeginInfoEXT                             StructureType = 1000081002
	StructureTypePhysicalDeviceShaderFloat16Int8FeaturesKHR                   StructureType = StructureTypePhysicalDeviceShaderFloat16Int8Features
	StructureTypePhysicalDeviceFloat16Int8FeaturesKHR                         StructureType = StructureTypePhysicalDeviceShaderFloat16Int8Features
	StructureTypePhysicalDevice16bitStorageFeaturesKHR                        StructureType = StructureTypePhysicalDevice16bitStorageFeatures
	StructureTypePresentRegionsKHR                                            StructureType = 1000084000
	StructureTypeDescriptorUpdateTemplateCreateInfoKHR                        StructureType = StructureTypeDescriptorUpdateTemplateCreateInfo
	StructureTypePipelineViewportWScalingStateCreateInfoNV                    StructureType = 1000087000
	StructureTypeSurfaceCapabilities2EXT                                      StructureType = 1000090000
	StructureTypeDisplayPowerInfoEXT                                          StructureType = 1000091000
	StructureTypeDeviceEventInfoEXT                                           StructureType = 1000091001
	StructureTypeDisplayEventInfoEXT                                          StructureType = 1000091002
	StructureTypeSwapchainCounterCreateInfoEXT                                StructureType = 1000091003
	StructureTypePresentTimesInfoGOOGLE                                       StructureType = 1000092000
	StructureTypePhysicalDeviceMultiviewPerViewAttributesPropertiesNVX        StructureType = 1000097000
	StructureTypeMultiviewPerViewAttributesInfoNVX                            StructureType = 1000044009
	StructureTypePipelineViewportSwizzleStateCreateInfoNV                     StructureType = 1000098000
	StructureTypePhysicalDeviceDiscardRectanglePropertiesEXT                  StructureType = 1000099000
	StructureTypePipelineDiscardRectangleStateCreateInfoEXT                   StructureType = 1000099001
	StructureTypePhysicalDeviceConservativeRasterizationPropertiesEXT         StructureType = 1000101000
	StructureTypePipelineRasterizationConservativeStateCreateInfoEXT          StructureType = 1000101001
	StructureTypePhysicalDeviceDepthClipEnableFeaturesEXT                     StructureType = 1000102000
	StructureTypePipelineRasterizationDepthClipStateCreateInfoEXT             StructureType = 1000102001
	StructureTypeHdrMetadataEXT                                               StructureType = 1000105000
	StructureTypePhysicalDeviceImagelessFramebufferFeaturesKHR                StructureType = StructureTypePhysicalDeviceImagelessFramebufferFeatures
	StructureTypeFramebufferAttachmentsCreateInfoKHR                          StructureType = StructureTypeFramebufferAttachmentsCreateInfo
	StructureTypeFramebufferAttachmentImageInfoKHR                            StructureType = StructureTypeFramebufferAttachmentImageInfo
	StructureTypeRenderPassAttachmentBeginInfoKHR                             StructureType = StructureTypeRenderPassAttachmentBeginInfo
	StructureTypeAttachmentDescription2KHR                                    StructureType = StructureTypeAttachmentDescription2
	StructureTypeAttachmentReference2KHR                                      StructureType = StructureTypeAttachmentReference2
	StructureTypeSubpassDescription2KHR                                       StructureType = StructureTypeSubpassDescription2
	StructureTypeSubpassDependency2KHR                                        StructureType = StructureTypeSubpassDependency2
	StructureTypeRenderPassCreateInfo2KHR                                     StructureType = StructureTypeRenderPassCreateInfo2
	StructureTypeSubpassBeginInfoKHR                                          StructureType = StructureTypeSubpassBeginInfo
	StructureTypeSubpassEndInfoKHR                                            StructureType = StructureTypeSubpassEndInfo
	StructureTypePhysicalDeviceRelaxedLineRasterizationFeaturesIMG            StructureType = 1000110000
	StructureTypeSharedPresentSurfaceCapabilitiesKHR                          StructureType = 1000111000
	StructureTypePhysicalDeviceExternalFenceInfoKHR                           StructureType = StructureTypePhysicalDeviceExternalFenceInfo
	StructureTypeExternalFencePropertiesKHR                                   StructureType = StructureTypeExternalFenceProperties
	StructureTypeExportFenceCreateInfoKHR                                     StructureType = StructureTypeExportFenceCreateInfo
	StructureTypeImportFenceWin32HandleInfoKHR                                StructureType = 1000114000
	StructureTypeExportFenceWin32HandleInfoKHR                                StructureType = 1000114001
	StructureTypeFenceGetWin32HandleInfoKHR                                   StructureType = 1000114002
	StructureTypeImportFenceFdInfoKHR                                         StructureType = 1000115000
	StructureTypeFenceGetFdInfoKHR                                            StructureType = 1000115001
	StructureTypePhysicalDevicePerformanceQueryFeaturesKHR                    StructureType = 1000116000
	StructureTypePhysicalDevicePerformanceQueryPropertiesKHR                  StructureType = 1000116001
	StructureTypeQueryPoolPerformanceCreateInfoKHR                            StructureType = 1000116002
	StructureTypePerformanceQuerySubmitInfoKHR                                StructureType = 1000116003
	StructureTypeAcquireProfilingLockInfoKHR                                  StructureType = 1000116004
	StructureTypePerformanceCounterKHR                                        StructureType = 1000116005
	StructureTypePerformanceCounterDescriptionKHR                             StructureType = 1000116006
	StructureTypePerformanceQueryReservationInfoKHR                           StructureType = 1000116007
	StructureTypePhysicalDevicePointClippingPropertiesKHR                     StructureType = StructureTypePhysicalDevicePointClippingProperties
	StructureTypeRenderPassInputAttachmentAspectCreateInfoKHR                 StructureType = StructureTypeRenderPassInputAttachmentAspectCreateInfo
	StructureTypeImageViewUsageCreateInfoKHR                                  StructureType = StructureTypeImageViewUsageCreateInfo
	StructureTypePipelineTessellationDomainOriginStateCreateInfoKHR           StructureType = StructureTypePipelineTessellationDomainOriginStateCreateInfo
	StructureTypePhysicalDeviceSurfaceInfo2KHR                                StructureType = 1000119000
	StructureTypeSurfaceCapabilities2KHR                                      StructureType = 1000119001
	StructureTypeSurfaceFormat2KHR                                            StructureType = 1000119002
	StructureTypePhysicalDeviceVariablePointersFeaturesKHR                    StructureType = StructureTypePhysicalDeviceVariablePointersFeatures
	StructureTypePhysicalDeviceVariablePointerFeaturesKHR                     StructureType = StructureTypePhysicalDeviceVariablePointersFeaturesKHR
	StructureTypeDisplayProperties2KHR                                        StructureType = 1000121000
	StructureTypeDisplayPlaneProperties2KHR                                   StructureType = 1000121001
	StructureTypeDisplayModeProperties2KHR                                    StructureType = 1000121002
	StructureTypeDisplayPlaneInfo2KHR                                         StructureType = 1000121003
	StructureTypeDisplayPlaneCapabilities2KHR                                 StructureType = 1000121004
	StructureTypeIosSurfaceCreateInfoMvk                                      StructureType = 1000122000
	StructureTypeMacosSurfaceCreateInfoMvk                                    StructureType = 1000123000
	StructureTypeMemoryDedicatedRequirementsKHR                               StructureType = StructureTypeMemoryDedicatedRequirements
	StructureTypeMemoryDedicatedAllocateInfoKHR                               StructureType = StructureTypeMemoryDedicatedAllocateInfo
	StructureTypeDebugUtilsObjectNameInfoEXT                                  StructureType = 1000128000
	StructureTypeDebugUtilsObjectTagInfoEXT                                   StructureType = 1000128001
	StructureTypeDebugUtilsLabelEXT                                           StructureType = 1000128002
	StructureTypeDebugUtilsMessengerCallbackDataEXT                           StructureType = 1000128003
	StructureTypeDebugUtilsMessengerCreateInfoEXT                             StructureType = 1000128004
	StructureTypeAndroidHardwareBufferUsageAndroid                            StructureType = 1000129000
	StructureTypeAndroidHardwareBufferPropertiesAndroid                       StructureType = 1000129001
	StructureTypeAndroidHardwareBufferFormatPropertiesAndroid                 StructureType = 1000129002
	StructureTypeImportAndroidHardwareBufferInfoAndroid                       StructureType = 1000129003
	StructureTypeMemoryGetAndroidHardwareBufferInfoAndroid                    StructureType = 1000129004
	StructureTypeExternalFormatAndroid                                        StructureType = 1000129005
	StructureTypeAndroidHardwareBufferFormatProperties2Android                StructureType = 1000129006
	StructureTypePhysicalDeviceSamplerFilterMinmaxPropertiesEXT               StructureType = StructureTypePhysicalDeviceSamplerFilterMinmaxProperties
	StructureTypeSamplerReductionModeCreateInfoEXT                            StructureType = StructureTypeSamplerReductionModeCreateInfo
	StructureTypePhysicalDeviceShaderEnqueueFeaturesAmdx                      StructureType = 1000134000
	StructureTypePhysicalDeviceShaderEnqueuePropertiesAmdx                    StructureType = 1000134001
	StructureTypeExecutionGraphPipelineScratchSizeAmdx                        StructureType = 1000134002
	StructureTypeExecutionGraphPipelineCreateInfoAmdx                         StructureType = 1000134003
	StructureTypePipelineShaderStageNodeCreateInfoAmdx                        StructureType = 1000134004
	StructureTypeTexelBufferDescriptorInfoEXT                                 StructureType = 1000135000
	StructureTypeImageDescriptorInfoEXT                                       StructureType = 1000135001
	StructureTypeResourceDescriptorInfoEXT                                    StructureType = 1000135002
	StructureTypeBindHeapInfoEXT                                              StructureType = 1000135003
	StructureTypePushDataInfoEXT                                              StructureType = 1000135004
	StructureTypeDescriptorSetAndBindingMappingEXT                            StructureType = 1000135005
	StructureTypeShaderDescriptorSetAndBindingMappingInfoEXT                  StructureType = 1000135006
	StructureTypeOpaqueCaptureDataCreateInfoEXT                               StructureType = 1000135007
	StructureTypePhysicalDeviceDescriptorHeapPropertiesEXT                    StructureType = 1000135008
	StructureTypePhysicalDeviceDescriptorHeapFeaturesEXT                      StructureType = 1000135009
	StructureTypeCommandBufferInheritanceDescriptorHeapInfoEXT                StructureType = 1000135010
	StructureTypeSamplerCustomBorderColorIndexCreateInfoEXT                   StructureType = 1000135011
	StructureTypeIndirectCommandsLayoutPushDataTokenNV                        StructureType = 1000135012
	StructureTypeSubsampledImageFormatPropertiesEXT                           StructureType = 1000135013
	StructureTypePhysicalDeviceDescriptorHeapTensorPropertiesARM              StructureType = 1000135014
	StructureTypeAttachmentSampleCountInfoAMD                                 StructureType = 1000044008
	StructureTypePhysicalDeviceInlineUniformBlockFeaturesEXT                  StructureType = StructureTypePhysicalDeviceInlineUniformBlockFeatures
	StructureTypePhysicalDeviceInlineUniformBlockPropertiesEXT                StructureType = StructureTypePhysicalDeviceInlineUniformBlockProperties
	StructureTypeWriteDescriptorSetInlineUniformBlockEXT                      StructureType = StructureTypeWriteDescriptorSetInlineUniformBlock
	StructureTypeDescriptorPoolInlineUniformBlockCreateInfoEXT                StructureType = StructureTypeDescriptorPoolInlineUniformBlockCreateInfo
	StructureTypePhysicalDeviceShaderBfloat16FeaturesKHR                      StructureType = 1000141000
	StructureTypeSampleLocationsInfoEXT                                       StructureType = 1000143000
	StructureTypeRenderPassSampleLocationsBeginInfoEXT                        StructureType = 1000143001
	StructureTypePipelineSampleLocationsStateCreateInfoEXT                    StructureType = 1000143002
	StructureTypePhysicalDeviceSampleLocationsPropertiesEXT                   StructureType = 1000143003
	StructureTypeMultisamplePropertiesEXT                                     StructureType = 1000143004
	StructureTypeBufferMemoryRequirementsInfo2KHR                             StructureType = StructureTypeBufferMemoryRequirementsInfo2
	StructureTypeImageMemoryRequirementsInfo2KHR                              StructureType = StructureTypeImageMemoryRequirementsInfo2
	StructureTypeImageSparseMemoryRequirementsInfo2KHR                        StructureType = StructureTypeImageSparseMemoryRequirementsInfo2
	StructureTypeMemoryRequirements2KHR                                       StructureType = StructureTypeMemoryRequirements2
	StructureTypeSparseImageMemoryRequirements2KHR                            StructureType = StructureTypeSparseImageMemoryRequirements2
	StructureTypeImageFormatListCreateInfoKHR                                 StructureType = StructureTypeImageFormatListCreateInfo
	StructureTypePhysicalDeviceBlendOperationAdvancedFeaturesEXT              StructureType = 1000148000
	StructureTypePhysicalDeviceBlendOperationAdvancedPropertiesEXT            StructureType = 1000148001
	StructureTypePipelineColorBlendAdvancedStateCreateInfoEXT                 StructureType = 1000148002
	StructureTypePipelineCoverageToColorStateCreateInfoNV                     StructureType = 1000149000
	StructureTypeWriteDescriptorSetAccelerationStructureKHR                   StructureType = 1000150007
	StructureTypeAccelerationStructureBuildGeometryInfoKHR                    StructureType = 1000150000
	StructureTypeAccelerationStructureDeviceAddressInfoKHR                    StructureType = 1000150002
	StructureTypeAccelerationStructureGeometryAabbsDataKHR                    StructureType = 1000150003
	StructureTypeAccelerationStructureGeometryInstancesDataKHR                StructureType = 1000150004
	StructureTypeAccelerationStructureGeometryTrianglesDataKHR                StructureType = 1000150005
	StructureTypeAccelerationStructureGeometryKHR                             StructureType = 1000150006
	StructureTypeAccelerationStructureVersionInfoKHR                          StructureType = 1000150009
	StructureTypeCopyAccelerationStructureInfoKHR                             StructureType = 1000150010
	StructureTypeCopyAccelerationStructureToMemoryInfoKHR                     StructureType = 1000150011
	StructureTypeCopyMemoryToAccelerationStructureInfoKHR                     StructureType = 1000150012
	StructureTypePhysicalDeviceAccelerationStructureFeaturesKHR               StructureType = 1000150013
	StructureTypePhysicalDeviceAccelerationStructurePropertiesKHR             StructureType = 1000150014
	StructureTypeAccelerationStructureCreateInfoKHR                           StructureType = 1000150017
	StructureTypeAccelerationStructureBuildSizesInfoKHR                       StructureType = 1000150020
	StructureTypePhysicalDeviceRayTracingPipelineFeaturesKHR                  StructureType = 1000151000
	StructureTypePhysicalDeviceRayTracingPipelinePropertiesKHR                StructureType = 1000151001
	StructureTypeRayTracingPipelineCreateInfoKHR                              StructureType = 1000150015
	StructureTypeRayTracingShaderGroupCreateInfoKHR                           StructureType = 1000150016
	StructureTypeRayTracingPipelineInterfaceCreateInfoKHR                     StructureType = 1000150018
	StructureTypePhysicalDeviceRayQueryFeaturesKHR                            StructureType = 1000152013
	StructureTypePipelineCoverageModulationStateCreateInfoNV                  StructureType = 1000154000
	StructureTypeAttachmentSampleCountInfoNV                                  StructureType = StructureTypeAttachmentSampleCountInfoAMD
	StructureTypePhysicalDeviceShaderSmBuiltinsFeaturesNV                     StructureType = 1000156000
	StructureTypePhysicalDeviceShaderSmBuiltinsPropertiesNV                   StructureType = 1000156001
	StructureTypeSamplerYcbcrConversionCreateInfoKHR                          StructureType = StructureTypeSamplerYcbcrConversionCreateInfo
	StructureTypeSamplerYcbcrConversionInfoKHR                                StructureType = StructureTypeSamplerYcbcrConversionInfo
	StructureTypeBindImagePlaneMemoryInfoKHR                                  StructureType = StructureTypeBindImagePlaneMemoryInfo
	StructureTypeImagePlaneMemoryRequirementsInfoKHR                          StructureType = StructureTypeImagePlaneMemoryRequirementsInfo
	StructureTypePhysicalDeviceSamplerYcbcrConversionFeaturesKHR              StructureType = StructureTypePhysicalDeviceSamplerYcbcrConversionFeatures
	StructureTypeSamplerYcbcrConversionImageFormatPropertiesKHR               StructureType = StructureTypeSamplerYcbcrConversionImageFormatProperties
	StructureTypeBindBufferMemoryInfoKHR                                      StructureType = StructureTypeBindBufferMemoryInfo
	StructureTypeBindImageMemoryInfoKHR                                       StructureType = StructureTypeBindImageMemoryInfo
	StructureTypeDrmFormatModifierPropertiesListEXT                           StructureType = 1000160000
	StructureTypePhysicalDeviceImageDrmFormatModifierInfoEXT                  StructureType = 1000160002
	StructureTypeImageDrmFormatModifierListCreateInfoEXT                      StructureType = 1000160003
	StructureTypeImageDrmFormatModifierExplicitCreateInfoEXT                  StructureType = 1000160004
	StructureTypeImageDrmFormatModifierPropertiesEXT                          StructureType = 1000160005
	StructureTypeDrmFormatModifierPropertiesList2EXT                          StructureType = 1000160006
	StructureTypeValidationCacheCreateInfoEXT                                 StructureType = 1000162000
	StructureTypeShaderModuleValidationCacheCreateInfoEXT                     StructureType = 1000162001
	StructureTypeDescriptorSetLayoutBindingFlagsCreateInfoEXT                 StructureType = StructureTypeDescriptorSetLayoutBindingFlagsCreateInfo
	StructureTypePhysicalDeviceDescriptorIndexingFeaturesEXT                  StructureType = StructureTypePhysicalDeviceDescriptorIndexingFeatures
	StructureTypePhysicalDeviceDescriptorIndexingPropertiesEXT                StructureType = StructureTypePhysicalDeviceDescriptorIndexingProperties
	StructureTypeDescriptorSetVariableDescriptorCountAllocateInfoEXT          StructureType = StructureTypeDescriptorSetVariableDescriptorCountAllocateInfo
	StructureTypeDescriptorSetVariableDescriptorCountLayoutSupportEXT         StructureType = StructureTypeDescriptorSetVariableDescriptorCountLayoutSupport
	StructureTypePhysicalDevicePortabilitySubsetFeaturesKHR                   StructureType = 1000165000
	StructureTypePhysicalDevicePortabilitySubsetPropertiesKHR                 StructureType = 1000165001
	StructureTypePipelineViewportShadingRateImageStateCreateInfoNV            StructureType = 1000166000
	StructureTypePhysicalDeviceShadingRateImageFeaturesNV                     StructureType = 1000166001
	StructureTypePhysicalDeviceShadingRateImagePropertiesNV                   StructureType = 1000166002
	StructureTypePipelineViewportCoarseSampleOrderStateCreateInfoNV           StructureType = 1000166005
	StructureTypeRayTracingPipelineCreateInfoNV                               StructureType = 1000167000
	StructureTypeAccelerationStructureCreateInfoNV                            StructureType = 1000167001
	StructureTypeGeometryNV                                                   StructureType = 1000167003
	StructureTypeGeometryTrianglesNV                                          StructureType = 1000167004
	StructureTypeGeometryAabbNV                                               StructureType = 1000167005
	StructureTypeBindAccelerationStructureMemoryInfoNV                        StructureType = 1000167006
	StructureTypeWriteDescriptorSetAccelerationStructureNV                    StructureType = 1000167007
	StructureTypeAccelerationStructureMemoryRequirementsInfoNV                StructureType = 1000167008
	StructureTypePhysicalDeviceRayTracingPropertiesNV                         StructureType = 1000167009
	StructureTypeRayTracingShaderGroupCreateInfoNV                            StructureType = 1000167011
	StructureTypeAccelerationStructureInfoNV                                  StructureType = 1000167012
	StructureTypePhysicalDeviceRepresentativeFragmentTestFeaturesNV           StructureType = 1000168000
	StructureTypePipelineRepresentativeFragmentTestStateCreateInfoNV          StructureType = 1000168001
	StructureTypePhysicalDeviceMaintenance3PropertiesKHR                      StructureType = StructureTypePhysicalDeviceMaintenance3Properties
	StructureTypeDescriptorSetLayoutSupportKHR                                StructureType = StructureTypeDescriptorSetLayoutSupport
	StructureTypePhysicalDeviceImageViewImageFormatInfoEXT                    StructureType = 1000172000
	StructureTypeFilterCubicImageViewImageFormatPropertiesEXT                 StructureType = 1000172001
	StructureTypePhysicalDeviceCooperativeMatrixConversionFeaturesQCOM        StructureType = 1000174000
	StructureTypeDeviceQueueGlobalPriorityCreateInfoEXT                       StructureType = StructureTypeDeviceQueueGlobalPriorityCreateInfo
	StructureTypePhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR         StructureType = StructureTypePhysicalDeviceShaderSubgroupExtendedTypesFeatures
	StructureTypePhysicalDevice8bitStorageFeaturesKHR                         StructureType = StructureTypePhysicalDevice8bitStorageFeatures
	StructureTypeImportMemoryHostPointerInfoEXT                               StructureType = 1000180000
	StructureTypeMemoryHostPointerPropertiesEXT                               StructureType = 1000180001
	StructureTypePhysicalDeviceExternalMemoryHostPropertiesEXT                StructureType = 1000180002
	StructureTypePhysicalDeviceShaderAtomicInt64FeaturesKHR                   StructureType = StructureTypePhysicalDeviceShaderAtomicInt64Features
	StructureTypePhysicalDeviceShaderClockFeaturesKHR                         StructureType = 1000183000
	StructureTypePipelineCompilerControlCreateInfoAMD                         StructureType = 1000185000
	StructureTypeCalibratedTimestampInfoEXT                                   StructureType = StructureTypeCalibratedTimestampInfoKHR
	StructureTypePhysicalDeviceShaderCorePropertiesAMD                        StructureType = 1000187000
	StructureTypeVideoDecodeH265CapabilitiesKHR                               StructureType = 1000189000
	StructureTypeVideoDecodeH265SessionParametersCreateInfoKHR                StructureType = 1000189001
	StructureTypeVideoDecodeH265SessionParametersAddInfoKHR                   StructureType = 1000189002
	StructureTypeVideoDecodeH265ProfileInfoKHR                                StructureType = 1000189003
	StructureTypeVideoDecodeH265PictureInfoKHR                                StructureType = 1000189004
	StructureTypeVideoDecodeH265DpbSlotInfoKHR                                StructureType = 1000189005
	StructureTypeDeviceQueueGlobalPriorityCreateInfoKHR                       StructureType = StructureTypeDeviceQueueGlobalPriorityCreateInfo
	StructureTypePhysicalDeviceGlobalPriorityQueryFeaturesKHR                 StructureType = StructureTypePhysicalDeviceGlobalPriorityQueryFeatures
	StructureTypeQueueFamilyGlobalPriorityPropertiesKHR                       StructureType = StructureTypeQueueFamilyGlobalPriorityProperties
	StructureTypeDeviceMemoryOverallocationCreateInfoAMD                      StructureType = 1000191000
	StructureTypePhysicalDeviceVertexAttributeDivisorPropertiesEXT            StructureType = 1000192000
	StructureTypePipelineVertexInputDivisorStateCreateInfoEXT                 StructureType = StructureTypePipelineVertexInputDivisorStateCreateInfo
	StructureTypePhysicalDeviceVertexAttributeDivisorFeaturesEXT              StructureType = StructureTypePhysicalDeviceVertexAttributeDivisorFeatures
	StructureTypePresentFrameTokenGgp                                         StructureType = 1000193000
	StructureTypePipelineCreationFeedbackCreateInfoEXT                        StructureType = StructureTypePipelineCreationFeedbackCreateInfo
	StructureTypePhysicalDeviceDriverPropertiesKHR                            StructureType = StructureTypePhysicalDeviceDriverProperties
	StructureTypePhysicalDeviceFloatControlsPropertiesKHR                     StructureType = StructureTypePhysicalDeviceFloatControlsProperties
	StructureTypePhysicalDeviceDepthStencilResolvePropertiesKHR               StructureType = StructureTypePhysicalDeviceDepthStencilResolveProperties
	StructureTypeSubpassDescriptionDepthStencilResolveKHR                     StructureType = StructureTypeSubpassDescriptionDepthStencilResolve
	StructureTypePhysicalDeviceComputeShaderDerivativesFeaturesNV             StructureType = StructureTypePhysicalDeviceComputeShaderDerivativesFeaturesKHR
	StructureTypePhysicalDeviceMeshShaderFeaturesNV                           StructureType = 1000204000
	StructureTypePhysicalDeviceMeshShaderPropertiesNV                         StructureType = 1000204001
	StructureTypePhysicalDeviceFragmentShaderBarycentricFeaturesNV            StructureType = StructureTypePhysicalDeviceFragmentShaderBarycentricFeaturesKHR
	StructureTypePhysicalDeviceShaderImageFootprintFeaturesNV                 StructureType = 1000206000
	StructureTypePipelineViewportExclusiveScissorStateCreateInfoNV            StructureType = 1000207000
	StructureTypePhysicalDeviceExclusiveScissorFeaturesNV                     StructureType = 1000207002
	StructureTypeCheckpointDataNV                                             StructureType = 1000208000
	StructureTypeQueueFamilyCheckpointPropertiesNV                            StructureType = 1000208001
	StructureTypeQueueFamilyCheckpointProperties2NV                           StructureType = 1000314008
	StructureTypeCheckpointData2NV                                            StructureType = 1000314009
	StructureTypePhysicalDeviceTimelineSemaphoreFeaturesKHR                   StructureType = StructureTypePhysicalDeviceTimelineSemaphoreFeatures
	StructureTypePhysicalDeviceTimelineSemaphorePropertiesKHR                 StructureType = StructureTypePhysicalDeviceTimelineSemaphoreProperties
	StructureTypeSemaphoreTypeCreateInfoKHR                                   StructureType = StructureTypeSemaphoreTypeCreateInfo
	StructureTypeTimelineSemaphoreSubmitInfoKHR                               StructureType = StructureTypeTimelineSemaphoreSubmitInfo
	StructureTypeSemaphoreWaitInfoKHR                                         StructureType = StructureTypeSemaphoreWaitInfo
	StructureTypeSemaphoreSignalInfoKHR                                       StructureType = StructureTypeSemaphoreSignalInfo
	StructureTypePhysicalDevicePresentTimingFeaturesEXT                       StructureType = 1000210000
	StructureTypeSwapchainTimingPropertiesEXT                                 StructureType = 1000210001
	StructureTypeSwapchainTimeDomainPropertiesEXT                             StructureType = 1000210002
	StructureTypePresentTimingsInfoEXT                                        StructureType = 1000210003
	StructureTypePresentTimingInfoEXT                                         StructureType = 1000210004
	StructureTypePastPresentationTimingInfoEXT                                StructureType = 1000210005
	StructureTypePastPresentationTimingPropertiesEXT                          StructureType = 1000210006
	StructureTypePastPresentationTimingEXT                                    StructureType = 1000210007
	StructureTypePresentTimingSurfaceCapabilitiesEXT                          StructureType = 1000210008
	StructureTypeSwapchainCalibratedTimestampInfoEXT                          StructureType = 1000210009
	StructureTypePhysicalDeviceShaderIntegerFunctions2FeaturesINTEL           StructureType = 1000211000
	StructureTypeQueryPoolPerformanceQueryCreateInfoINTEL                     StructureType = 1000212000
	StructureTypeQueryPoolCreateInfoINTEL                                     StructureType = StructureTypeQueryPoolPerformanceQueryCreateInfoINTEL
	StructureTypeInitializePerformanceApiInfoINTEL                            StructureType = 1000212001
	StructureTypePerformanceMarkerInfoINTEL                                   StructureType = 1000212002
	StructureTypePerformanceStreamMarkerInfoINTEL                             StructureType = 1000212003
	StructureTypePerformanceOverrideInfoINTEL                                 StructureType = 1000212004
	StructureTypePerformanceConfigurationAcquireInfoINTEL                     StructureType = 1000212005
	StructureTypePhysicalDeviceVulkanMemoryModelFeaturesKHR                   StructureType = StructureTypePhysicalDeviceVulkanMemoryModelFeatures
	StructureTypePhysicalDevicePciBusInfoPropertiesEXT                        StructureType = 1000214000
	StructureTypeDisplayNativeHdrSurfaceCapabilitiesAMD                       StructureType = 1000215000
	StructureTypeSwapchainDisplayNativeHdrCreateInfoAMD                       StructureType = 1000215001
	StructureTypeImagepipeSurfaceCreateInfoFuchsia                            StructureType = 1000216000
	StructureTypePhysicalDeviceShaderTerminateInvocationFeaturesKHR           StructureType = StructureTypePhysicalDeviceShaderTerminateInvocationFeatures
	StructureTypeMetalSurfaceCreateInfoEXT                                    StructureType = 1000219000
	StructureTypePhysicalDeviceFragmentDensityMapFeaturesEXT                  StructureType = 1000220000
	StructureTypePhysicalDeviceFragmentDensityMapPropertiesEXT                StructureType = 1000220001
	StructureTypeRenderPassFragmentDensityMapCreateInfoEXT                    StructureType = 1000220002
	StructureTypeRenderingFragmentDensityMapAttachmentInfoEXT                 StructureType = 1000044007
	StructureTypePhysicalDeviceScalarBlockLayoutFeaturesEXT                   StructureType = StructureTypePhysicalDeviceScalarBlockLayoutFeatures
	StructureTypePhysicalDeviceSubgroupSizeControlPropertiesEXT               StructureType = StructureTypePhysicalDeviceSubgroupSizeControlProperties
	StructureTypePipelineShaderStageRequiredSubgroupSizeCreateInfoEXT         StructureType = StructureTypePipelineShaderStageRequiredSubgroupSizeCreateInfo
	StructureTypePhysicalDeviceSubgroupSizeControlFeaturesEXT                 StructureType = StructureTypePhysicalDeviceSubgroupSizeControlFeatures
	StructureTypeFragmentShadingRateAttachmentInfoKHR                         StructureType = 1000228000
	StructureTypePipelineFragmentShadingRateStateCreateInfoKHR                StructureType = 1000228001
	StructureTypePhysicalDeviceFragmentShadingRatePropertiesKHR               StructureType = 1000228002
	StructureTypePhysicalDeviceFragmentShadingRateFeaturesKHR                 StructureType = 1000228003
	StructureTypePhysicalDeviceFragmentShadingRateKHR                         StructureType = 1000228004
	StructureTypeRenderingFragmentShadingRateAttachmentInfoKHR                StructureType = 1000044006
	StructureTypePhysicalDeviceShaderCoreProperties2AMD                       StructureType = 1000229000
	StructureTypePhysicalDeviceCoherentMemoryFeaturesAMD                      StructureType = 1000231000
	StructureTypePhysicalDeviceDynamicRenderingLocalReadFeaturesKHR           StructureType = StructureTypePhysicalDeviceDynamicRenderingLocalReadFeatures
	StructureTypeRenderingAttachmentLocationInfoKHR                           StructureType = StructureTypeRenderingAttachmentLocationInfo
	StructureTypeRenderingInputAttachmentIndexInfoKHR                         StructureType = StructureTypeRenderingInputAttachmentIndexInfo
	StructureTypePhysicalDeviceShaderImageAtomicInt64FeaturesEXT              StructureType = 1000236000
	StructureTypePhysicalDeviceShaderQuadControlFeaturesKHR                   StructureType = 1000237000
	StructureTypePhysicalDeviceMemoryBudgetPropertiesEXT                      StructureType = 1000239000
	StructureTypePhysicalDeviceMemoryPriorityFeaturesEXT                      StructureType = 1000240000
	StructureTypeMemoryPriorityAllocateInfoEXT                                StructureType = 1000240001
	StructureTypeSurfaceProtectedCapabilitiesKHR                              StructureType = 1000241000
	StructureTypePhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV     StructureType = 1000242000
	StructureTypePhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR         StructureType = StructureTypePhysicalDeviceSeparateDepthStencilLayoutsFeatures
	StructureTypeAttachmentReferenceStencilLayoutKHR                          StructureType = StructureTypeAttachmentReferenceStencilLayout
	StructureTypeAttachmentDescriptionStencilLayoutKHR                        StructureType = StructureTypeAttachmentDescriptionStencilLayout
	StructureTypePhysicalDeviceBufferDeviceAddressFeaturesEXT                 StructureType = 1000246000
	StructureTypePhysicalDeviceBufferAddressFeaturesEXT                       StructureType = StructureTypePhysicalDeviceBufferDeviceAddressFeaturesEXT
	StructureTypeBufferDeviceAddressInfoEXT                                   StructureType = StructureTypeBufferDeviceAddressInfo
	StructureTypeBufferDeviceAddressCreateInfoEXT                             StructureType = 1000246002
	StructureTypePhysicalDeviceToolPropertiesEXT                              StructureType = StructureTypePhysicalDeviceToolProperties
	StructureTypeImageStencilUsageCreateInfoEXT                               StructureType = StructureTypeImageStencilUsageCreateInfo
	StructureTypeValidationFeaturesEXT                                        StructureType = 1000249000
	StructureTypePhysicalDevicePresentWaitFeaturesKHR                         StructureType = 1000250000
	StructureTypePhysicalDeviceCooperativeMatrixFeaturesNV                    StructureType = 1000251000
	StructureTypeCooperativeMatrixPropertiesNV                                StructureType = 1000251001
	StructureTypePhysicalDeviceCooperativeMatrixPropertiesNV                  StructureType = 1000251002
	StructureTypePhysicalDeviceCoverageReductionModeFeaturesNV                StructureType = 1000252000
	StructureTypePipelineCoverageReductionStateCreateInfoNV                   StructureType = 1000252001
	StructureTypeFramebufferMixedSamplesCombinationNV                         StructureType = 1000252002
	StructureTypePhysicalDeviceFragmentShaderInterlockFeaturesEXT             StructureType = 1000253000
	StructureTypePhysicalDeviceYcbcrImageArraysFeaturesEXT                    StructureType = 1000254000
	StructureTypePhysicalDeviceUniformBufferStandardLayoutFeaturesKHR         StructureType = StructureTypePhysicalDeviceUniformBufferStandardLayoutFeatures
	StructureTypePhysicalDeviceProvokingVertexFeaturesEXT                     StructureType = 1000256000
	StructureTypePipelineRasterizationProvokingVertexStateCreateInfoEXT       StructureType = 1000256001
	StructureTypePhysicalDeviceProvokingVertexPropertiesEXT                   StructureType = 1000256002
	StructureTypeSurfaceFullScreenExclusiveInfoEXT                            StructureType = 1000257000
	StructureTypeSurfaceCapabilitiesFullScreenExclusiveEXT                    StructureType = 1000257002
	StructureTypeSurfaceFullScreenExclusiveWin32InfoEXT                       StructureType = 1000257001
	StructureTypeHeadlessSurfaceCreateInfoEXT                                 StructureType = 1000258000
	StructureTypePhysicalDeviceBufferDeviceAddressFeaturesKHR                 StructureType = StructureTypePhysicalDeviceBufferDeviceAddressFeatures
	StructureTypeBufferDeviceAddressInfoKHR                                   StructureType = StructureTypeBufferDeviceAddressInfo
	StructureTypeBufferOpaqueCaptureAddressCreateInfoKHR                      StructureType = StructureTypeBufferOpaqueCaptureAddressCreateInfo
	StructureTypeMemoryOpaqueCaptureAddressAllocateInfoKHR                    StructureType = StructureTypeMemoryOpaqueCaptureAddressAllocateInfo
	StructureTypeDeviceMemoryOpaqueCaptureAddressInfoKHR                      StructureType = StructureTypeDeviceMemoryOpaqueCaptureAddressInfo
	StructureTypePhysicalDeviceLineRasterizationFeaturesEXT                   StructureType = StructureTypePhysicalDeviceLineRasterizationFeatures
	StructureTypePipelineRasterizationLineStateCreateInfoEXT                  StructureType = StructureTypePipelineRasterizationLineStateCreateInfo
	StructureTypePhysicalDeviceLineRasterizationPropertiesEXT                 StructureType = StructureTypePhysicalDeviceLineRasterizationProperties
	StructureTypePhysicalDeviceShaderAtomicFloatFeaturesEXT                   StructureType = 1000262000
	StructureTypePhysicalDeviceHostQueryResetFeaturesEXT                      StructureType = StructureTypePhysicalDeviceHostQueryResetFeatures
	StructureTypePhysicalDeviceIndexTypeUint8FeaturesEXT                      StructureType = StructureTypePhysicalDeviceIndexTypeUint8Features
	StructureTypePhysicalDeviceExtendedDynamicStateFeaturesEXT                StructureType = 1000269000
	StructureTypePhysicalDevicePipelineExecutablePropertiesFeaturesKHR        StructureType = 1000271000
	StructureTypePipelineInfoKHR                                              StructureType = 1000271001
	StructureTypePipelineExecutablePropertiesKHR                              StructureType = 1000271002
	StructureTypePipelineExecutableInfoKHR                                    StructureType = 1000271003
	StructureTypePipelineExecutableStatisticKHR                               StructureType = 1000271004
	StructureTypePipelineExecutableInternalRepresentationKHR                  StructureType = 1000271005
	StructureTypePhysicalDeviceHostImageCopyFeaturesEXT                       StructureType = StructureTypePhysicalDeviceHostImageCopyFeatures
	StructureTypePhysicalDeviceHostImageCopyPropertiesEXT                     StructureType = StructureTypePhysicalDeviceHostImageCopyProperties
	StructureTypeMemoryToImageCopyEXT                                         StructureType = StructureTypeMemoryToImageCopy
	StructureTypeImageToMemoryCopyEXT                                         StructureType = StructureTypeImageToMemoryCopy
	StructureTypeCopyImageToMemoryInfoEXT                                     StructureType = StructureTypeCopyImageToMemoryInfo
	StructureTypeCopyMemoryToImageInfoEXT                                     StructureType = StructureTypeCopyMemoryToImageInfo
	StructureTypeHostImageLayoutTransitionInfoEXT                             StructureType = StructureTypeHostImageLayoutTransitionInfo
	StructureTypeCopyImageToImageInfoEXT                                      StructureType = StructureTypeCopyImageToImageInfo
	StructureTypeSubresourceHostMemcpySizeEXT                                 StructureType = StructureTypeSubresourceHostMemcpySize
	StructureTypeHostImageCopyDevicePerformanceQueryEXT                       StructureType = StructureTypeHostImageCopyDevicePerformanceQuery
	StructureTypeMemoryMapInfoKHR                                             StructureType = StructureTypeMemoryMapInfo
	StructureTypeMemoryUnmapInfoKHR                                           StructureType = StructureTypeMemoryUnmapInfo
	StructureTypePhysicalDeviceMapMemoryPlacedFeaturesEXT                     StructureType = 1000274000
	StructureTypePhysicalDeviceMapMemoryPlacedPropertiesEXT                   StructureType = 1000274001
	StructureTypeMemoryMapPlacedInfoEXT                                       StructureType = 1000274002
	StructureTypePhysicalDeviceShaderAtomicFloat2FeaturesEXT                  StructureType = 1000275000
	StructureTypeSurfacePresentModeEXT                                        StructureType = StructureTypeSurfacePresentModeKHR
	StructureTypeSurfacePresentScalingCapabilitiesEXT                         StructureType = StructureTypeSurfacePresentScalingCapabilitiesKHR
	StructureTypeSurfacePresentModeCompatibilityEXT                           StructureType = StructureTypeSurfacePresentModeCompatibilityKHR
	StructureTypePhysicalDeviceSwapchainMaintenance1FeaturesEXT               StructureType = StructureTypePhysicalDeviceSwapchainMaintenance1FeaturesKHR
	StructureTypeSwapchainPresentFenceInfoEXT                                 StructureType = StructureTypeSwapchainPresentFenceInfoKHR
	StructureTypeSwapchainPresentModesCreateInfoEXT                           StructureType = StructureTypeSwapchainPresentModesCreateInfoKHR
	StructureTypeSwapchainPresentModeInfoEXT                                  StructureType = StructureTypeSwapchainPresentModeInfoKHR
	StructureTypeSwapchainPresentScalingCreateInfoEXT                         StructureType = StructureTypeSwapchainPresentScalingCreateInfoKHR
	StructureTypeReleaseSwapchainImagesInfoEXT                                StructureType = StructureTypeReleaseSwapchainImagesInfoKHR
	StructureTypePhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT      StructureType = StructureTypePhysicalDeviceShaderDemoteToHelperInvocationFeatures
	StructureTypePhysicalDeviceDeviceGeneratedCommandsPropertiesNV            StructureType = 1000279000
	StructureTypeGraphicsShaderGroupCreateInfoNV                              StructureType = 1000279001
	StructureTypeGraphicsPipelineShaderGroupsCreateInfoNV                     StructureType = 1000279002
	StructureTypeIndirectCommandsLayoutTokenNV                                StructureType = 1000279003
	StructureTypeIndirectCommandsLayoutCreateInfoNV                           StructureType = 1000279004
	StructureTypeGeneratedCommandsInfoNV                                      StructureType = 1000279005
	StructureTypeGeneratedCommandsMemoryRequirementsInfoNV                    StructureType = 1000279006
	StructureTypePhysicalDeviceDeviceGeneratedCommandsFeaturesNV              StructureType = 1000279007
	StructureTypePhysicalDeviceInheritedViewportScissorFeaturesNV             StructureType = 1000280000
	StructureTypeCommandBufferInheritanceViewportScissorInfoNV                StructureType = 1000280001
	StructureTypePhysicalDeviceShaderIntegerDotProductFeaturesKHR             StructureType = StructureTypePhysicalDeviceShaderIntegerDotProductFeatures
	StructureTypePhysicalDeviceShaderIntegerDotProductPropertiesKHR           StructureType = StructureTypePhysicalDeviceShaderIntegerDotProductProperties
	StructureTypePhysicalDeviceTexelBufferAlignmentFeaturesEXT                StructureType = 1000283000
	StructureTypePhysicalDeviceTexelBufferAlignmentPropertiesEXT              StructureType = StructureTypePhysicalDeviceTexelBufferAlignmentProperties
	StructureTypeCommandBufferInheritanceRenderPassTransformInfoQCOM          StructureType = 1000284000
	StructureTypeRenderPassTransformBeginInfoQCOM                             StructureType = 1000284001
	StructureTypePhysicalDeviceDepthBiasControlFeaturesEXT                    StructureType = 1000285000
	StructureTypeDepthBiasInfoEXT                                             StructureType = 1000285001
	StructureTypeDepthBiasRepresentationInfoEXT                               StructureType = 1000285002
	StructureTypePhysicalDeviceDeviceMemoryReportFeaturesEXT                  StructureType = 1000286000
	StructureTypeDeviceDeviceMemoryReportCreateInfoEXT                        StructureType = 1000286001
	StructureTypeDeviceMemoryReportCallbackDataEXT                            StructureType = 1000286002
	StructureTypePhysicalDeviceRobustness2FeaturesEXT                         StructureType = StructureTypePhysicalDeviceRobustness2FeaturesKHR
	StructureTypePhysicalDeviceRobustness2PropertiesEXT                       StructureType = StructureTypePhysicalDeviceRobustness2PropertiesKHR
	StructureTypeSamplerCustomBorderColorCreateInfoEXT                        StructureType = 1000289000
	StructureTypePhysicalDeviceCustomBorderColorPropertiesEXT                 StructureType = 1000289001
	StructureTypePhysicalDeviceCustomBorderColorFeaturesEXT                   StructureType = 1000289002
	StructureTypePhysicalDeviceTextureCompressionAstc3dFeaturesEXT            StructureType = 1000290000
	StructureTypePipelineLibraryCreateInfoKHR                                 StructureType = 1000292000
	StructureTypePhysicalDevicePresentBarrierFeaturesNV                       StructureType = 1000294000
	StructureTypeSurfaceCapabilitiesPresentBarrierNV                          StructureType = 1000294001
	StructureTypeSwapchainPresentBarrierCreateInfoNV                          StructureType = 1000294002
	StructureTypePresentIdKHR                                                 StructureType = 1000296000
	StructureTypePhysicalDevicePresentIdFeaturesKHR                           StructureType = 1000296001
	StructureTypePhysicalDevicePrivateDataFeaturesEXT                         StructureType = StructureTypePhysicalDevicePrivateDataFeatures
	StructureTypeDevicePrivateDataCreateInfoEXT                               StructureType = StructureTypeDevicePrivateDataCreateInfo
	StructureTypePrivateDataSlotCreateInfoEXT                                 StructureType = StructureTypePrivateDataSlotCreateInfo
	StructureTypePhysicalDevicePipelineCreationCacheControlFeaturesEXT        StructureType = StructureTypePhysicalDevicePipelineCreationCacheControlFeatures
	StructureTypeVideoEncodeInfoKHR                                           StructureType = 1000301000
	StructureTypeVideoEncodeRateControlInfoKHR                                StructureType = 1000301001
	StructureTypeVideoEncodeRateControlLayerInfoKHR                           StructureType = 1000301002
	StructureTypeVideoEncodeCapabilitiesKHR                                   StructureType = 1000301003
	StructureTypeVideoEncodeUsageInfoKHR                                      StructureType = 1000301004
	StructureTypeQueryPoolVideoEncodeFeedbackCreateInfoKHR                    StructureType = 1000301005
	StructureTypePhysicalDeviceVideoEncodeQualityLevelInfoKHR                 StructureType = 1000301006
	StructureTypeVideoEncodeQualityLevelPropertiesKHR                         StructureType = 1000301007
	StructureTypeVideoEncodeQualityLevelInfoKHR                               StructureType = 1000301008
	StructureTypeVideoEncodeSessionParametersGetInfoKHR                       StructureType = 1000301009
	StructureTypeVideoEncodeSessionParametersFeedbackInfoKHR                  StructureType = 1000301010
	StructureTypePhysicalDeviceDiagnosticsConfigFeaturesNV                    StructureType = 1000302000
	StructureTypeDeviceDiagnosticsConfigCreateInfoNV                          StructureType = 1000302001
	StructureTypeCudaModuleCreateInfoNV                                       StructureType = 1000309000
	StructureTypeCudaFunctionCreateInfoNV                                     StructureType = 1000309001
	StructureTypeCudaLaunchInfoNV                                             StructureType = 1000309002
	StructureTypePhysicalDeviceCudaKernelLaunchFeaturesNV                     StructureType = 1000309003
	StructureTypePhysicalDeviceCudaKernelLaunchPropertiesNV                   StructureType = 1000309004
	StructureTypeRefreshObjectListKHR                                         StructureType = 1000310000
	StructureTypePhysicalDeviceTileShadingFeaturesQCOM                        StructureType = 1000311000
	StructureTypePhysicalDeviceTileShadingPropertiesQCOM                      StructureType = 1000311001
	StructureTypeRenderPassTileShadingCreateInfoQCOM                          StructureType = 1000311002
	StructureTypePerTileBeginInfoQCOM                                         StructureType = 1000311003
	StructureTypePerTileEndInfoQCOM                                           StructureType = 1000311004
	StructureTypeDispatchTileInfoQCOM                                         StructureType = 1000311005
	StructureTypeQueryLowLatencySupportNV                                     StructureType = 1000312000
	StructureTypeExportMetalObjectCreateInfoEXT                               StructureType = 1000313000
	StructureTypeExportMetalObjectsInfoEXT                                    StructureType = 1000313001
	StructureTypeExportMetalDeviceInfoEXT                                     StructureType = 1000313002
	StructureTypeExportMetalCommandQueueInfoEXT                               StructureType = 1000313003
	StructureTypeExportMetalBufferInfoEXT                                     StructureType = 1000313004
	StructureTypeImportMetalBufferInfoEXT                                     StructureType = 1000313005
	StructureTypeExportMetalTextureInfoEXT                                    StructureType = 1000313006
	StructureTypeImportMetalTextureInfoEXT                                    StructureType = 1000313007
	StructureTypeExportMetalIoSurfaceInfoEXT                                  StructureType = 1000313008
	StructureTypeImportMetalIoSurfaceInfoEXT                                  StructureType = 1000313009
	StructureTypeExportMetalSharedEventInfoEXT                                StructureType = 1000313010
	StructureTypeImportMetalSharedEventInfoEXT                                StructureType = 1000313011
	StructureTypeMemoryBarrier2KHR                                            StructureType = StructureTypeMemoryBarrier2
	StructureTypeBufferMemoryBarrier2KHR                                      StructureType = StructureTypeBufferMemoryBarrier2
	StructureTypeImageMemoryBarrier2KHR                                       StructureType = StructureTypeImageMemoryBarrier2
	StructureTypeDependencyInfoKHR                                            StructureType = StructureTypeDependencyInfo
	StructureTypeSubmitInfo2KHR                                               StructureType = StructureTypeSubmitInfo2
	StructureTypeSemaphoreSubmitInfoKHR                                       StructureType = StructureTypeSemaphoreSubmitInfo
	StructureTypeCommandBufferSubmitInfoKHR                                   StructureType = StructureTypeCommandBufferSubmitInfo
	StructureTypePhysicalDeviceSynchronization2FeaturesKHR                    StructureType = StructureTypePhysicalDeviceSynchronization2Features
	StructureTypePhysicalDeviceDescriptorBufferPropertiesEXT                  StructureType = 1000318000
	StructureTypePhysicalDeviceDescriptorBufferDensityMapPropertiesEXT        StructureType = 1000318001
	StructureTypePhysicalDeviceDescriptorBufferFeaturesEXT                    StructureType = 1000318002
	StructureTypeDescriptorAddressInfoEXT                                     StructureType = 1000318003
	StructureTypeDescriptorGetInfoEXT                                         StructureType = 1000318004
	StructureTypeBufferCaptureDescriptorDataInfoEXT                           StructureType = 1000318005
	StructureTypeImageCaptureDescriptorDataInfoEXT                            StructureType = 1000318006
	StructureTypeImageViewCaptureDescriptorDataInfoEXT                        StructureType = 1000318007
	StructureTypeSamplerCaptureDescriptorDataInfoEXT                          StructureType = 1000318008
	StructureTypeOpaqueCaptureDescriptorDataCreateInfoEXT                     StructureType = 1000318010
	StructureTypeDescriptorBufferBindingInfoEXT                               StructureType = 1000318011
	StructureTypeDescriptorBufferBindingPushDescriptorBufferHandleEXT         StructureType = 1000318012
	StructureTypeAccelerationStructureCaptureDescriptorDataInfoEXT            StructureType = 1000318009
	StructureTypePhysicalDeviceGraphicsPipelineLibraryFeaturesEXT             StructureType = 1000322000
	StructureTypePhysicalDeviceGraphicsPipelineLibraryPropertiesEXT           StructureType = 1000322001
	StructureTypeGraphicsPipelineLibraryCreateInfoEXT                         StructureType = 1000322002
	StructureTypePhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD     StructureType = 1000323000
	StructureTypePhysicalDeviceFragmentShaderBarycentricFeaturesKHR           StructureType = 1000203000
	StructureTypePhysicalDeviceFragmentShaderBarycentricPropertiesKHR         StructureType = 1000324000
	StructureTypePhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR    StructureType = 1000325000
	StructureTypePhysicalDeviceZeroInitializeWorkgroupMemoryFeaturesKHR       StructureType = StructureTypePhysicalDeviceZeroInitializeWorkgroupMemoryFeatures
	StructureTypePhysicalDeviceFragmentShadingRateEnumsPropertiesNV           StructureType = 1000328000
	StructureTypePhysicalDeviceFragmentShadingRateEnumsFeaturesNV             StructureType = 1000328001
	StructureTypePipelineFragmentShadingRateEnumStateCreateInfoNV             StructureType = 1000328002
	StructureTypeAccelerationStructureGeometryMotionTrianglesDataNV           StructureType = 1000329000
	StructureTypePhysicalDeviceRayTracingMotionBlurFeaturesNV                 StructureType = 1000329001
	StructureTypeAccelerationStructureMotionInfoNV                            StructureType = 1000329002
	StructureTypePhysicalDeviceMeshShaderFeaturesEXT                          StructureType = 1000330000
	StructureTypePhysicalDeviceMeshShaderPropertiesEXT                        StructureType = 1000330001
	StructureTypePhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT               StructureType = 1000332000
	StructureTypePhysicalDeviceFragmentDensityMap2FeaturesEXT                 StructureType = 1000334000
	StructureTypePhysicalDeviceFragmentDensityMap2PropertiesEXT               StructureType = 1000334001
	StructureTypeCopyCommandTransformInfoQCOM                                 StructureType = 1000335000
	StructureTypePhysicalDeviceImageRobustnessFeaturesEXT                     StructureType = StructureTypePhysicalDeviceImageRobustnessFeatures
	StructureTypePhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR       StructureType = 1000338000
	StructureTypeCopyBufferInfo2KHR                                           StructureType = StructureTypeCopyBufferInfo2
	StructureTypeCopyImageInfo2KHR                                            StructureType = StructureTypeCopyImageInfo2
	StructureTypeCopyBufferToImageInfo2KHR                                    StructureType = StructureTypeCopyBufferToImageInfo2
	StructureTypeCopyImageToBufferInfo2KHR                                    StructureType = StructureTypeCopyImageToBufferInfo2
	StructureTypeBlitImageInfo2KHR                                            StructureType = StructureTypeBlitImageInfo2
	StructureTypeResolveImageInfo2KHR                                         StructureType = StructureTypeResolveImageInfo2
	StructureTypeBufferCopy2KHR                                               StructureType = StructureTypeBufferCopy2
	StructureTypeImageCopy2KHR                                                StructureType = StructureTypeImageCopy2
	StructureTypeImageBlit2KHR                                                StructureType = StructureTypeImageBlit2
	StructureTypeBufferImageCopy2KHR                                          StructureType = StructureTypeBufferImageCopy2
	StructureTypeImageResolve2KHR                                             StructureType = StructureTypeImageResolve2
	StructureTypePhysicalDeviceImageCompressionControlFeaturesEXT             StructureType = 1000340000
	StructureTypeImageCompressionControlEXT                                   StructureType = 1000340001
	StructureTypeSubresourceLayout2EXT                                        StructureType = StructureTypeSubresourceLayout2
	StructureTypeImageSubresource2EXT                                         StructureType = StructureTypeImageSubresource2
	StructureTypeImageCompressionPropertiesEXT                                StructureType = 1000340004
	StructureTypePhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT        StructureType = 1000341000
	StructureTypePhysicalDevice4444FormatsFeaturesEXT                         StructureType = 1000342000
	StructureTypePhysicalDeviceFaultFeaturesEXT                               StructureType = 1000343000
	StructureTypeDeviceFaultCountsEXT                                         StructureType = 1000343001
	StructureTypeDeviceFaultInfoEXT                                           StructureType = 1000343002
	StructureTypePhysicalDeviceRasterizationOrderAttachmentAccessFeaturesARM  StructureType = StructureTypePhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT
	StructureTypePhysicalDeviceRgba10x6FormatsFeaturesEXT                     StructureType = 1000346000
	StructureTypeDirectfbSurfaceCreateInfoEXT                                 StructureType = 1000348000
	StructureTypePhysicalDeviceMutableDescriptorTypeFeaturesVALVE             StructureType = StructureTypePhysicalDeviceMutableDescriptorTypeFeaturesEXT
	StructureTypeMutableDescriptorTypeCreateInfoVALVE                         StructureType = StructureTypeMutableDescriptorTypeCreateInfoEXT
	StructureTypePhysicalDeviceVertexInputDynamicStateFeaturesEXT             StructureType = 1000352000
	StructureTypeVertexInputBindingDescription2EXT                            StructureType = 1000352001
	StructureTypeVertexInputAttributeDescription2EXT                          StructureType = 1000352002
	StructureTypePhysicalDeviceDrmPropertiesEXT                               StructureType = 1000353000
	StructureTypePhysicalDeviceAddressBindingReportFeaturesEXT                StructureType = 1000354000
	StructureTypeDeviceAddressBindingCallbackDataEXT                          StructureType = 1000354001
	StructureTypePhysicalDeviceDepthClipControlFeaturesEXT                    StructureType = 1000355000
	StructureTypePipelineViewportDepthClipControlCreateInfoEXT                StructureType = 1000355001
	StructureTypePhysicalDevicePrimitiveTopologyListRestartFeaturesEXT        StructureType = 1000356000
	StructureTypeFormatProperties3KHR                                         StructureType = StructureTypeFormatProperties3
	StructureTypePhysicalDevicePresentModeFifoLatestReadyFeaturesEXT          StructureType = StructureTypePhysicalDevicePresentModeFifoLatestReadyFeaturesKHR
	StructureTypeImportMemoryZirconHandleInfoFuchsia                          StructureType = 1000364000
	StructureTypeMemoryZirconHandlePropertiesFuchsia                          StructureType = 1000364001
	StructureTypeMemoryGetZirconHandleInfoFuchsia                             StructureType = 1000364002
	StructureTypeImportSemaphoreZirconHandleInfoFuchsia                       StructureType = 1000365000
	StructureTypeSemaphoreGetZirconHandleInfoFuchsia                          StructureType = 1000365001
	StructureTypeBufferCollectionCreateInfoFuchsia                            StructureType = 1000366000
	StructureTypeImportMemoryBufferCollectionFuchsia                          StructureType = 1000366001
	StructureTypeBufferCollectionImageCreateInfoFuchsia                       StructureType = 1000366002
	StructureTypeBufferCollectionPropertiesFuchsia                            StructureType = 1000366003
	StructureTypeBufferConstraintsInfoFuchsia                                 StructureType = 1000366004
	StructureTypeBufferCollectionBufferCreateInfoFuchsia                      StructureType = 1000366005
	StructureTypeImageConstraintsInfoFuchsia                                  StructureType = 1000366006
	StructureTypeImageFormatConstraintsInfoFuchsia                            StructureType = 1000366007
	StructureTypeSysmemColorSpaceFuchsia                                      StructureType = 1000366008
	StructureTypeBufferCollectionConstraintsInfoFuchsia                       StructureType = 1000366009
	StructureTypeSubpassShadingPipelineCreateInfoHUAWEI                       StructureType = 1000369000
	StructureTypePhysicalDeviceSubpassShadingFeaturesHUAWEI                   StructureType = 1000369001
	StructureTypePhysicalDeviceSubpassShadingPropertiesHUAWEI                 StructureType = 1000369002
	StructureTypePhysicalDeviceInvocationMaskFeaturesHUAWEI                   StructureType = 1000370000
	StructureTypeMemoryGetRemoteAddressInfoNV                                 StructureType = 1000371000
	StructureTypePhysicalDeviceExternalMemoryRdmaFeaturesNV                   StructureType = 1000371001
	StructureTypePipelinePropertiesIdentifierEXT                              StructureType = 1000372000
	StructureTypePhysicalDevicePipelinePropertiesFeaturesEXT                  StructureType = 1000372001
	StructureTypePipelineInfoEXT                                              StructureType = StructureTypePipelineInfoKHR
	StructureTypeImportFenceSciSyncInfoNV                                     StructureType = 1000373000
	StructureTypeExportFenceSciSyncInfoNV                                     StructureType = 1000373001
	StructureTypeFenceGetSciSyncInfoNV                                        StructureType = 1000373002
	StructureTypeSciSyncAttributesInfoNV                                      StructureType = 1000373003
	StructureTypeImportSemaphoreSciSyncInfoNV                                 StructureType = 1000373004
	StructureTypeExportSemaphoreSciSyncInfoNV                                 StructureType = 1000373005
	StructureTypeSemaphoreGetSciSyncInfoNV                                    StructureType = 1000373006
	StructureTypePhysicalDeviceExternalSciSyncFeaturesNV                      StructureType = 1000373007
	StructureTypeImportMemorySciBufInfoNV                                     StructureType = 1000374000
	StructureTypeExportMemorySciBufInfoNV                                     StructureType = 1000374001
	StructureTypeMemoryGetSciBufInfoNV                                        StructureType = 1000374002
	StructureTypeMemorySciBufPropertiesNV                                     StructureType = 1000374003
	StructureTypePhysicalDeviceExternalMemorySciBufFeaturesNV                 StructureType = 1000374004
	StructureTypePhysicalDeviceExternalSciBufFeaturesNV                       StructureType = StructureTypePhysicalDeviceExternalMemorySciBufFeaturesNV
	StructureTypePhysicalDeviceFrameBoundaryFeaturesEXT                       StructureType = 1000375000
	StructureTypeFrameBoundaryEXT                                             StructureType = 1000375001
	StructureTypePhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT   StructureType = 1000376000
	StructureTypeSubpassResolvePerformanceQueryEXT                            StructureType = 1000376001
	StructureTypeMultisampledRenderToSingleSampledInfoEXT                     StructureType = 1000376002
	StructureTypePhysicalDeviceExtendedDynamicState2FeaturesEXT               StructureType = 1000377000
	StructureTypeScreenSurfaceCreateInfoQnx                                   StructureType = 1000378000
	StructureTypePhysicalDeviceColorWriteEnableFeaturesEXT                    StructureType = 1000381000
	StructureTypePipelineColorWriteCreateInfoEXT                              StructureType = 1000381001
	StructureTypePhysicalDevicePrimitivesGeneratedQueryFeaturesEXT            StructureType = 1000382000
	StructureTypePhysicalDeviceRayTracingMaintenance1FeaturesKHR              StructureType = 1000386000
	StructureTypePhysicalDeviceShaderUntypedPointersFeaturesKHR               StructureType = 1000387000
	StructureTypePhysicalDeviceGlobalPriorityQueryFeaturesEXT                 StructureType = StructureTypePhysicalDeviceGlobalPriorityQueryFeatures
	StructureTypeQueueFamilyGlobalPriorityPropertiesEXT                       StructureType = StructureTypeQueueFamilyGlobalPriorityProperties
	StructureTypePhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE          StructureType = 1000390000
	StructureTypeVideoEncodeRgbConversionCapabilitiesVALVE                    StructureType = 1000390001
	StructureTypeVideoEncodeProfileRgbConversionInfoVALVE                     StructureType = 1000390002
	StructureTypeVideoEncodeSessionRgbConversionCreateInfoVALVE               StructureType = 1000390003
	StructureTypePhysicalDeviceImageViewMinLodFeaturesEXT                     StructureType = 1000391000
	StructureTypeImageViewMinLodCreateInfoEXT                                 StructureType = 1000391001
	StructureTypePhysicalDeviceMultiDrawFeaturesEXT                           StructureType = 1000392000
	StructureTypePhysicalDeviceMultiDrawPropertiesEXT                         StructureType = 1000392001
	StructureTypePhysicalDeviceImage2dViewOf3dFeaturesEXT                     StructureType = 1000393000
	StructureTypePhysicalDeviceShaderTileImageFeaturesEXT                     StructureType = 1000395000
	StructureTypePhysicalDeviceShaderTileImagePropertiesEXT                   StructureType = 1000395001
	StructureTypeMicromapBuildInfoEXT                                         StructureType = 1000396000
	StructureTypeMicromapVersionInfoEXT                                       StructureType = 1000396001
	StructureTypeCopyMicromapInfoEXT                                          StructureType = 1000396002
	StructureTypeCopyMicromapToMemoryInfoEXT                                  StructureType = 1000396003
	StructureTypeCopyMemoryToMicromapInfoEXT                                  StructureType = 1000396004
	StructureTypePhysicalDeviceOpacityMicromapFeaturesEXT                     StructureType = 1000396005
	StructureTypePhysicalDeviceOpacityMicromapPropertiesEXT                   StructureType = 1000396006
	StructureTypeMicromapCreateInfoEXT                                        StructureType = 1000396007
	StructureTypeMicromapBuildSizesInfoEXT                                    StructureType = 1000396008
	StructureTypeAccelerationStructureTrianglesOpacityMicromapEXT             StructureType = 1000396009
	StructureTypePhysicalDeviceDisplacementMicromapFeaturesNV                 StructureType = 1000397000
	StructureTypePhysicalDeviceDisplacementMicromapPropertiesNV               StructureType = 1000397001
	StructureTypeAccelerationStructureTrianglesDisplacementMicromapNV         StructureType = 1000397002
	StructureTypePhysicalDeviceClusterCullingShaderFeaturesHUAWEI             StructureType = 1000404000
	StructureTypePhysicalDeviceClusterCullingShaderPropertiesHUAWEI           StructureType = 1000404001
	StructureTypePhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI          StructureType = 1000404002
	StructureTypePhysicalDeviceBorderColorSwizzleFeaturesEXT                  StructureType = 1000411000
	StructureTypeSamplerBorderColorComponentMappingCreateInfoEXT              StructureType = 1000411001
	StructureTypePhysicalDevicePageableDeviceLocalMemoryFeaturesEXT           StructureType = 1000412000
	StructureTypePhysicalDeviceMaintenance4FeaturesKHR                        StructureType = StructureTypePhysicalDeviceMaintenance4Features
	StructureTypePhysicalDeviceMaintenance4PropertiesKHR                      StructureType = StructureTypePhysicalDeviceMaintenance4Properties
	StructureTypeDeviceBufferMemoryRequirementsKHR                            StructureType = StructureTypeDeviceBufferMemoryRequirements
	StructureTypeDeviceImageMemoryRequirementsKHR                             StructureType = StructureTypeDeviceImageMemoryRequirements
	StructureTypePhysicalDeviceShaderCorePropertiesARM                        StructureType = 1000415000
	StructureTypePhysicalDeviceShaderSubgroupRotateFeaturesKHR                StructureType = StructureTypePhysicalDeviceShaderSubgroupRotateFeatures
	StructureTypeDeviceQueueShaderCoreControlCreateInfoARM                    StructureType = 1000417000
	StructureTypePhysicalDeviceSchedulingControlsFeaturesARM                  StructureType = 1000417001
	StructureTypePhysicalDeviceSchedulingControlsPropertiesARM                StructureType = 1000417002
	StructureTypePhysicalDeviceImageSlicedViewOf3dFeaturesEXT                 StructureType = 1000418000
	StructureTypeImageViewSlicedCreateInfoEXT                                 StructureType = 1000418001
	StructureTypePhysicalDeviceDescriptorSetHostMappingFeaturesVALVE          StructureType = 1000420000
	StructureTypeDescriptorSetBindingReferenceVALVE                           StructureType = 1000420001
	StructureTypeDescriptorSetLayoutHostMappingInfoVALVE                      StructureType = 1000420002
	StructureTypePhysicalDeviceDepthClampZeroOneFeaturesEXT                   StructureType = StructureTypePhysicalDeviceDepthClampZeroOneFeaturesKHR
	StructureTypePhysicalDeviceNonSeamlessCubeMapFeaturesEXT                  StructureType = 1000422000
	StructureTypePhysicalDeviceRenderPassStripedFeaturesARM                   StructureType = 1000424000
	StructureTypePhysicalDeviceRenderPassStripedPropertiesARM                 StructureType = 1000424001
	StructureTypeRenderPassStripeBeginInfoARM                                 StructureType = 1000424002
	StructureTypeRenderPassStripeInfoARM                                      StructureType = 1000424003
	StructureTypeRenderPassStripeSubmitInfoARM                                StructureType = 1000424004
	StructureTypePhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM           StructureType = StructureTypePhysicalDeviceFragmentDensityMapOffsetFeaturesEXT
	StructureTypePhysicalDeviceFragmentDensityMapOffsetPropertiesQCOM         StructureType = StructureTypePhysicalDeviceFragmentDensityMapOffsetPropertiesEXT
	StructureTypeSubpassFragmentDensityMapOffsetEndInfoQCOM                   StructureType = StructureTypeRenderPassFragmentDensityMapOffsetEndInfoEXT
	StructureTypePhysicalDeviceCopyMemoryIndirectFeaturesNV                   StructureType = 1000426000
	StructureTypePhysicalDeviceCopyMemoryIndirectPropertiesNV                 StructureType = StructureTypePhysicalDeviceCopyMemoryIndirectPropertiesKHR
	StructureTypePhysicalDeviceMemoryDecompressionFeaturesNV                  StructureType = StructureTypePhysicalDeviceMemoryDecompressionFeaturesEXT
	StructureTypePhysicalDeviceMemoryDecompressionPropertiesNV                StructureType = StructureTypePhysicalDeviceMemoryDecompressionPropertiesEXT
	StructureTypePhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV       StructureType = 1000428000
	StructureTypeComputePipelineIndirectBufferInfoNV                          StructureType = 1000428001
	StructureTypePipelineIndirectDeviceAddressInfoNV                          StructureType = 1000428002
	StructureTypePhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV         StructureType = 1000429008
	StructureTypeAccelerationStructureGeometryLinearSweptSpheresDataNV        StructureType = 1000429009
	StructureTypeAccelerationStructureGeometrySpheresDataNV                   StructureType = 1000429010
	StructureTypePhysicalDeviceLinearColorAttachmentFeaturesNV                StructureType = 1000430000
	StructureTypePhysicalDeviceShaderMaximalReconvergenceFeaturesKHR          StructureType = 1000434000
	StructureTypeApplicationParametersEXT                                     StructureType = 1000435000
	StructureTypePhysicalDeviceImageCompressionControlSwapchainFeaturesEXT    StructureType = 1000437000
	StructureTypePhysicalDeviceImageProcessingFeaturesQCOM                    StructureType = 1000440000
	StructureTypePhysicalDeviceImageProcessingPropertiesQCOM                  StructureType = 1000440001
	StructureTypeImageViewSampleWeightCreateInfoQCOM                          StructureType = 1000440002
	StructureTypePhysicalDeviceNestedCommandBufferFeaturesEXT                 StructureType = 1000451000
	StructureTypePhysicalDeviceNestedCommandBufferPropertiesEXT               StructureType = 1000451001
	StructureTypeNativeBufferUsageOhos                                        StructureType = 1000452000
	StructureTypeNativeBufferPropertiesOhos                                   StructureType = 1000452001
	StructureTypeNativeBufferFormatPropertiesOhos                             StructureType = 1000452002
	StructureTypeImportNativeBufferInfoOhos                                   StructureType = 1000452003
	StructureTypeMemoryGetNativeBufferInfoOhos                                StructureType = 1000452004
	StructureTypeExternalFormatOhos                                           StructureType = 1000452005
	StructureTypeExternalMemoryAcquireUnmodifiedEXT                           StructureType = 1000453000
	StructureTypePhysicalDeviceExtendedDynamicState3FeaturesEXT               StructureType = 1000455000
	StructureTypePhysicalDeviceExtendedDynamicState3PropertiesEXT             StructureType = 1000455001
	StructureTypePhysicalDeviceSubpassMergeFeedbackFeaturesEXT                StructureType = 1000458000
	StructureTypeRenderPassCreationControlEXT                                 StructureType = 1000458001
	StructureTypeRenderPassCreationFeedbackCreateInfoEXT                      StructureType = 1000458002
	StructureTypeRenderPassSubpassFeedbackCreateInfoEXT                       StructureType = 1000458003
	StructureTypeDirectDriverLoadingInfoLUNARG                                StructureType = 1000459000
	StructureTypeDirectDriverLoadingListLUNARG                                StructureType = 1000459001
	StructureTypeTensorCreateInfoARM                                          StructureType = 1000460000
	StructureTypeTensorViewCreateInfoARM                                      StructureType = 1000460001
	StructureTypeBindTensorMemoryInfoARM                                      StructureType = 1000460002
	StructureTypeWriteDescriptorSetTensorARM                                  StructureType = 1000460003
	StructureTypePhysicalDeviceTensorPropertiesARM                            StructureType = 1000460004
	StructureTypeTensorFormatPropertiesARM                                    StructureType = 1000460005
	StructureTypeTensorDescriptionARM                                         StructureType = 1000460006
	StructureTypeTensorMemoryRequirementsInfoARM                              StructureType = 1000460007
	StructureTypeTensorMemoryBarrierARM                                       StructureType = 1000460008
	StructureTypePhysicalDeviceTensorFeaturesARM                              StructureType = 1000460009
	StructureTypeDeviceTensorMemoryRequirementsARM                            StructureType = 1000460010
	StructureTypeCopyTensorInfoARM                                            StructureType = 1000460011
	StructureTypeTensorCopyARM                                                StructureType = 1000460012
	StructureTypeTensorDependencyInfoARM                                      StructureType = 1000460013
	StructureTypeMemoryDedicatedAllocateInfoTensorARM                         StructureType = 1000460014
	StructureTypePhysicalDeviceExternalTensorInfoARM                          StructureType = 1000460015
	StructureTypeExternalTensorPropertiesARM                                  StructureType = 1000460016
	StructureTypeExternalMemoryTensorCreateInfoARM                            StructureType = 1000460017
	StructureTypePhysicalDeviceDescriptorBufferTensorFeaturesARM              StructureType = 1000460018
	StructureTypePhysicalDeviceDescriptorBufferTensorPropertiesARM            StructureType = 1000460019
	StructureTypeDescriptorGetTensorInfoARM                                   StructureType = 1000460020
	StructureTypeTensorCaptureDescriptorDataInfoARM                           StructureType = 1000460021
	StructureTypeTensorViewCaptureDescriptorDataInfoARM                       StructureType = 1000460022
	StructureTypeFrameBoundaryTensorsARM                                      StructureType = 1000460023
	StructureTypePhysicalDeviceShaderModuleIdentifierFeaturesEXT              StructureType = 1000462000
	StructureTypePhysicalDeviceShaderModuleIdentifierPropertiesEXT            StructureType = 1000462001
	StructureTypePipelineShaderStageModuleIdentifierCreateInfoEXT             StructureType = 1000462002
	StructureTypeShaderModuleIdentifierEXT                                    StructureType = 1000462003
	StructureTypePhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT  StructureType = 1000342000
	StructureTypePhysicalDeviceOpticalFlowFeaturesNV                          StructureType = 1000464000
	StructureTypePhysicalDeviceOpticalFlowPropertiesNV                        StructureType = 1000464001
	StructureTypeOpticalFlowImageFormatInfoNV                                 StructureType = 1000464002
	StructureTypeOpticalFlowImageFormatPropertiesNV                           StructureType = 1000464003
	StructureTypeOpticalFlowSessionCreateInfoNV                               StructureType = 1000464004
	StructureTypeOpticalFlowExecuteInfoNV                                     StructureType = 1000464005
	StructureTypeOpticalFlowSessionCreatePrivateDataInfoNV                    StructureType = 1000464010
	StructureTypePhysicalDeviceLegacyDitheringFeaturesEXT                     StructureType = 1000465000
	StructureTypePhysicalDevicePipelineProtectedAccessFeaturesEXT             StructureType = StructureTypePhysicalDevicePipelineProtectedAccessFeatures
	StructureTypePhysicalDeviceExternalFormatResolveFeaturesAndroid           StructureType = 1000468000
	StructureTypePhysicalDeviceExternalFormatResolvePropertiesAndroid         StructureType = 1000468001
	StructureTypeAndroidHardwareBufferFormatResolvePropertiesAndroid          StructureType = 1000468002
	StructureTypePhysicalDeviceMaintenance5FeaturesKHR                        StructureType = StructureTypePhysicalDeviceMaintenance5Features
	StructureTypePhysicalDeviceMaintenance5PropertiesKHR                      StructureType = StructureTypePhysicalDeviceMaintenance5Properties
	StructureTypeRenderingAreaInfoKHR                                         StructureType = StructureTypeRenderingAreaInfo
	StructureTypeDeviceImageSubresourceInfoKHR                                StructureType = StructureTypeDeviceImageSubresourceInfo
	StructureTypeSubresourceLayout2KHR                                        StructureType = StructureTypeSubresourceLayout2
	StructureTypeImageSubresource2KHR                                         StructureType = StructureTypeImageSubresource2
	StructureTypePipelineCreateFlags2CreateInfoKHR                            StructureType = StructureTypePipelineCreateFlags2CreateInfo
	StructureTypeBufferUsageFlags2CreateInfoKHR                               StructureType = StructureTypeBufferUsageFlags2CreateInfo
	StructureTypePhysicalDeviceAntiLagFeaturesAMD                             StructureType = 1000476000
	StructureTypeAntiLagDataAMD                                               StructureType = 1000476001
	StructureTypeAntiLagPresentationInfoAMD                                   StructureType = 1000476002
	StructureTypePhysicalDeviceDenseGeometryFormatFeaturesAmdx                StructureType = 1000478000
	StructureTypeAccelerationStructureDenseGeometryFormatTrianglesDataAmdx    StructureType = 1000478001
	StructureTypeSurfaceCapabilitiesPresentId2KHR                             StructureType = 1000479000
	StructureTypePresentId2KHR                                                StructureType = 1000479001
	StructureTypePhysicalDevicePresentId2FeaturesKHR                          StructureType = 1000479002
	StructureTypeSurfaceCapabilitiesPresentWait2KHR                           StructureType = 1000480000
	StructureTypePhysicalDevicePresentWait2FeaturesKHR                        StructureType = 1000480001
	StructureTypePresentWait2InfoKHR                                          StructureType = 1000480002
	StructureTypePhysicalDeviceRayTracingPositionFetchFeaturesKHR             StructureType = 1000481000
	StructureTypePhysicalDeviceShaderObjectFeaturesEXT                        StructureType = 1000482000
	StructureTypePhysicalDeviceShaderObjectPropertiesEXT                      StructureType = 1000482001
	StructureTypeShaderCreateInfoEXT                                          StructureType = 1000482002
	StructureTypeShaderRequiredSubgroupSizeCreateInfoEXT                      StructureType = StructureTypePipelineShaderStageRequiredSubgroupSizeCreateInfo
	StructureTypePhysicalDevicePipelineBinaryFeaturesKHR                      StructureType = 1000483000
	StructureTypePipelineBinaryCreateInfoKHR                                  StructureType = 1000483001
	StructureTypePipelineBinaryInfoKHR                                        StructureType = 1000483002
	StructureTypePipelineBinaryKeyKHR                                         StructureType = 1000483003
	StructureTypePhysicalDevicePipelineBinaryPropertiesKHR                    StructureType = 1000483004
	StructureTypeReleaseCapturedPipelineDataInfoKHR                           StructureType = 1000483005
	StructureTypePipelineBinaryDataInfoKHR                                    StructureType = 1000483006
	StructureTypePipelineCreateInfoKHR                                        StructureType = 1000483007
	StructureTypeDevicePipelineBinaryInternalCacheControlKHR                  StructureType = 1000483008
	StructureTypePipelineBinaryHandlesInfoKHR                                 StructureType = 1000483009
	StructureTypePhysicalDeviceTilePropertiesFeaturesQCOM                     StructureType = 1000484000
	StructureTypeTilePropertiesQCOM                                           StructureType = 1000484001
	StructureTypePhysicalDeviceAmigoProfilingFeaturesSEC                      StructureType = 1000485000
	StructureTypeAmigoProfilingSubmitInfoSEC                                  StructureType = 1000485001
	StructureTypeSurfacePresentModeKHR                                        StructureType = 1000274000
	StructureTypeSurfacePresentScalingCapabilitiesKHR                         StructureType = 1000274001
	StructureTypeSurfacePresentModeCompatibilityKHR                           StructureType = 1000274002
	StructureTypePhysicalDeviceSwapchainMaintenance1FeaturesKHR               StructureType = 1000275000
	StructureTypeSwapchainPresentFenceInfoKHR                                 StructureType = 1000275001
	StructureTypeSwapchainPresentModesCreateInfoKHR                           StructureType = 1000275002
	StructureTypeSwapchainPresentModeInfoKHR                                  StructureType = 1000275003
	StructureTypeSwapchainPresentScalingCreateInfoKHR                         StructureType = 1000275004
	StructureTypeReleaseSwapchainImagesInfoKHR                                StructureType = 1000275005
	StructureTypePhysicalDeviceMultiviewPerViewViewportsFeaturesQCOM          StructureType = 1000488000
	StructureTypeSemaphoreSciSyncPoolCreateInfoNV                             StructureType = 1000489000
	StructureTypeSemaphoreSciSyncCreateInfoNV                                 StructureType = 1000489001
	StructureTypePhysicalDeviceExternalSciSync2FeaturesNV                     StructureType = 1000489002
	StructureTypeDeviceSemaphoreSciSyncPoolReservationCreateInfoNV            StructureType = 1000489003
	StructureTypePhysicalDeviceRayTracingInvocationReorderFeaturesNV          StructureType = 1000490000
	StructureTypePhysicalDeviceRayTracingInvocationReorderPropertiesNV        StructureType = 1000490001
	StructureTypePhysicalDeviceCooperativeVectorFeaturesNV                    StructureType = 1000491000
	StructureTypePhysicalDeviceCooperativeVectorPropertiesNV                  StructureType = 1000491001
	StructureTypeCooperativeVectorPropertiesNV                                StructureType = 1000491002
	StructureTypeConvertCooperativeVectorMatrixInfoNV                         StructureType = 1000491004
	StructureTypePhysicalDeviceExtendedSparseAddressSpaceFeaturesNV           StructureType = 1000492000
	StructureTypePhysicalDeviceExtendedSparseAddressSpacePropertiesNV         StructureType = 1000492001
	StructureTypePhysicalDeviceMutableDescriptorTypeFeaturesEXT               StructureType = 1000351000
	StructureTypeMutableDescriptorTypeCreateInfoEXT                           StructureType = 1000351002
	StructureTypePhysicalDeviceLegacyVertexAttributesFeaturesEXT              StructureType = 1000495000
	StructureTypePhysicalDeviceLegacyVertexAttributesPropertiesEXT            StructureType = 1000495001
	StructureTypeLayerSettingsCreateInfoEXT                                   StructureType = 1000496000
	StructureTypePhysicalDeviceShaderCoreBuiltinsFeaturesARM                  StructureType = 1000497000
	StructureTypePhysicalDeviceShaderCoreBuiltinsPropertiesARM                StructureType = 1000497001
	StructureTypePhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT         StructureType = 1000498000
	StructureTypePhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT   StructureType = 1000499000
	StructureTypePhysicalDeviceInternallySynchronizedQueuesFeaturesKHR        StructureType = 1000504000
	StructureTypeLatencySleepModeInfoNV                                       StructureType = 1000505000
	StructureTypeLatencySleepInfoNV                                           StructureType = 1000505001
	StructureTypeSetLatencyMarkerInfoNV                                       StructureType = 1000505002
	StructureTypeGetLatencyMarkerInfoNV                                       StructureType = 1000505003
	StructureTypeLatencyTimingsFrameReportNV                                  StructureType = 1000505004
	StructureTypeLatencySubmissionPresentIdNV                                 StructureType = 1000505005
	StructureTypeOutOfBandQueueTypeInfoNV                                     StructureType = 1000505006
	StructureTypeSwapchainLatencyCreateInfoNV                                 StructureType = 1000505007
	StructureTypeLatencySurfaceCapabilitiesNV                                 StructureType = 1000505008
	StructureTypePhysicalDeviceCooperativeMatrixFeaturesKHR                   StructureType = 1000506000
	StructureTypeCooperativeMatrixPropertiesKHR                               StructureType = 1000506001
	StructureTypePhysicalDeviceCooperativeMatrixPropertiesKHR                 StructureType = 1000506002
	StructureTypeDataGraphPipelineCreateInfoARM                               StructureType = 1000507000
	StructureTypeDataGraphPipelineSessionCreateInfoARM                        StructureType = 1000507001
	StructureTypeDataGraphPipelineResourceInfoARM                             StructureType = 1000507002
	StructureTypeDataGraphPipelineConstantARM                                 StructureType = 1000507003
	StructureTypeDataGraphPipelineSessionMemoryRequirementsInfoARM            StructureType = 1000507004
	StructureTypeBindDataGraphPipelineSessionMemoryInfoARM                    StructureType = 1000507005
	StructureTypePhysicalDeviceDataGraphFeaturesARM                           StructureType = 1000507006
	StructureTypeDataGraphPipelineShaderModuleCreateInfoARM                   StructureType = 1000507007
	StructureTypeDataGraphPipelinePropertyQueryResultARM                      StructureType = 1000507008
	StructureTypeDataGraphPipelineInfoARM                                     StructureType = 1000507009
	StructureTypeDataGraphPipelineCompilerControlCreateInfoARM                StructureType = 1000507010
	StructureTypeDataGraphPipelineSessionBindPointRequirementsInfoARM         StructureType = 1000507011
	StructureTypeDataGraphPipelineSessionBindPointRequirementARM              StructureType = 1000507012
	StructureTypeDataGraphPipelineIdentifierCreateInfoARM                     StructureType = 1000507013
	StructureTypeDataGraphPipelineDispatchInfoARM                             StructureType = 1000507014
	StructureTypeDataGraphProcessingEngineCreateInfoARM                       StructureType = 1000507016
	StructureTypeQueueFamilyDataGraphProcessingEnginePropertiesARM            StructureType = 1000507017
	StructureTypeQueueFamilyDataGraphPropertiesARM                            StructureType = 1000507018
	StructureTypePhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM    StructureType = 1000507019
	StructureTypeDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM StructureType = 1000507015
	StructureTypePhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM        StructureType = 1000510000
	StructureTypeMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM           StructureType = 1000510001
	StructureTypePhysicalDeviceComputeShaderDerivativesFeaturesKHR            StructureType = 1000201000
	StructureTypePhysicalDeviceComputeShaderDerivativesPropertiesKHR          StructureType = 1000511000
	StructureTypeVideoDecodeAv1CapabilitiesKHR                                StructureType = 1000512000
	StructureTypeVideoDecodeAv1PictureInfoKHR                                 StructureType = 1000512001
	StructureTypeVideoDecodeAv1ProfileInfoKHR                                 StructureType = 1000512003
	StructureTypeVideoDecodeAv1SessionParametersCreateInfoKHR                 StructureType = 1000512004
	StructureTypeVideoDecodeAv1DpbSlotInfoKHR                                 StructureType = 1000512005
	StructureTypeVideoEncodeAv1CapabilitiesKHR                                StructureType = 1000513000
	StructureTypeVideoEncodeAv1SessionParametersCreateInfoKHR                 StructureType = 1000513001
	StructureTypeVideoEncodeAv1PictureInfoKHR                                 StructureType = 1000513002
	StructureTypeVideoEncodeAv1DpbSlotInfoKHR                                 StructureType = 1000513003
	StructureTypePhysicalDeviceVideoEncodeAv1FeaturesKHR                      StructureType = 1000513004
	StructureTypeVideoEncodeAv1ProfileInfoKHR                                 StructureType = 1000513005
	StructureTypeVideoEncodeAv1RateControlInfoKHR                             StructureType = 1000513006
	StructureTypeVideoEncodeAv1RateControlLayerInfoKHR                        StructureType = 1000513007
	StructureTypeVideoEncodeAv1QualityLevelPropertiesKHR                      StructureType = 1000513008
	StructureTypeVideoEncodeAv1SessionCreateInfoKHR                           StructureType = 1000513009
	StructureTypeVideoEncodeAv1GopRemainingFrameInfoKHR                       StructureType = 1000513010
	StructureTypePhysicalDeviceVideoDecodeVp9FeaturesKHR                      StructureType = 1000514000
	StructureTypeVideoDecodeVp9CapabilitiesKHR                                StructureType = 1000514001
	StructureTypeVideoDecodeVp9PictureInfoKHR                                 StructureType = 1000514002
	StructureTypeVideoDecodeVp9ProfileInfoKHR                                 StructureType = 1000514003
	StructureTypePhysicalDeviceVideoMaintenance1FeaturesKHR                   StructureType = 1000515000
	StructureTypeVideoInlineQueryInfoKHR                                      StructureType = 1000515001
	StructureTypePhysicalDevicePerStageDescriptorSetFeaturesNV                StructureType = 1000516000
	StructureTypePhysicalDeviceImageProcessing2FeaturesQCOM                   StructureType = 1000518000
	StructureTypePhysicalDeviceImageProcessing2PropertiesQCOM                 StructureType = 1000518001
	StructureTypeSamplerBlockMatchWindowCreateInfoQCOM                        StructureType = 1000518002
	StructureTypeSamplerCubicWeightsCreateInfoQCOM                            StructureType = 1000519000
	StructureTypePhysicalDeviceCubicWeightsFeaturesQCOM                       StructureType = 1000519001
	StructureTypeBlitImageCubicWeightsInfoQCOM                                StructureType = 1000519002
	StructureTypePhysicalDeviceYcbcrDegammaFeaturesQCOM                       StructureType = 1000520000
	StructureTypeSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM             StructureType = 1000520001
	StructureTypePhysicalDeviceCubicClampFeaturesQCOM                         StructureType = 1000521000
	StructureTypePhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT  StructureType = 1000524000
	StructureTypePhysicalDeviceVertexAttributeDivisorPropertiesKHR            StructureType = StructureTypePhysicalDeviceVertexAttributeDivisorProperties
	StructureTypePipelineVertexInputDivisorStateCreateInfoKHR                 StructureType = StructureTypePipelineVertexInputDivisorStateCreateInfo
	StructureTypePhysicalDeviceVertexAttributeDivisorFeaturesKHR              StructureType = StructureTypePhysicalDeviceVertexAttributeDivisorFeatures
	StructureTypePhysicalDeviceUnifiedImageLayoutsFeaturesKHR                 StructureType = 1000527000
	StructureTypeAttachmentFeedbackLoopInfoEXT                                StructureType = 1000527001
	StructureTypePhysicalDeviceShaderFloatControls2FeaturesKHR                StructureType = StructureTypePhysicalDeviceShaderFloatControls2Features
	StructureTypeScreenBufferPropertiesQnx                                    StructureType = 1000529000
	StructureTypeScreenBufferFormatPropertiesQnx                              StructureType = 1000529001
	StructureTypeImportScreenBufferInfoQnx                                    StructureType = 1000529002
	StructureTypeExternalFormatQnx                                            StructureType = 1000529003
	StructureTypePhysicalDeviceExternalMemoryScreenBufferFeaturesQnx          StructureType = 1000529004
	StructureTypePhysicalDeviceLayeredDriverPropertiesMsft                    StructureType = 1000530000
	StructureTypePhysicalDeviceIndexTypeUint8FeaturesKHR                      StructureType = StructureTypePhysicalDeviceIndexTypeUint8Features
	StructureTypePhysicalDeviceLineRasterizationFeaturesKHR                   StructureType = StructureTypePhysicalDeviceLineRasterizationFeatures
	StructureTypePipelineRasterizationLineStateCreateInfoKHR                  StructureType = StructureTypePipelineRasterizationLineStateCreateInfo
	StructureTypePhysicalDeviceLineRasterizationPropertiesKHR                 StructureType = StructureTypePhysicalDeviceLineRasterizationProperties
	StructureTypeCalibratedTimestampInfoKHR                                   StructureType = 1000184000
	StructureTypePhysicalDeviceShaderExpectAssumeFeaturesKHR                  StructureType = StructureTypePhysicalDeviceShaderExpectAssumeFeatures
	StructureTypePhysicalDeviceMaintenance6FeaturesKHR                        StructureType = StructureTypePhysicalDeviceMaintenance6Features
	StructureTypePhysicalDeviceMaintenance6PropertiesKHR                      StructureType = StructureTypePhysicalDeviceMaintenance6Properties
	StructureTypeBindMemoryStatusKHR                                          StructureType = StructureTypeBindMemoryStatus
	StructureTypeBindDescriptorSetsInfoKHR                                    StructureType = StructureTypeBindDescriptorSetsInfo
	StructureTypePushConstantsInfoKHR                                         StructureType = StructureTypePushConstantsInfo
	StructureTypePushDescriptorSetInfoKHR                                     StructureType = StructureTypePushDescriptorSetInfo
	StructureTypePushDescriptorSetWithTemplateInfoKHR                         StructureType = StructureTypePushDescriptorSetWithTemplateInfo
	StructureTypeSetDescriptorBufferOffsetsInfoEXT                            StructureType = 1000545007
	StructureTypeBindDescriptorBufferEmbeddedSamplersInfoEXT                  StructureType = 1000545008
	StructureTypePhysicalDeviceDescriptorPoolOverallocationFeaturesNV         StructureType = 1000546000
	StructureTypePhysicalDeviceTileMemoryHeapFeaturesQCOM                     StructureType = 1000547000
	StructureTypePhysicalDeviceTileMemoryHeapPropertiesQCOM                   StructureType = 1000547001
	StructureTypeTileMemoryRequirementsQCOM                                   StructureType = 1000547002
	StructureTypeTileMemoryBindInfoQCOM                                       StructureType = 1000547003
	StructureTypeTileMemorySizeInfoQCOM                                       StructureType = 1000547004
	StructureTypePhysicalDeviceCopyMemoryIndirectFeaturesKHR                  StructureType = 1000549000
	StructureTypePhysicalDeviceCopyMemoryIndirectPropertiesKHR                StructureType = 1000426001
	StructureTypeCopyMemoryIndirectInfoKHR                                    StructureType = 1000549002
	StructureTypeCopyMemoryToImageIndirectInfoKHR                             StructureType = 1000549003
	StructureTypePhysicalDeviceMemoryDecompressionFeaturesEXT                 StructureType = 1000427000
	StructureTypePhysicalDeviceMemoryDecompressionPropertiesEXT               StructureType = 1000427001
	StructureTypeDecompressMemoryInfoEXT                                      StructureType = 1000550002
	StructureTypeDisplaySurfaceStereoCreateInfoNV                             StructureType = 1000551000
	StructureTypeDisplayModeStereoPropertiesNV                                StructureType = 1000551001
	StructureTypeVideoEncodeIntraRefreshCapabilitiesKHR                       StructureType = 1000552000
	StructureTypeVideoEncodeSessionIntraRefreshCreateInfoKHR                  StructureType = 1000552001
	StructureTypeVideoEncodeIntraRefreshInfoKHR                               StructureType = 1000552002
	StructureTypeVideoReferenceIntraRefreshInfoKHR                            StructureType = 1000552003
	StructureTypePhysicalDeviceVideoEncodeIntraRefreshFeaturesKHR             StructureType = 1000552004
	StructureTypeVideoEncodeQuantizationMapCapabilitiesKHR                    StructureType = 1000553000
	StructureTypeVideoFormatQuantizationMapPropertiesKHR                      StructureType = 1000553001
	StructureTypeVideoEncodeQuantizationMapInfoKHR                            StructureType = 1000553002
	StructureTypeVideoEncodeQuantizationMapSessionParametersCreateInfoKHR     StructureType = 1000553005
	StructureTypePhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR          StructureType = 1000553009
	StructureTypeVideoEncodeH264QuantizationMapCapabilitiesKHR                StructureType = 1000553003
	StructureTypeVideoEncodeH265QuantizationMapCapabilitiesKHR                StructureType = 1000553004
	StructureTypeVideoFormatH265QuantizationMapPropertiesKHR                  StructureType = 1000553006
	StructureTypeVideoEncodeAv1QuantizationMapCapabilitiesKHR                 StructureType = 1000553007
	StructureTypeVideoFormatAv1QuantizationMapPropertiesKHR                   StructureType = 1000553008
	StructureTypePhysicalDeviceRawAccessChainsFeaturesNV                      StructureType = 1000555000
	StructureTypeExternalComputeQueueDeviceCreateInfoNV                       StructureType = 1000556000
	StructureTypeExternalComputeQueueCreateInfoNV                             StructureType = 1000556001
	StructureTypeExternalComputeQueueDataParamsNV                             StructureType = 1000556002
	StructureTypePhysicalDeviceExternalComputeQueuePropertiesNV               StructureType = 1000556003
	StructureTypePhysicalDeviceShaderRelaxedExtendedInstructionFeaturesKHR    StructureType = 1000558000
	StructureTypePhysicalDeviceCommandBufferInheritanceFeaturesNV             StructureType = 1000559000
	StructureTypePhysicalDeviceMaintenance7FeaturesKHR                        StructureType = 1000562000
	StructureTypePhysicalDeviceMaintenance7PropertiesKHR                      StructureType = 1000562001
	StructureTypePhysicalDeviceLayeredApiPropertiesListKHR                    StructureType = 1000562002
	StructureTypePhysicalDeviceLayeredApiPropertiesKHR                        StructureType = 1000562003
	StructureTypePhysicalDeviceLayeredApiVulkanPropertiesKHR                  StructureType = 1000562004
	StructureTypePhysicalDeviceShaderAtomicFloat16VectorFeaturesNV            StructureType = 1000563000
	StructureTypePhysicalDeviceShaderReplicatedCompositesFeaturesEXT          StructureType = 1000564000
	StructureTypePhysicalDeviceShaderFloat8FeaturesEXT                        StructureType = 1000567000
	StructureTypePhysicalDeviceRayTracingValidationFeaturesNV                 StructureType = 1000568000
	StructureTypePhysicalDeviceClusterAccelerationStructureFeaturesNV         StructureType = 1000569000
	StructureTypePhysicalDeviceClusterAccelerationStructurePropertiesNV       StructureType = 1000569001
	StructureTypeClusterAccelerationStructureClustersBottomLevelInputNV       StructureType = 1000569002
	StructureTypeClusterAccelerationStructureTriangleClusterInputNV           StructureType = 1000569003
	StructureTypeClusterAccelerationStructureMoveObjectsInputNV               StructureType = 1000569004
	StructureTypeClusterAccelerationStructureInputInfoNV                      StructureType = 1000569005
	StructureTypeClusterAccelerationStructureCommandsInfoNV                   StructureType = 1000569006
	StructureTypeRayTracingPipelineClusterAccelerationStructureCreateInfoNV   StructureType = 1000569007
	StructureTypePhysicalDevicePartitionedAccelerationStructureFeaturesNV     StructureType = 1000570000
	StructureTypePhysicalDevicePartitionedAccelerationStructurePropertiesNV   StructureType = 1000570001
	StructureTypeWriteDescriptorSetPartitionedAccelerationStructureNV         StructureType = 1000570002
	StructureTypePartitionedAccelerationStructureInstancesInputNV             StructureType = 1000570003
	StructureTypeBuildPartitionedAccelerationStructureInfoNV                  StructureType = 1000570004
	StructureTypePartitionedAccelerationStructureFlagsNV                      StructureType = 1000570005
	StructureTypePhysicalDeviceDeviceGeneratedCommandsFeaturesEXT             StructureType = 1000572000
	StructureTypePhysicalDeviceDeviceGeneratedCommandsPropertiesEXT           StructureType = 1000572001
	StructureTypeGeneratedCommandsMemoryRequirementsInfoEXT                   StructureType = 1000572002
	StructureTypeIndirectExecutionSetCreateInfoEXT                            StructureType = 1000572003
	StructureTypeGeneratedCommandsInfoEXT                                     StructureType = 1000572004
	StructureTypeIndirectCommandsLayoutCreateInfoEXT                          StructureType = 1000572006
	StructureTypeIndirectCommandsLayoutTokenEXT                               StructureType = 1000572007
	StructureTypeWriteIndirectExecutionSetPipelineEXT                         StructureType = 1000572008
	StructureTypeWriteIndirectExecutionSetShaderEXT                           StructureType = 1000572009
	StructureTypeIndirectExecutionSetPipelineInfoEXT                          StructureType = 1000572010
	StructureTypeIndirectExecutionSetShaderInfoEXT                            StructureType = 1000572011
	StructureTypeIndirectExecutionSetShaderLayoutInfoEXT                      StructureType = 1000572012
	StructureTypeGeneratedCommandsPipelineInfoEXT                             StructureType = 1000572013
	StructureTypeGeneratedCommandsShaderInfoEXT                               StructureType = 1000572014
	StructureTypePhysicalDeviceMaintenance8FeaturesKHR                        StructureType = 1000574000
	StructureTypeMemoryBarrierAccessFlags3KHR                                 StructureType = 1000574002
	StructureTypePhysicalDeviceImageAlignmentControlFeaturesMESA              StructureType = 1000575000
	StructureTypePhysicalDeviceImageAlignmentControlPropertiesMESA            StructureType = 1000575001
	StructureTypeImageAlignmentControlCreateInfoMESA                          StructureType = 1000575002
	StructureTypePhysicalDeviceShaderFmaFeaturesKHR                           StructureType = 1000579000
	StructureTypePushConstantBankInfoNV                                       StructureType = 1000580000
	StructureTypePhysicalDevicePushConstantBankFeaturesNV                     StructureType = 1000580001
	StructureTypePhysicalDevicePushConstantBankPropertiesNV                   StructureType = 1000580002
	StructureTypePhysicalDeviceRayTracingInvocationReorderFeaturesEXT         StructureType = 1000581000
	StructureTypePhysicalDeviceRayTracingInvocationReorderPropertiesEXT       StructureType = 1000581001
	StructureTypePhysicalDeviceDepthClampControlFeaturesEXT                   StructureType = 1000582000
	StructureTypePipelineViewportDepthClampControlCreateInfoEXT               StructureType = 1000582001
	StructureTypePhysicalDeviceMaintenance9FeaturesKHR                        StructureType = 1000584000
	StructureTypePhysicalDeviceMaintenance9PropertiesKHR                      StructureType = 1000584001
	StructureTypeQueueFamilyOwnershipTransferPropertiesKHR                    StructureType = 1000584002
	StructureTypePhysicalDeviceVideoMaintenance2FeaturesKHR                   StructureType = 1000586000
	StructureTypeVideoDecodeH264InlineSessionParametersInfoKHR                StructureType = 1000586001
	StructureTypeVideoDecodeH265InlineSessionParametersInfoKHR                StructureType = 1000586002
	StructureTypeVideoDecodeAv1InlineSessionParametersInfoKHR                 StructureType = 1000586003
	StructureTypeSurfaceCreateInfoOhos                                        StructureType = 1000685000
	StructureTypeNativeBufferOhos                                             StructureType = 1000453001
	StructureTypeSwapchainImageCreateInfoOhos                                 StructureType = 1000453002
	StructureTypePhysicalDevicePresentationPropertiesOhos                     StructureType = 1000453003
	StructureTypePhysicalDeviceHdrVividFeaturesHUAWEI                         StructureType = 1000591000
	StructureTypeHdrVividDynamicMetadataHUAWEI                                StructureType = 1000591001
	StructureTypePhysicalDeviceCooperativeMatrix2FeaturesNV                   StructureType = 1000594000
	StructureTypeCooperativeMatrixFlexibleDimensionsPropertiesNV              StructureType = 1000594001
	StructureTypePhysicalDeviceCooperativeMatrix2PropertiesNV                 StructureType = 1000594002
	StructureTypePhysicalDevicePipelineOpacityMicromapFeaturesARM             StructureType = 1000597000
	StructureTypeImportMemoryMetalHandleInfoEXT                               StructureType = 1000603000
	StructureTypeMemoryMetalHandlePropertiesEXT                               StructureType = 1000603001
	StructureTypeMemoryGetMetalHandleInfoEXT                                  StructureType = 1000603002
	StructureTypePhysicalDeviceDepthClampZeroOneFeaturesKHR                   StructureType = 1000421000
	StructureTypePhysicalDevicePerformanceCountersByRegionFeaturesARM         StructureType = 1000606000
	StructureTypePhysicalDevicePerformanceCountersByRegionPropertiesARM       StructureType = 1000606001
	StructureTypePerformanceCounterARM                                        StructureType = 1000606002
	StructureTypePerformanceCounterDescriptionARM                             StructureType = 1000606003
	StructureTypeRenderPassPerformanceCountersByRegionBeginInfoARM            StructureType = 1000606004
	StructureTypePhysicalDeviceVertexAttributeRobustnessFeaturesEXT           StructureType = 1000609000
	StructureTypePhysicalDeviceFormatPackFeaturesARM                          StructureType = 1000610000
	StructureTypePhysicalDeviceFragmentDensityMapLayeredFeaturesVALVE         StructureType = 1000612000
	StructureTypePhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE       StructureType = 1000612001
	StructureTypePipelineFragmentDensityMapLayeredCreateInfoVALVE             StructureType = 1000612002
	StructureTypePhysicalDeviceRobustness2FeaturesKHR                         StructureType = 1000286000
	StructureTypePhysicalDeviceRobustness2PropertiesKHR                       StructureType = 1000286001
	StructureTypeSetPresentConfigNV                                           StructureType = 1000614000
	StructureTypePhysicalDevicePresentMeteringFeaturesNV                      StructureType = 1000614001
	StructureTypePhysicalDeviceFragmentDensityMapOffsetFeaturesEXT            StructureType = 1000425000
	StructureTypePhysicalDeviceFragmentDensityMapOffsetPropertiesEXT          StructureType = 1000425001
	StructureTypeRenderPassFragmentDensityMapOffsetEndInfoEXT                 StructureType = 1000425002
	StructureTypeRenderingEndInfoEXT                                          StructureType = StructureTypeRenderingEndInfoKHR
	StructureTypePhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT          StructureType = 1000620000
	StructureTypePhysicalDevicePresentModeFifoLatestReadyFeaturesKHR          StructureType = 1000361000
	StructureTypePhysicalDeviceShader64BitIndexingFeaturesEXT                 StructureType = 1000627000
	StructureTypePhysicalDeviceCustomResolveFeaturesEXT                       StructureType = 1000628000
	StructureTypeBeginCustomResolveInfoEXT                                    StructureType = 1000628001
	StructureTypeCustomResolveCreateInfoEXT                                   StructureType = 1000628002
	StructureTypePhysicalDeviceDataGraphModelFeaturesQCOM                     StructureType = 1000629000
	StructureTypeDataGraphPipelineBuiltinModelCreateInfoQCOM                  StructureType = 1000629001
	StructureTypePhysicalDeviceMaintenance10FeaturesKHR                       StructureType = 1000630000
	StructureTypePhysicalDeviceMaintenance10PropertiesKHR                     StructureType = 1000630001
	StructureTypeRenderingAttachmentFlagsInfoKHR                              StructureType = 1000630002
	StructureTypeRenderingEndInfoKHR                                          StructureType = 1000619003
	StructureTypeResolveImageModeInfoKHR                                      StructureType = 1000630004
	StructureTypePhysicalDeviceShaderLongVectorFeaturesEXT                    StructureType = 1000635000
	StructureTypePhysicalDeviceShaderLongVectorPropertiesEXT                  StructureType = 1000635001
	StructureTypePhysicalDevicePipelineCacheIncrementalModeFeaturesSEC        StructureType = 1000637000
	StructureTypePhysicalDeviceShaderUniformBufferUnsizedArrayFeaturesEXT     StructureType = 1000642000
	StructureTypeComputeOccupancyPriorityParametersNV                         StructureType = 1000645000
	StructureTypePhysicalDeviceComputeOccupancyPriorityFeaturesNV             StructureType = 1000645001
	StructureTypePhysicalDeviceShaderSubgroupPartitionedFeaturesEXT           StructureType = 1000662000
	StructureTypeUbmSurfaceCreateInfoSEC                                      StructureType = 1000664000
	StructureTypePhysicalDeviceShaderMixedFloatDotProductFeaturesVALVE        StructureType = 1000673000
)

type SubpassContents uint32

const (
	SubpassContentsInline                              SubpassContents = 0
	SubpassContentsSecondaryCommandBuffers             SubpassContents = 1
	SubpassContentsInlineAndSecondaryCommandBuffersEXT SubpassContents = SubpassContentsInlineAndSecondaryCommandBuffersKHR
	SubpassContentsInlineAndSecondaryCommandBuffersKHR SubpassContents = 1000451000
)

type SubpassDescriptionFlagBits uint32

const (
	SubpassDescriptionPerViewAttributesBitNVX                         SubpassDescriptionFlagBits = 1 << 0
	SubpassDescriptionPerViewPositionXOnlyBitNVX                      SubpassDescriptionFlagBits = 1 << 1
	SubpassDescriptionFragmentRegionBitQCOM                           SubpassDescriptionFlagBits = SubpassDescriptionFragmentRegionBitEXT
	SubpassDescriptionShaderResolveBitQCOM                            SubpassDescriptionFlagBits = SubpassDescriptionCustomResolveBitEXT
	SubpassDescriptionTileShadingApronBitQCOM                         SubpassDescriptionFlagBits = 1 << 8
	SubpassDescriptionRasterizationOrderAttachmentColorAccessBitARM   SubpassDescriptionFlagBits = SubpassDescriptionRasterizationOrderAttachmentColorAccessBitEXT
	SubpassDescriptionRasterizationOrderAttachmentDepthAccessBitARM   SubpassDescriptionFlagBits = SubpassDescriptionRasterizationOrderAttachmentDepthAccessBitEXT
	SubpassDescriptionRasterizationOrderAttachmentStencilAccessBitARM SubpassDescriptionFlagBits = SubpassDescriptionRasterizationOrderAttachmentStencilAccessBitEXT
	SubpassDescriptionRasterizationOrderAttachmentColorAccessBitEXT   SubpassDescriptionFlagBits = 1 << 4
	SubpassDescriptionRasterizationOrderAttachmentDepthAccessBitEXT   SubpassDescriptionFlagBits = 1 << 5
	SubpassDescriptionRasterizationOrderAttachmentStencilAccessBitEXT SubpassDescriptionFlagBits = 1 << 6
	SubpassDescriptionEnableLegacyDitheringBitEXT                     SubpassDescriptionFlagBits = 1 << 7
	SubpassDescriptionFragmentRegionBitEXT                            SubpassDescriptionFlagBits = 1 << 2
	SubpassDescriptionCustomResolveBitEXT                             SubpassDescriptionFlagBits = 1 << 3
)

type SystemAllocationScope uint32

const (
	SystemAllocationScopeCommand  SystemAllocationScope = 0
	SystemAllocationScopeObject   SystemAllocationScope = 1
	SystemAllocationScopeCache    SystemAllocationScope = 2
	SystemAllocationScopeDevice   SystemAllocationScope = 3
	SystemAllocationScopeInstance SystemAllocationScope = 4
)

type TessellationDomainOrigin uint32

const (
	TessellationDomainOriginUpperLeft    TessellationDomainOrigin = 0
	TessellationDomainOriginLowerLeft    TessellationDomainOrigin = 1
	TessellationDomainOriginUpperLeftKHR TessellationDomainOrigin = TessellationDomainOriginUpperLeft
	TessellationDomainOriginLowerLeftKHR TessellationDomainOrigin = TessellationDomainOriginLowerLeft
)

type VertexInputRate uint32

const (
	VertexInputRateVertex   VertexInputRate = 0
	VertexInputRateInstance VertexInputRate = 1
)

type AccessFlags uint32

type AttachmentDescriptionFlags uint32

type ColorComponentFlags uint32

type CullModeFlags uint32

type DependencyFlags uint32

type FramebufferCreateFlags uint32

type ImageAspectFlags uint32

type ImageCreateFlags uint32

type ImageUsageFlags uint32

type PipelineColorBlendStateCreateFlags uint32

type PipelineCreateFlags uint32

type PipelineDepthStencilStateCreateFlags uint32

type PipelineDynamicStateCreateFlags uint32

type PipelineInputAssemblyStateCreateFlags uint32

type PipelineMultisampleStateCreateFlags uint32

type PipelineRasterizationStateCreateFlags uint32

type PipelineShaderStageCreateFlags uint32

type PipelineStageFlags uint32

type PipelineTessellationStateCreateFlags uint32

type PipelineVertexInputStateCreateFlags uint32

type PipelineViewportStateCreateFlags uint32

type RenderPassCreateFlags uint32

type RenderingFlags uint32

type ResolveModeFlags uint32

type StencilFaceFlags uint32

type SubpassDescriptionFlags uint32

type AllocationCallbacks struct {
	UserData              unsafe.Pointer
	PfnAllocation         AllocationFunction
	PfnReallocation       ReallocationFunction
	PfnFree               FreeFunction
	PfnInternalAllocation InternalAllocationNotification
	PfnInternalFree       InternalFreeNotification
}

func (s *AllocationCallbacks) toC() (*C.VkAllocationCallbacks, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkAllocationCallbacks)(C.malloc(C.size_t(C.sizeof_VkAllocationCallbacks)))
	*p = C.VkAllocationCallbacks{}
	p.pUserData = s.UserData
	var fn0 C.PFN_vkAllocationFunction
	if s.PfnAllocation != nil {
		fn0 = *(*C.PFN_vkAllocationFunction)(unsafe.Pointer(&s.PfnAllocation))
	}
	p.pfnAllocation = fn0
	var fn1 C.PFN_vkReallocationFunction
	if s.PfnReallocation != nil {
		fn1 = *(*C.PFN_vkReallocationFunction)(unsafe.Pointer(&s.PfnReallocation))
	}
	p.pfnReallocation = fn1
	var fn2 C.PFN_vkFreeFunction
	if s.PfnFree != nil {
		fn2 = *(*C.PFN_vkFreeFunction)(unsafe.Pointer(&s.PfnFree))
	}
	p.pfnFree = fn2
	var fn3 C.PFN_vkInternalAllocationNotification
	if s.PfnInternalAllocation != nil {
		fn3 = *(*C.PFN_vkInternalAllocationNotification)(unsafe.Pointer(&s.PfnInternalAllocation))
	}
	p.pfnInternalAllocation = fn3
	var fn4 C.PFN_vkInternalFreeNotification
	if s.PfnInternalFree != nil {
		fn4 = *(*C.PFN_vkInternalFreeNotification)(unsafe.Pointer(&s.PfnInternalFree))
	}
	p.pfnInternalFree = fn4
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type AttachmentDescription struct {
	Flags          AttachmentDescriptionFlags
	Format         Format
	Samples        SampleCountFlagBits
	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp
	InitialLayout  ImageLayout
	FinalLayout    ImageLayout
}

func (s *AttachmentDescription) toC() (*C.VkAttachmentDescription, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkAttachmentDescription)(C.malloc(C.size_t(C.sizeof_VkAttachmentDescription)))
	*p = C.VkAttachmentDescription{}
	val0 := C.VkAttachmentDescriptionFlags(s.Flags)
	p.flags = val0
	val1 := C.VkFormat(s.Format)
	p.format = val1
	val2 := C.VkSampleCountFlagBits(s.Samples)
	p.samples = val2
	val3 := C.VkAttachmentLoadOp(s.LoadOp)
	p.loadOp = val3
	val4 := C.VkAttachmentStoreOp(s.StoreOp)
	p.storeOp = val4
	val5 := C.VkAttachmentLoadOp(s.StencilLoadOp)
	p.stencilLoadOp = val5
	val6 := C.VkAttachmentStoreOp(s.StencilStoreOp)
	p.stencilStoreOp = val6
	val7 := C.VkImageLayout(s.InitialLayout)
	p.initialLayout = val7
	val8 := C.VkImageLayout(s.FinalLayout)
	p.finalLayout = val8
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type AttachmentDescription2 struct {
	Next           Structure
	Flags          AttachmentDescriptionFlags
	Format         Format
	Samples        SampleCountFlagBits
	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp
	InitialLayout  ImageLayout
	FinalLayout    ImageLayout
}

func (s *AttachmentDescription2) GetType() StructureType {
	return StructureTypeAttachmentDescription2
}

func (s *AttachmentDescription2) toC() (*C.VkAttachmentDescription2, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkAttachmentDescription2)(C.malloc(C.size_t(C.sizeof_VkAttachmentDescription2)))
	*p = C.VkAttachmentDescription2{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkAttachmentDescriptionFlags(s.Flags)
	p.flags = val0
	val1 := C.VkFormat(s.Format)
	p.format = val1
	val2 := C.VkSampleCountFlagBits(s.Samples)
	p.samples = val2
	val3 := C.VkAttachmentLoadOp(s.LoadOp)
	p.loadOp = val3
	val4 := C.VkAttachmentStoreOp(s.StoreOp)
	p.storeOp = val4
	val5 := C.VkAttachmentLoadOp(s.StencilLoadOp)
	p.stencilLoadOp = val5
	val6 := C.VkAttachmentStoreOp(s.StencilStoreOp)
	p.stencilStoreOp = val6
	val7 := C.VkImageLayout(s.InitialLayout)
	p.initialLayout = val7
	val8 := C.VkImageLayout(s.FinalLayout)
	p.finalLayout = val8
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type AttachmentDescriptionStencilLayout struct {
	Next                 Structure
	StencilInitialLayout ImageLayout
	StencilFinalLayout   ImageLayout
}

func (s *AttachmentDescriptionStencilLayout) GetType() StructureType {
	return StructureTypeAttachmentDescriptionStencilLayout
}

func (s *AttachmentDescriptionStencilLayout) toC() (*C.VkAttachmentDescriptionStencilLayout, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkAttachmentDescriptionStencilLayout)(C.malloc(C.size_t(C.sizeof_VkAttachmentDescriptionStencilLayout)))
	*p = C.VkAttachmentDescriptionStencilLayout{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkImageLayout(s.StencilInitialLayout)
	p.stencilInitialLayout = val0
	val1 := C.VkImageLayout(s.StencilFinalLayout)
	p.stencilFinalLayout = val1
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type AttachmentReference struct {
	Attachment uint32
	Layout     ImageLayout
}

func (s *AttachmentReference) toC() (*C.VkAttachmentReference, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkAttachmentReference)(C.malloc(C.size_t(C.sizeof_VkAttachmentReference)))
	*p = C.VkAttachmentReference{}
	val0 := C.uint32_t(s.Attachment)
	p.attachment = val0
	val1 := C.VkImageLayout(s.Layout)
	p.layout = val1
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type AttachmentReference2 struct {
	Next       Structure
	Attachment uint32
	Layout     ImageLayout
	AspectMask ImageAspectFlags
}

func (s *AttachmentReference2) GetType() StructureType {
	return StructureTypeAttachmentReference2
}

func (s *AttachmentReference2) toC() (*C.VkAttachmentReference2, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkAttachmentReference2)(C.malloc(C.size_t(C.sizeof_VkAttachmentReference2)))
	*p = C.VkAttachmentReference2{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.uint32_t(s.Attachment)
	p.attachment = val0
	val1 := C.VkImageLayout(s.Layout)
	p.layout = val1
	val2 := C.VkImageAspectFlags(s.AspectMask)
	p.aspectMask = val2
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type AttachmentReferenceStencilLayout struct {
	Next          Structure
	StencilLayout ImageLayout
}

func (s *AttachmentReferenceStencilLayout) GetType() StructureType {
	return StructureTypeAttachmentReferenceStencilLayout
}

func (s *AttachmentReferenceStencilLayout) toC() (*C.VkAttachmentReferenceStencilLayout, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkAttachmentReferenceStencilLayout)(C.malloc(C.size_t(C.sizeof_VkAttachmentReferenceStencilLayout)))
	*p = C.VkAttachmentReferenceStencilLayout{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkImageLayout(s.StencilLayout)
	p.stencilLayout = val0
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type BlitImageInfo2 struct {
	Next           Structure
	SrcImage       *Image
	SrcImageLayout ImageLayout
	DstImage       *Image
	DstImageLayout ImageLayout
	Regions        []ImageBlit2
	Filter         Filter
}

func (s *BlitImageInfo2) GetType() StructureType {
	return StructureTypeBlitImageInfo2
}

func (s *BlitImageInfo2) toC() (*C.VkBlitImageInfo2, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkBlitImageInfo2)(C.malloc(C.size_t(C.sizeof_VkBlitImageInfo2)))
	*p = C.VkBlitImageInfo2{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	h0 := C.VkImage(unsafe.Pointer(s.SrcImage.handle))
	p.srcImage = h0
	val1 := C.VkImageLayout(s.SrcImageLayout)
	p.srcImageLayout = val1
	h2 := C.VkImage(unsafe.Pointer(s.DstImage.handle))
	p.dstImage = h2
	val3 := C.VkImageLayout(s.DstImageLayout)
	p.dstImageLayout = val3
	len4 := len(s.Regions)

	var arr5 *C.VkImageBlit2
	if len4 > 0 {
		arr5 = (*C.VkImageBlit2)(C.malloc(C.size_t(len4) * C.size_t(unsafe.Sizeof(*new(C.VkImageBlit2)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr5)) })
	}
	for i6, elem7 := range s.Regions {
		val8, cancel9 := elem7.toC()
		cancels = append(cancels, cancel9)
		(*[1 << 30]C.VkImageBlit2)(unsafe.Pointer(arr5))[i6] = *val8
	}
	p.pRegions = arr5
	p.regionCount = C.uint32_t(len(s.Regions))
	val10 := C.VkFilter(s.Filter)
	p.filter = val10
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ClearAttachment struct {
	AspectMask      ImageAspectFlags
	ColorAttachment uint32
	ClearValue      ClearValue
}

func (s *ClearAttachment) toC() (*C.VkClearAttachment, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkClearAttachment)(C.malloc(C.size_t(C.sizeof_VkClearAttachment)))
	*p = C.VkClearAttachment{}
	val0 := C.VkImageAspectFlags(s.AspectMask)
	p.aspectMask = val0
	val1 := C.uint32_t(s.ColorAttachment)
	p.colorAttachment = val1
	val2, cancel3 := s.ClearValue.toC()
	cancels = append(cancels, cancel3)
	p.clearValue = *val2
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ClearColorValue [C.sizeof_VkClearColorValue]byte

func NewClearColorValueFloat32(val [4]float32) ClearColorValue {
	var u ClearColorValue
	*(*[4]float32)(unsafe.Pointer(&u[0])) = val
	return u
}

func NewClearColorValueInt32(val [4]int32) ClearColorValue {
	var u ClearColorValue
	*(*[4]int32)(unsafe.Pointer(&u[0])) = val
	return u
}

func NewClearColorValueUint32(val [4]uint32) ClearColorValue {
	var u ClearColorValue
	*(*[4]uint32)(unsafe.Pointer(&u[0])) = val
	return u
}

func (u *ClearColorValue) Float32() [4]float32 {
	return *(*[4]float32)(unsafe.Pointer(&u[0]))
}

func (u *ClearColorValue) Int32() [4]int32 {
	return *(*[4]int32)(unsafe.Pointer(&u[0]))
}

func (u *ClearColorValue) Uint32() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&u[0]))
}

func (s *ClearColorValue) toC() (*C.VkClearColorValue, func()) {
	p := (*C.VkClearColorValue)(C.malloc(C.size_t(C.sizeof_VkClearColorValue)))
	C.memcpy(unsafe.Pointer(p), unsafe.Pointer(&s[0]), C.sizeof_VkClearColorValue)
	return p, func() { C.free(unsafe.Pointer(p)) }
}

type ClearDepthStencilValue struct {
	Depth   float32
	Stencil uint32
}

func (s *ClearDepthStencilValue) toC() (*C.VkClearDepthStencilValue, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkClearDepthStencilValue)(C.malloc(C.size_t(C.sizeof_VkClearDepthStencilValue)))
	*p = C.VkClearDepthStencilValue{}
	val0 := C.float(s.Depth)
	p.depth = val0
	val1 := C.uint32_t(s.Stencil)
	p.stencil = val1
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ClearRect struct {
	Rect           Rect2D
	BaseArrayLayer uint32
	LayerCount     uint32
}

func (s *ClearRect) toC() (*C.VkClearRect, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkClearRect)(C.malloc(C.size_t(C.sizeof_VkClearRect)))
	*p = C.VkClearRect{}
	val0, cancel1 := s.Rect.toC()
	cancels = append(cancels, cancel1)
	p.rect = *val0
	val2 := C.uint32_t(s.BaseArrayLayer)
	p.baseArrayLayer = val2
	val3 := C.uint32_t(s.LayerCount)
	p.layerCount = val3
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ClearValue [C.sizeof_VkClearValue]byte

func NewClearValueColor(val ClearColorValue) ClearValue {
	var u ClearValue
	*(*ClearColorValue)(unsafe.Pointer(&u[0])) = val
	return u
}

func NewClearValueDepthStencil(val ClearDepthStencilValue) ClearValue {
	var u ClearValue
	*(*ClearDepthStencilValue)(unsafe.Pointer(&u[0])) = val
	return u
}

func (u *ClearValue) Color() ClearColorValue {
	return *(*ClearColorValue)(unsafe.Pointer(&u[0]))
}

func (u *ClearValue) DepthStencil() ClearDepthStencilValue {
	return *(*ClearDepthStencilValue)(unsafe.Pointer(&u[0]))
}

func (s *ClearValue) toC() (*C.VkClearValue, func()) {
	p := (*C.VkClearValue)(C.malloc(C.size_t(C.sizeof_VkClearValue)))
	C.memcpy(unsafe.Pointer(p), unsafe.Pointer(&s[0]), C.sizeof_VkClearValue)
	return p, func() { C.free(unsafe.Pointer(p)) }
}

type CommandBufferInheritanceRenderingInfo struct {
	Next                    Structure
	Flags                   RenderingFlags
	ViewMask                uint32
	ColorAttachmentFormats  []Format
	DepthAttachmentFormat   Format
	StencilAttachmentFormat Format
	RasterizationSamples    SampleCountFlagBits
}

func (s *CommandBufferInheritanceRenderingInfo) GetType() StructureType {
	return StructureTypeCommandBufferInheritanceRenderingInfo
}

func (s *CommandBufferInheritanceRenderingInfo) toC() (*C.VkCommandBufferInheritanceRenderingInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkCommandBufferInheritanceRenderingInfo)(C.malloc(C.size_t(C.sizeof_VkCommandBufferInheritanceRenderingInfo)))
	*p = C.VkCommandBufferInheritanceRenderingInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkRenderingFlags(s.Flags)
	p.flags = val0
	val1 := C.uint32_t(s.ViewMask)
	p.viewMask = val1
	len2 := len(s.ColorAttachmentFormats)

	var arr3 *C.VkFormat
	if len2 > 0 {
		arr3 = (*C.VkFormat)(C.malloc(C.size_t(len2) * C.size_t(unsafe.Sizeof(*new(C.VkFormat)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr3)) })
	}
	for i4, elem5 := range s.ColorAttachmentFormats {
		val6 := C.VkFormat(elem5)
		(*[1 << 30]C.VkFormat)(unsafe.Pointer(arr3))[i4] = val6
	}
	p.pColorAttachmentFormats = arr3
	p.colorAttachmentCount = C.uint32_t(len(s.ColorAttachmentFormats))
	val7 := C.VkFormat(s.DepthAttachmentFormat)
	p.depthAttachmentFormat = val7
	val8 := C.VkFormat(s.StencilAttachmentFormat)
	p.stencilAttachmentFormat = val8
	val9 := C.VkSampleCountFlagBits(s.RasterizationSamples)
	p.rasterizationSamples = val9
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type DeviceGroupRenderPassBeginInfo struct {
	Next              Structure
	DeviceMask        uint32
	DeviceRenderAreas []Rect2D
}

func (s *DeviceGroupRenderPassBeginInfo) GetType() StructureType {
	return StructureTypeDeviceGroupRenderPassBeginInfo
}

func (s *DeviceGroupRenderPassBeginInfo) toC() (*C.VkDeviceGroupRenderPassBeginInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkDeviceGroupRenderPassBeginInfo)(C.malloc(C.size_t(C.sizeof_VkDeviceGroupRenderPassBeginInfo)))
	*p = C.VkDeviceGroupRenderPassBeginInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.uint32_t(s.DeviceMask)
	p.deviceMask = val0
	len1 := len(s.DeviceRenderAreas)

	var arr2 *C.VkRect2D
	if len1 > 0 {
		arr2 = (*C.VkRect2D)(C.malloc(C.size_t(len1) * C.size_t(unsafe.Sizeof(*new(C.VkRect2D)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr2)) })
	}
	for i3, elem4 := range s.DeviceRenderAreas {
		val5, cancel6 := elem4.toC()
		cancels = append(cancels, cancel6)
		(*[1 << 30]C.VkRect2D)(unsafe.Pointer(arr2))[i3] = *val5
	}
	p.pDeviceRenderAreas = arr2
	p.deviceRenderAreaCount = C.uint32_t(len(s.DeviceRenderAreas))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type DrawIndexedIndirectCommand struct {
	IndexCount    uint32
	InstanceCount uint32
	FirstIndex    uint32
	VertexOffset  int32
	FirstInstance uint32
}

func (s *DrawIndexedIndirectCommand) toC() (*C.VkDrawIndexedIndirectCommand, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkDrawIndexedIndirectCommand)(C.malloc(C.size_t(C.sizeof_VkDrawIndexedIndirectCommand)))
	*p = C.VkDrawIndexedIndirectCommand{}
	val0 := C.uint32_t(s.IndexCount)
	p.indexCount = val0
	val1 := C.uint32_t(s.InstanceCount)
	p.instanceCount = val1
	val2 := C.uint32_t(s.FirstIndex)
	p.firstIndex = val2
	val3 := C.int32_t(s.VertexOffset)
	p.vertexOffset = val3
	val4 := C.uint32_t(s.FirstInstance)
	p.firstInstance = val4
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type DrawIndirectCommand struct {
	VertexCount   uint32
	InstanceCount uint32
	FirstVertex   uint32
	FirstInstance uint32
}

func (s *DrawIndirectCommand) toC() (*C.VkDrawIndirectCommand, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkDrawIndirectCommand)(C.malloc(C.size_t(C.sizeof_VkDrawIndirectCommand)))
	*p = C.VkDrawIndirectCommand{}
	val0 := C.uint32_t(s.VertexCount)
	p.vertexCount = val0
	val1 := C.uint32_t(s.InstanceCount)
	p.instanceCount = val1
	val2 := C.uint32_t(s.FirstVertex)
	p.firstVertex = val2
	val3 := C.uint32_t(s.FirstInstance)
	p.firstInstance = val3
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type Extent2D struct {
	Width  uint32
	Height uint32
}

func (s *Extent2D) toC() (*C.VkExtent2D, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkExtent2D)(C.malloc(C.size_t(C.sizeof_VkExtent2D)))
	*p = C.VkExtent2D{}
	val0 := C.uint32_t(s.Width)
	p.width = val0
	val1 := C.uint32_t(s.Height)
	p.height = val1
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type Extent3D struct {
	Width  uint32
	Height uint32
	Depth  uint32
}

func (s *Extent3D) toC() (*C.VkExtent3D, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkExtent3D)(C.malloc(C.size_t(C.sizeof_VkExtent3D)))
	*p = C.VkExtent3D{}
	val0 := C.uint32_t(s.Width)
	p.width = val0
	val1 := C.uint32_t(s.Height)
	p.height = val1
	val2 := C.uint32_t(s.Depth)
	p.depth = val2
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type FramebufferAttachmentImageInfo struct {
	Next        Structure
	Flags       ImageCreateFlags
	Usage       ImageUsageFlags
	Width       uint32
	Height      uint32
	LayerCount  uint32
	ViewFormats []Format
}

func (s *FramebufferAttachmentImageInfo) GetType() StructureType {
	return StructureTypeFramebufferAttachmentImageInfo
}

func (s *FramebufferAttachmentImageInfo) toC() (*C.VkFramebufferAttachmentImageInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkFramebufferAttachmentImageInfo)(C.malloc(C.size_t(C.sizeof_VkFramebufferAttachmentImageInfo)))
	*p = C.VkFramebufferAttachmentImageInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkImageCreateFlags(s.Flags)
	p.flags = val0
	val1 := C.VkImageUsageFlags(s.Usage)
	p.usage = val1
	val2 := C.uint32_t(s.Width)
	p.width = val2
	val3 := C.uint32_t(s.Height)
	p.height = val3
	val4 := C.uint32_t(s.LayerCount)
	p.layerCount = val4
	len5 := len(s.ViewFormats)

	var arr6 *C.VkFormat
	if len5 > 0 {
		arr6 = (*C.VkFormat)(C.malloc(C.size_t(len5) * C.size_t(unsafe.Sizeof(*new(C.VkFormat)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr6)) })
	}
	for i7, elem8 := range s.ViewFormats {
		val9 := C.VkFormat(elem8)
		(*[1 << 30]C.VkFormat)(unsafe.Pointer(arr6))[i7] = val9
	}
	p.pViewFormats = arr6
	p.viewFormatCount = C.uint32_t(len(s.ViewFormats))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type FramebufferAttachmentsCreateInfo struct {
	Next                 Structure
	AttachmentImageInfos []FramebufferAttachmentImageInfo
}

func (s *FramebufferAttachmentsCreateInfo) GetType() StructureType {
	return StructureTypeFramebufferAttachmentsCreateInfo
}

func (s *FramebufferAttachmentsCreateInfo) toC() (*C.VkFramebufferAttachmentsCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkFramebufferAttachmentsCreateInfo)(C.malloc(C.size_t(C.sizeof_VkFramebufferAttachmentsCreateInfo)))
	*p = C.VkFramebufferAttachmentsCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	len0 := len(s.AttachmentImageInfos)

	var arr1 *C.VkFramebufferAttachmentImageInfo
	if len0 > 0 {
		arr1 = (*C.VkFramebufferAttachmentImageInfo)(C.malloc(C.size_t(len0) * C.size_t(unsafe.Sizeof(*new(C.VkFramebufferAttachmentImageInfo)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr1)) })
	}
	for i2, elem3 := range s.AttachmentImageInfos {
		val4, cancel5 := elem3.toC()
		cancels = append(cancels, cancel5)
		(*[1 << 30]C.VkFramebufferAttachmentImageInfo)(unsafe.Pointer(arr1))[i2] = *val4
	}
	p.pAttachmentImageInfos = arr1
	p.attachmentImageInfoCount = C.uint32_t(len(s.AttachmentImageInfos))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type FramebufferCreateInfo struct {
	Next        Structure
	Flags       FramebufferCreateFlags
	RenderPass  *RenderPass
	Attachments []*ImageView
	Width       uint32
	Height      uint32
	Layers      uint32
}

func (s *FramebufferCreateInfo) GetType() StructureType {
	return StructureTypeFramebufferCreateInfo
}

func (s *FramebufferCreateInfo) toC() (*C.VkFramebufferCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkFramebufferCreateInfo)(C.malloc(C.size_t(C.sizeof_VkFramebufferCreateInfo)))
	*p = C.VkFramebufferCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkFramebufferCreateFlags(s.Flags)
	p.flags = val0
	h1 := C.VkRenderPass(unsafe.Pointer(s.RenderPass.handle))
	p.renderPass = h1
	len2 := len(s.Attachments)

	var arr3 *C.VkImageView
	if len2 > 0 {
		arr3 = (*C.VkImageView)(C.malloc(C.size_t(len2) * C.size_t(unsafe.Sizeof(*new(C.VkImageView)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr3)) })
	}
	for i4, elem5 := range s.Attachments {
		h6 := C.VkImageView(unsafe.Pointer(elem5.handle))
		(*[1 << 30]C.VkImageView)(unsafe.Pointer(arr3))[i4] = h6
	}
	p.pAttachments = arr3
	p.attachmentCount = C.uint32_t(len(s.Attachments))
	val7 := C.uint32_t(s.Width)
	p.width = val7
	val8 := C.uint32_t(s.Height)
	p.height = val8
	val9 := C.uint32_t(s.Layers)
	p.layers = val9
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type GraphicsPipelineCreateInfo struct {
	Next               Structure
	Flags              PipelineCreateFlags
	Stages             []PipelineShaderStageCreateInfo
	VertexInputState   *PipelineVertexInputStateCreateInfo
	InputAssemblyState *PipelineInputAssemblyStateCreateInfo
	TessellationState  *PipelineTessellationStateCreateInfo
	ViewportState      *PipelineViewportStateCreateInfo
	RasterizationState *PipelineRasterizationStateCreateInfo
	MultisampleState   *PipelineMultisampleStateCreateInfo
	DepthStencilState  *PipelineDepthStencilStateCreateInfo
	ColorBlendState    *PipelineColorBlendStateCreateInfo
	DynamicState       *PipelineDynamicStateCreateInfo
	Layout             *PipelineLayout
	RenderPass         *RenderPass
	Subpass            uint32
	BasePipelineHandle *Pipeline
	BasePipelineIndex  int32
}

func (s *GraphicsPipelineCreateInfo) GetType() StructureType {
	return StructureTypeGraphicsPipelineCreateInfo
}

func (s *GraphicsPipelineCreateInfo) toC() (*C.VkGraphicsPipelineCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkGraphicsPipelineCreateInfo)(C.malloc(C.size_t(C.sizeof_VkGraphicsPipelineCreateInfo)))
	*p = C.VkGraphicsPipelineCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPipelineCreateFlags(s.Flags)
	p.flags = val0
	len1 := len(s.Stages)

	var arr2 *C.VkPipelineShaderStageCreateInfo
	if len1 > 0 {
		arr2 = (*C.VkPipelineShaderStageCreateInfo)(C.malloc(C.size_t(len1) * C.size_t(unsafe.Sizeof(*new(C.VkPipelineShaderStageCreateInfo)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr2)) })
	}
	for i3, elem4 := range s.Stages {
		val5, cancel6 := elem4.toC()
		cancels = append(cancels, cancel6)
		(*[1 << 30]C.VkPipelineShaderStageCreateInfo)(unsafe.Pointer(arr2))[i3] = *val5
	}
	p.pStages = arr2
	p.stageCount = C.uint32_t(len(s.Stages))
	var ptr7 *C.VkPipelineVertexInputStateCreateInfo
	if s.VertexInputState != nil {
		val8, cancel9 := s.VertexInputState.toC()
		cancels = append(cancels, cancel9)
		ptr7 = val8
	}
	p.pVertexInputState = ptr7
	var ptr10 *C.VkPipelineInputAssemblyStateCreateInfo
	if s.InputAssemblyState != nil {
		val11, cancel12 := s.InputAssemblyState.toC()
		cancels = append(cancels, cancel12)
		ptr10 = val11
	}
	p.pInputAssemblyState = ptr10
	var ptr13 *C.VkPipelineTessellationStateCreateInfo
	if s.TessellationState != nil {
		val14, cancel15 := s.TessellationState.toC()
		cancels = append(cancels, cancel15)
		ptr13 = val14
	}
	p.pTessellationState = ptr13
	var ptr16 *C.VkPipelineViewportStateCreateInfo
	if s.ViewportState != nil {
		val17, cancel18 := s.ViewportState.toC()
		cancels = append(cancels, cancel18)
		ptr16 = val17
	}
	p.pViewportState = ptr16
	var ptr19 *C.VkPipelineRasterizationStateCreateInfo
	if s.RasterizationState != nil {
		val20, cancel21 := s.RasterizationState.toC()
		cancels = append(cancels, cancel21)
		ptr19 = val20
	}
	p.pRasterizationState = ptr19
	var ptr22 *C.VkPipelineMultisampleStateCreateInfo
	if s.MultisampleState != nil {
		val23, cancel24 := s.MultisampleState.toC()
		cancels = append(cancels, cancel24)
		ptr22 = val23
	}
	p.pMultisampleState = ptr22
	var ptr25 *C.VkPipelineDepthStencilStateCreateInfo
	if s.DepthStencilState != nil {
		val26, cancel27 := s.DepthStencilState.toC()
		cancels = append(cancels, cancel27)
		ptr25 = val26
	}
	p.pDepthStencilState = ptr25
	var ptr28 *C.VkPipelineColorBlendStateCreateInfo
	if s.ColorBlendState != nil {
		val29, cancel30 := s.ColorBlendState.toC()
		cancels = append(cancels, cancel30)
		ptr28 = val29
	}
	p.pColorBlendState = ptr28
	var ptr31 *C.VkPipelineDynamicStateCreateInfo
	if s.DynamicState != nil {
		val32, cancel33 := s.DynamicState.toC()
		cancels = append(cancels, cancel33)
		ptr31 = val32
	}
	p.pDynamicState = ptr31
	h34 := C.VkPipelineLayout(unsafe.Pointer(s.Layout.handle))
	p.layout = h34
	h35 := C.VkRenderPass(unsafe.Pointer(s.RenderPass.handle))
	p.renderPass = h35
	val36 := C.uint32_t(s.Subpass)
	p.subpass = val36
	h37 := C.VkPipeline(unsafe.Pointer(s.BasePipelineHandle.handle))
	p.basePipelineHandle = h37
	val38 := C.int32_t(s.BasePipelineIndex)
	p.basePipelineIndex = val38
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ImageBlit struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffsets     [2]Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffsets     [2]Offset3D
}

func (s *ImageBlit) toC() (*C.VkImageBlit, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkImageBlit)(C.malloc(C.size_t(C.sizeof_VkImageBlit)))
	*p = C.VkImageBlit{}
	val0, cancel1 := s.SrcSubresource.toC()
	cancels = append(cancels, cancel1)
	p.srcSubresource = *val0
	var arr2 [2]C.VkOffset3D
	for i3, elem4 := range s.SrcOffsets {
		val5, cancel6 := elem4.toC()
		cancels = append(cancels, cancel6)
		arr2[i3] = *val5
	}
	p.srcOffsets = arr2
	val7, cancel8 := s.DstSubresource.toC()
	cancels = append(cancels, cancel8)
	p.dstSubresource = *val7
	var arr9 [2]C.VkOffset3D
	for i10, elem11 := range s.DstOffsets {
		val12, cancel13 := elem11.toC()
		cancels = append(cancels, cancel13)
		arr9[i10] = *val12
	}
	p.dstOffsets = arr9
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ImageBlit2 struct {
	Next           Structure
	SrcSubresource ImageSubresourceLayers
	SrcOffsets     [2]Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffsets     [2]Offset3D
}

func (s *ImageBlit2) GetType() StructureType {
	return StructureTypeImageBlit2
}

func (s *ImageBlit2) toC() (*C.VkImageBlit2, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkImageBlit2)(C.malloc(C.size_t(C.sizeof_VkImageBlit2)))
	*p = C.VkImageBlit2{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0, cancel1 := s.SrcSubresource.toC()
	cancels = append(cancels, cancel1)
	p.srcSubresource = *val0
	var arr2 [2]C.VkOffset3D
	for i3, elem4 := range s.SrcOffsets {
		val5, cancel6 := elem4.toC()
		cancels = append(cancels, cancel6)
		arr2[i3] = *val5
	}
	p.srcOffsets = arr2
	val7, cancel8 := s.DstSubresource.toC()
	cancels = append(cancels, cancel8)
	p.dstSubresource = *val7
	var arr9 [2]C.VkOffset3D
	for i10, elem11 := range s.DstOffsets {
		val12, cancel13 := elem11.toC()
		cancels = append(cancels, cancel13)
		arr9[i10] = *val12
	}
	p.dstOffsets = arr9
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ImageResolve struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffset      Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffset      Offset3D
	Extent         Extent3D
}

func (s *ImageResolve) toC() (*C.VkImageResolve, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkImageResolve)(C.malloc(C.size_t(C.sizeof_VkImageResolve)))
	*p = C.VkImageResolve{}
	val0, cancel1 := s.SrcSubresource.toC()
	cancels = append(cancels, cancel1)
	p.srcSubresource = *val0
	val2, cancel3 := s.SrcOffset.toC()
	cancels = append(cancels, cancel3)
	p.srcOffset = *val2
	val4, cancel5 := s.DstSubresource.toC()
	cancels = append(cancels, cancel5)
	p.dstSubresource = *val4
	val6, cancel7 := s.DstOffset.toC()
	cancels = append(cancels, cancel7)
	p.dstOffset = *val6
	val8, cancel9 := s.Extent.toC()
	cancels = append(cancels, cancel9)
	p.extent = *val8
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ImageResolve2 struct {
	Next           Structure
	SrcSubresource ImageSubresourceLayers
	SrcOffset      Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffset      Offset3D
	Extent         Extent3D
}

func (s *ImageResolve2) GetType() StructureType {
	return StructureTypeImageResolve2
}

func (s *ImageResolve2) toC() (*C.VkImageResolve2, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkImageResolve2)(C.malloc(C.size_t(C.sizeof_VkImageResolve2)))
	*p = C.VkImageResolve2{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0, cancel1 := s.SrcSubresource.toC()
	cancels = append(cancels, cancel1)
	p.srcSubresource = *val0
	val2, cancel3 := s.SrcOffset.toC()
	cancels = append(cancels, cancel3)
	p.srcOffset = *val2
	val4, cancel5 := s.DstSubresource.toC()
	cancels = append(cancels, cancel5)
	p.dstSubresource = *val4
	val6, cancel7 := s.DstOffset.toC()
	cancels = append(cancels, cancel7)
	p.dstOffset = *val6
	val8, cancel9 := s.Extent.toC()
	cancels = append(cancels, cancel9)
	p.extent = *val8
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ImageStencilUsageCreateInfo struct {
	Next         Structure
	StencilUsage ImageUsageFlags
}

func (s *ImageStencilUsageCreateInfo) GetType() StructureType {
	return StructureTypeImageStencilUsageCreateInfo
}

func (s *ImageStencilUsageCreateInfo) toC() (*C.VkImageStencilUsageCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkImageStencilUsageCreateInfo)(C.malloc(C.size_t(C.sizeof_VkImageStencilUsageCreateInfo)))
	*p = C.VkImageStencilUsageCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkImageUsageFlags(s.StencilUsage)
	p.stencilUsage = val0
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ImageSubresourceLayers struct {
	AspectMask     ImageAspectFlags
	MipLevel       uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

func (s *ImageSubresourceLayers) toC() (*C.VkImageSubresourceLayers, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkImageSubresourceLayers)(C.malloc(C.size_t(C.sizeof_VkImageSubresourceLayers)))
	*p = C.VkImageSubresourceLayers{}
	val0 := C.VkImageAspectFlags(s.AspectMask)
	p.aspectMask = val0
	val1 := C.uint32_t(s.MipLevel)
	p.mipLevel = val1
	val2 := C.uint32_t(s.BaseArrayLayer)
	p.baseArrayLayer = val2
	val3 := C.uint32_t(s.LayerCount)
	p.layerCount = val3
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ImageSubresourceRange struct {
	AspectMask     ImageAspectFlags
	BaseMipLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

func (s *ImageSubresourceRange) toC() (*C.VkImageSubresourceRange, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkImageSubresourceRange)(C.malloc(C.size_t(C.sizeof_VkImageSubresourceRange)))
	*p = C.VkImageSubresourceRange{}
	val0 := C.VkImageAspectFlags(s.AspectMask)
	p.aspectMask = val0
	val1 := C.uint32_t(s.BaseMipLevel)
	p.baseMipLevel = val1
	val2 := C.uint32_t(s.LevelCount)
	p.levelCount = val2
	val3 := C.uint32_t(s.BaseArrayLayer)
	p.baseArrayLayer = val3
	val4 := C.uint32_t(s.LayerCount)
	p.layerCount = val4
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type InputAttachmentAspectReference struct {
	Subpass              uint32
	InputAttachmentIndex uint32
	AspectMask           ImageAspectFlags
}

func (s *InputAttachmentAspectReference) toC() (*C.VkInputAttachmentAspectReference, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkInputAttachmentAspectReference)(C.malloc(C.size_t(C.sizeof_VkInputAttachmentAspectReference)))
	*p = C.VkInputAttachmentAspectReference{}
	val0 := C.uint32_t(s.Subpass)
	p.subpass = val0
	val1 := C.uint32_t(s.InputAttachmentIndex)
	p.inputAttachmentIndex = val1
	val2 := C.VkImageAspectFlags(s.AspectMask)
	p.aspectMask = val2
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type Offset2D struct {
	X int32
	Y int32
}

func (s *Offset2D) toC() (*C.VkOffset2D, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkOffset2D)(C.malloc(C.size_t(C.sizeof_VkOffset2D)))
	*p = C.VkOffset2D{}
	val0 := C.int32_t(s.X)
	p.x = val0
	val1 := C.int32_t(s.Y)
	p.y = val1
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type Offset3D struct {
	X int32
	Y int32
	Z int32
}

func (s *Offset3D) toC() (*C.VkOffset3D, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkOffset3D)(C.malloc(C.size_t(C.sizeof_VkOffset3D)))
	*p = C.VkOffset3D{}
	val0 := C.int32_t(s.X)
	p.x = val0
	val1 := C.int32_t(s.Y)
	p.y = val1
	val2 := C.int32_t(s.Z)
	p.z = val2
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceDepthStencilResolveProperties struct {
	Next                         Structure
	SupportedDepthResolveModes   ResolveModeFlags
	SupportedStencilResolveModes ResolveModeFlags
	IndependentResolveNone       bool
	IndependentResolve           bool
}

func (s *PhysicalDeviceDepthStencilResolveProperties) GetType() StructureType {
	return StructureTypePhysicalDeviceDepthStencilResolveProperties
}

func (s *PhysicalDeviceDepthStencilResolveProperties) toC() (*C.VkPhysicalDeviceDepthStencilResolveProperties, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceDepthStencilResolveProperties)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceDepthStencilResolveProperties)))
	*p = C.VkPhysicalDeviceDepthStencilResolveProperties{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkResolveModeFlags(s.SupportedDepthResolveModes)
	p.supportedDepthResolveModes = val0
	val1 := C.VkResolveModeFlags(s.SupportedStencilResolveModes)
	p.supportedStencilResolveModes = val1
	val2 := C.VkBool32(0)
	if s.IndependentResolveNone {
		val2 = C.VkBool32(1)
	}
	p.independentResolveNone = val2
	val3 := C.VkBool32(0)
	if s.IndependentResolve {
		val3 = C.VkBool32(1)
	}
	p.independentResolve = val3
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceDynamicRenderingFeatures struct {
	Next             Structure
	DynamicRendering bool
}

func (s *PhysicalDeviceDynamicRenderingFeatures) GetType() StructureType {
	return StructureTypePhysicalDeviceDynamicRenderingFeatures
}

func (s *PhysicalDeviceDynamicRenderingFeatures) toC() (*C.VkPhysicalDeviceDynamicRenderingFeatures, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceDynamicRenderingFeatures)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceDynamicRenderingFeatures)))
	*p = C.VkPhysicalDeviceDynamicRenderingFeatures{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkBool32(0)
	if s.DynamicRendering {
		val0 = C.VkBool32(1)
	}
	p.dynamicRendering = val0
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceDynamicRenderingLocalReadFeatures struct {
	Next                      Structure
	DynamicRenderingLocalRead bool
}

func (s *PhysicalDeviceDynamicRenderingLocalReadFeatures) GetType() StructureType {
	return StructureTypePhysicalDeviceDynamicRenderingLocalReadFeatures
}

func (s *PhysicalDeviceDynamicRenderingLocalReadFeatures) toC() (*C.VkPhysicalDeviceDynamicRenderingLocalReadFeatures, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceDynamicRenderingLocalReadFeatures)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceDynamicRenderingLocalReadFeatures)))
	*p = C.VkPhysicalDeviceDynamicRenderingLocalReadFeatures{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkBool32(0)
	if s.DynamicRenderingLocalRead {
		val0 = C.VkBool32(1)
	}
	p.dynamicRenderingLocalRead = val0
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceImagelessFramebufferFeatures struct {
	Next                 Structure
	ImagelessFramebuffer bool
}

func (s *PhysicalDeviceImagelessFramebufferFeatures) GetType() StructureType {
	return StructureTypePhysicalDeviceImagelessFramebufferFeatures
}

func (s *PhysicalDeviceImagelessFramebufferFeatures) toC() (*C.VkPhysicalDeviceImagelessFramebufferFeatures, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceImagelessFramebufferFeatures)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceImagelessFramebufferFeatures)))
	*p = C.VkPhysicalDeviceImagelessFramebufferFeatures{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkBool32(0)
	if s.ImagelessFramebuffer {
		val0 = C.VkBool32(1)
	}
	p.imagelessFramebuffer = val0
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceLineRasterizationFeatures struct {
	Next                     Structure
	RectangularLines         bool
	BresenhamLines           bool
	SmoothLines              bool
	StippledRectangularLines bool
	StippledBresenhamLines   bool
	StippledSmoothLines      bool
}

func (s *PhysicalDeviceLineRasterizationFeatures) GetType() StructureType {
	return StructureTypePhysicalDeviceLineRasterizationFeatures
}

func (s *PhysicalDeviceLineRasterizationFeatures) toC() (*C.VkPhysicalDeviceLineRasterizationFeatures, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceLineRasterizationFeatures)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceLineRasterizationFeatures)))
	*p = C.VkPhysicalDeviceLineRasterizationFeatures{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkBool32(0)
	if s.RectangularLines {
		val0 = C.VkBool32(1)
	}
	p.rectangularLines = val0
	val1 := C.VkBool32(0)
	if s.BresenhamLines {
		val1 = C.VkBool32(1)
	}
	p.bresenhamLines = val1
	val2 := C.VkBool32(0)
	if s.SmoothLines {
		val2 = C.VkBool32(1)
	}
	p.smoothLines = val2
	val3 := C.VkBool32(0)
	if s.StippledRectangularLines {
		val3 = C.VkBool32(1)
	}
	p.stippledRectangularLines = val3
	val4 := C.VkBool32(0)
	if s.StippledBresenhamLines {
		val4 = C.VkBool32(1)
	}
	p.stippledBresenhamLines = val4
	val5 := C.VkBool32(0)
	if s.StippledSmoothLines {
		val5 = C.VkBool32(1)
	}
	p.stippledSmoothLines = val5
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceLineRasterizationProperties struct {
	Next                      Structure
	LineSubPixelPrecisionBits uint32
}

func (s *PhysicalDeviceLineRasterizationProperties) GetType() StructureType {
	return StructureTypePhysicalDeviceLineRasterizationProperties
}

func (s *PhysicalDeviceLineRasterizationProperties) toC() (*C.VkPhysicalDeviceLineRasterizationProperties, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceLineRasterizationProperties)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceLineRasterizationProperties)))
	*p = C.VkPhysicalDeviceLineRasterizationProperties{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.uint32_t(s.LineSubPixelPrecisionBits)
	p.lineSubPixelPrecisionBits = val0
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceMultiviewFeatures struct {
	Next                        Structure
	Multiview                   bool
	MultiviewGeometryShader     bool
	MultiviewTessellationShader bool
}

func (s *PhysicalDeviceMultiviewFeatures) GetType() StructureType {
	return StructureTypePhysicalDeviceMultiviewFeatures
}

func (s *PhysicalDeviceMultiviewFeatures) toC() (*C.VkPhysicalDeviceMultiviewFeatures, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceMultiviewFeatures)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceMultiviewFeatures)))
	*p = C.VkPhysicalDeviceMultiviewFeatures{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkBool32(0)
	if s.Multiview {
		val0 = C.VkBool32(1)
	}
	p.multiview = val0
	val1 := C.VkBool32(0)
	if s.MultiviewGeometryShader {
		val1 = C.VkBool32(1)
	}
	p.multiviewGeometryShader = val1
	val2 := C.VkBool32(0)
	if s.MultiviewTessellationShader {
		val2 = C.VkBool32(1)
	}
	p.multiviewTessellationShader = val2
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceMultiviewProperties struct {
	Next                      Structure
	MaxMultiviewViewCount     uint32
	MaxMultiviewInstanceIndex uint32
}

func (s *PhysicalDeviceMultiviewProperties) GetType() StructureType {
	return StructureTypePhysicalDeviceMultiviewProperties
}

func (s *PhysicalDeviceMultiviewProperties) toC() (*C.VkPhysicalDeviceMultiviewProperties, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceMultiviewProperties)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceMultiviewProperties)))
	*p = C.VkPhysicalDeviceMultiviewProperties{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.uint32_t(s.MaxMultiviewViewCount)
	p.maxMultiviewViewCount = val0
	val1 := C.uint32_t(s.MaxMultiviewInstanceIndex)
	p.maxMultiviewInstanceIndex = val1
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDevicePointClippingProperties struct {
	Next                  Structure
	PointClippingBehavior PointClippingBehavior
}

func (s *PhysicalDevicePointClippingProperties) GetType() StructureType {
	return StructureTypePhysicalDevicePointClippingProperties
}

func (s *PhysicalDevicePointClippingProperties) toC() (*C.VkPhysicalDevicePointClippingProperties, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDevicePointClippingProperties)(C.malloc(C.size_t(C.sizeof_VkPhysicalDevicePointClippingProperties)))
	*p = C.VkPhysicalDevicePointClippingProperties{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPointClippingBehavior(s.PointClippingBehavior)
	p.pointClippingBehavior = val0
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceSeparateDepthStencilLayoutsFeatures struct {
	Next                        Structure
	SeparateDepthStencilLayouts bool
}

func (s *PhysicalDeviceSeparateDepthStencilLayoutsFeatures) GetType() StructureType {
	return StructureTypePhysicalDeviceSeparateDepthStencilLayoutsFeatures
}

func (s *PhysicalDeviceSeparateDepthStencilLayoutsFeatures) toC() (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures)))
	*p = C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkBool32(0)
	if s.SeparateDepthStencilLayouts {
		val0 = C.VkBool32(1)
	}
	p.separateDepthStencilLayouts = val0
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceShaderDrawParameterFeatures struct {
}

func (s *PhysicalDeviceShaderDrawParameterFeatures) toC() (*C.VkPhysicalDeviceShaderDrawParameterFeatures, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceShaderDrawParameterFeatures)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceShaderDrawParameterFeatures)))
	*p = C.VkPhysicalDeviceShaderDrawParameterFeatures{}
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceShaderDrawParametersFeatures struct {
	Next                 Structure
	ShaderDrawParameters bool
}

func (s *PhysicalDeviceShaderDrawParametersFeatures) GetType() StructureType {
	return StructureTypePhysicalDeviceShaderDrawParametersFeatures
}

func (s *PhysicalDeviceShaderDrawParametersFeatures) toC() (*C.VkPhysicalDeviceShaderDrawParametersFeatures, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceShaderDrawParametersFeatures)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceShaderDrawParametersFeatures)))
	*p = C.VkPhysicalDeviceShaderDrawParametersFeatures{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkBool32(0)
	if s.ShaderDrawParameters {
		val0 = C.VkBool32(1)
	}
	p.shaderDrawParameters = val0
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceVertexAttributeDivisorFeatures struct {
	Next                                   Structure
	VertexAttributeInstanceRateDivisor     bool
	VertexAttributeInstanceRateZeroDivisor bool
}

func (s *PhysicalDeviceVertexAttributeDivisorFeatures) GetType() StructureType {
	return StructureTypePhysicalDeviceVertexAttributeDivisorFeatures
}

func (s *PhysicalDeviceVertexAttributeDivisorFeatures) toC() (*C.VkPhysicalDeviceVertexAttributeDivisorFeatures, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceVertexAttributeDivisorFeatures)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceVertexAttributeDivisorFeatures)))
	*p = C.VkPhysicalDeviceVertexAttributeDivisorFeatures{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkBool32(0)
	if s.VertexAttributeInstanceRateDivisor {
		val0 = C.VkBool32(1)
	}
	p.vertexAttributeInstanceRateDivisor = val0
	val1 := C.VkBool32(0)
	if s.VertexAttributeInstanceRateZeroDivisor {
		val1 = C.VkBool32(1)
	}
	p.vertexAttributeInstanceRateZeroDivisor = val1
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PhysicalDeviceVertexAttributeDivisorProperties struct {
	Next                         Structure
	MaxVertexAttribDivisor       uint32
	SupportsNonZeroFirstInstance bool
}

func (s *PhysicalDeviceVertexAttributeDivisorProperties) GetType() StructureType {
	return StructureTypePhysicalDeviceVertexAttributeDivisorProperties
}

func (s *PhysicalDeviceVertexAttributeDivisorProperties) toC() (*C.VkPhysicalDeviceVertexAttributeDivisorProperties, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPhysicalDeviceVertexAttributeDivisorProperties)(C.malloc(C.size_t(C.sizeof_VkPhysicalDeviceVertexAttributeDivisorProperties)))
	*p = C.VkPhysicalDeviceVertexAttributeDivisorProperties{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.uint32_t(s.MaxVertexAttribDivisor)
	p.maxVertexAttribDivisor = val0
	val1 := C.VkBool32(0)
	if s.SupportsNonZeroFirstInstance {
		val1 = C.VkBool32(1)
	}
	p.supportsNonZeroFirstInstance = val1
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineColorBlendAttachmentState struct {
	BlendEnable         bool
	SrcColorBlendFactor BlendFactor
	DstColorBlendFactor BlendFactor
	ColorBlendOp        BlendOp
	SrcAlphaBlendFactor BlendFactor
	DstAlphaBlendFactor BlendFactor
	AlphaBlendOp        BlendOp
	ColorWriteMask      ColorComponentFlags
}

func (s *PipelineColorBlendAttachmentState) toC() (*C.VkPipelineColorBlendAttachmentState, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineColorBlendAttachmentState)(C.malloc(C.size_t(C.sizeof_VkPipelineColorBlendAttachmentState)))
	*p = C.VkPipelineColorBlendAttachmentState{}
	val0 := C.VkBool32(0)
	if s.BlendEnable {
		val0 = C.VkBool32(1)
	}
	p.blendEnable = val0
	val1 := C.VkBlendFactor(s.SrcColorBlendFactor)
	p.srcColorBlendFactor = val1
	val2 := C.VkBlendFactor(s.DstColorBlendFactor)
	p.dstColorBlendFactor = val2
	val3 := C.VkBlendOp(s.ColorBlendOp)
	p.colorBlendOp = val3
	val4 := C.VkBlendFactor(s.SrcAlphaBlendFactor)
	p.srcAlphaBlendFactor = val4
	val5 := C.VkBlendFactor(s.DstAlphaBlendFactor)
	p.dstAlphaBlendFactor = val5
	val6 := C.VkBlendOp(s.AlphaBlendOp)
	p.alphaBlendOp = val6
	val7 := C.VkColorComponentFlags(s.ColorWriteMask)
	p.colorWriteMask = val7
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineColorBlendStateCreateInfo struct {
	Next           Structure
	Flags          PipelineColorBlendStateCreateFlags
	LogicOpEnable  bool
	LogicOp        LogicOp
	Attachments    []PipelineColorBlendAttachmentState
	BlendConstants [4]float32
}

func (s *PipelineColorBlendStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineColorBlendStateCreateInfo
}

func (s *PipelineColorBlendStateCreateInfo) toC() (*C.VkPipelineColorBlendStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineColorBlendStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineColorBlendStateCreateInfo)))
	*p = C.VkPipelineColorBlendStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPipelineColorBlendStateCreateFlags(s.Flags)
	p.flags = val0
	val1 := C.VkBool32(0)
	if s.LogicOpEnable {
		val1 = C.VkBool32(1)
	}
	p.logicOpEnable = val1
	val2 := C.VkLogicOp(s.LogicOp)
	p.logicOp = val2
	len3 := len(s.Attachments)

	var arr4 *C.VkPipelineColorBlendAttachmentState
	if len3 > 0 {
		arr4 = (*C.VkPipelineColorBlendAttachmentState)(C.malloc(C.size_t(len3) * C.size_t(unsafe.Sizeof(*new(C.VkPipelineColorBlendAttachmentState)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr4)) })
	}
	for i5, elem6 := range s.Attachments {
		val7, cancel8 := elem6.toC()
		cancels = append(cancels, cancel8)
		(*[1 << 30]C.VkPipelineColorBlendAttachmentState)(unsafe.Pointer(arr4))[i5] = *val7
	}
	p.pAttachments = arr4
	p.attachmentCount = C.uint32_t(len(s.Attachments))
	var arr9 [4]C.float
	for i10, elem11 := range s.BlendConstants {
		val12 := C.float(elem11)
		arr9[i10] = val12
	}
	p.blendConstants = arr9
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineDepthStencilStateCreateInfo struct {
	Next                  Structure
	Flags                 PipelineDepthStencilStateCreateFlags
	DepthTestEnable       bool
	DepthWriteEnable      bool
	DepthCompareOp        CompareOp
	DepthBoundsTestEnable bool
	StencilTestEnable     bool
	Front                 StencilOpState
	Back                  StencilOpState
	MinDepthBounds        float32
	MaxDepthBounds        float32
}

func (s *PipelineDepthStencilStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineDepthStencilStateCreateInfo
}

func (s *PipelineDepthStencilStateCreateInfo) toC() (*C.VkPipelineDepthStencilStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineDepthStencilStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineDepthStencilStateCreateInfo)))
	*p = C.VkPipelineDepthStencilStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPipelineDepthStencilStateCreateFlags(s.Flags)
	p.flags = val0
	val1 := C.VkBool32(0)
	if s.DepthTestEnable {
		val1 = C.VkBool32(1)
	}
	p.depthTestEnable = val1
	val2 := C.VkBool32(0)
	if s.DepthWriteEnable {
		val2 = C.VkBool32(1)
	}
	p.depthWriteEnable = val2
	val3 := C.VkCompareOp(s.DepthCompareOp)
	p.depthCompareOp = val3
	val4 := C.VkBool32(0)
	if s.DepthBoundsTestEnable {
		val4 = C.VkBool32(1)
	}
	p.depthBoundsTestEnable = val4
	val5 := C.VkBool32(0)
	if s.StencilTestEnable {
		val5 = C.VkBool32(1)
	}
	p.stencilTestEnable = val5
	val6, cancel7 := s.Front.toC()
	cancels = append(cancels, cancel7)
	p.front = *val6
	val8, cancel9 := s.Back.toC()
	cancels = append(cancels, cancel9)
	p.back = *val8
	val10 := C.float(s.MinDepthBounds)
	p.minDepthBounds = val10
	val11 := C.float(s.MaxDepthBounds)
	p.maxDepthBounds = val11
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineDynamicStateCreateInfo struct {
	Next          Structure
	Flags         PipelineDynamicStateCreateFlags
	DynamicStates []DynamicState
}

func (s *PipelineDynamicStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineDynamicStateCreateInfo
}

func (s *PipelineDynamicStateCreateInfo) toC() (*C.VkPipelineDynamicStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineDynamicStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineDynamicStateCreateInfo)))
	*p = C.VkPipelineDynamicStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPipelineDynamicStateCreateFlags(s.Flags)
	p.flags = val0
	len1 := len(s.DynamicStates)

	var arr2 *C.VkDynamicState
	if len1 > 0 {
		arr2 = (*C.VkDynamicState)(C.malloc(C.size_t(len1) * C.size_t(unsafe.Sizeof(*new(C.VkDynamicState)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr2)) })
	}
	for i3, elem4 := range s.DynamicStates {
		val5 := C.VkDynamicState(elem4)
		(*[1 << 30]C.VkDynamicState)(unsafe.Pointer(arr2))[i3] = val5
	}
	p.pDynamicStates = arr2
	p.dynamicStateCount = C.uint32_t(len(s.DynamicStates))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineInputAssemblyStateCreateInfo struct {
	Next                   Structure
	Flags                  PipelineInputAssemblyStateCreateFlags
	Topology               PrimitiveTopology
	PrimitiveRestartEnable bool
}

func (s *PipelineInputAssemblyStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineInputAssemblyStateCreateInfo
}

func (s *PipelineInputAssemblyStateCreateInfo) toC() (*C.VkPipelineInputAssemblyStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineInputAssemblyStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineInputAssemblyStateCreateInfo)))
	*p = C.VkPipelineInputAssemblyStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPipelineInputAssemblyStateCreateFlags(s.Flags)
	p.flags = val0
	val1 := C.VkPrimitiveTopology(s.Topology)
	p.topology = val1
	val2 := C.VkBool32(0)
	if s.PrimitiveRestartEnable {
		val2 = C.VkBool32(1)
	}
	p.primitiveRestartEnable = val2
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineMultisampleStateCreateInfo struct {
	Next                  Structure
	Flags                 PipelineMultisampleStateCreateFlags
	RasterizationSamples  SampleCountFlagBits
	SampleShadingEnable   bool
	MinSampleShading      float32
	SampleMask            []uint32
	AlphaToCoverageEnable bool
	AlphaToOneEnable      bool
}

func (s *PipelineMultisampleStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineMultisampleStateCreateInfo
}

func (s *PipelineMultisampleStateCreateInfo) toC() (*C.VkPipelineMultisampleStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineMultisampleStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineMultisampleStateCreateInfo)))
	*p = C.VkPipelineMultisampleStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPipelineMultisampleStateCreateFlags(s.Flags)
	p.flags = val0
	val1 := C.VkSampleCountFlagBits(s.RasterizationSamples)
	p.rasterizationSamples = val1
	val2 := C.VkBool32(0)
	if s.SampleShadingEnable {
		val2 = C.VkBool32(1)
	}
	p.sampleShadingEnable = val2
	val3 := C.float(s.MinSampleShading)
	p.minSampleShading = val3
	len4 := len(s.SampleMask)

	var arr5 *C.VkSampleMask
	if len4 > 0 {
		arr5 = (*C.VkSampleMask)(C.malloc(C.size_t(len4) * C.size_t(unsafe.Sizeof(*new(C.VkSampleMask)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr5)) })
	}
	for i6, elem7 := range s.SampleMask {
		val8 := C.VkSampleMask(elem7)
		(*[1 << 30]C.VkSampleMask)(unsafe.Pointer(arr5))[i6] = val8
	}
	p.pSampleMask = arr5
	val9 := C.VkBool32(0)
	if s.AlphaToCoverageEnable {
		val9 = C.VkBool32(1)
	}
	p.alphaToCoverageEnable = val9
	val10 := C.VkBool32(0)
	if s.AlphaToOneEnable {
		val10 = C.VkBool32(1)
	}
	p.alphaToOneEnable = val10
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineRasterizationLineStateCreateInfo struct {
	Next                  Structure
	LineRasterizationMode LineRasterizationMode
	StippledLineEnable    bool
	LineStippleFactor     uint32
	LineStipplePattern    uint16
}

func (s *PipelineRasterizationLineStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineRasterizationLineStateCreateInfo
}

func (s *PipelineRasterizationLineStateCreateInfo) toC() (*C.VkPipelineRasterizationLineStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineRasterizationLineStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineRasterizationLineStateCreateInfo)))
	*p = C.VkPipelineRasterizationLineStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkLineRasterizationMode(s.LineRasterizationMode)
	p.lineRasterizationMode = val0
	val1 := C.VkBool32(0)
	if s.StippledLineEnable {
		val1 = C.VkBool32(1)
	}
	p.stippledLineEnable = val1
	val2 := C.uint32_t(s.LineStippleFactor)
	p.lineStippleFactor = val2
	val3 := C.uint16_t(s.LineStipplePattern)
	p.lineStipplePattern = val3
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineRasterizationStateCreateInfo struct {
	Next                    Structure
	Flags                   PipelineRasterizationStateCreateFlags
	DepthClampEnable        bool
	RasterizerDiscardEnable bool
	PolygonMode             PolygonMode
	CullMode                CullModeFlags
	FrontFace               FrontFace
	DepthBiasEnable         bool
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
	LineWidth               float32
}

func (s *PipelineRasterizationStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineRasterizationStateCreateInfo
}

func (s *PipelineRasterizationStateCreateInfo) toC() (*C.VkPipelineRasterizationStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineRasterizationStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineRasterizationStateCreateInfo)))
	*p = C.VkPipelineRasterizationStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPipelineRasterizationStateCreateFlags(s.Flags)
	p.flags = val0
	val1 := C.VkBool32(0)
	if s.DepthClampEnable {
		val1 = C.VkBool32(1)
	}
	p.depthClampEnable = val1
	val2 := C.VkBool32(0)
	if s.RasterizerDiscardEnable {
		val2 = C.VkBool32(1)
	}
	p.rasterizerDiscardEnable = val2
	val3 := C.VkPolygonMode(s.PolygonMode)
	p.polygonMode = val3
	val4 := C.VkCullModeFlags(s.CullMode)
	p.cullMode = val4
	val5 := C.VkFrontFace(s.FrontFace)
	p.frontFace = val5
	val6 := C.VkBool32(0)
	if s.DepthBiasEnable {
		val6 = C.VkBool32(1)
	}
	p.depthBiasEnable = val6
	val7 := C.float(s.DepthBiasConstantFactor)
	p.depthBiasConstantFactor = val7
	val8 := C.float(s.DepthBiasClamp)
	p.depthBiasClamp = val8
	val9 := C.float(s.DepthBiasSlopeFactor)
	p.depthBiasSlopeFactor = val9
	val10 := C.float(s.LineWidth)
	p.lineWidth = val10
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineRenderingCreateInfo struct {
	Next                    Structure
	ViewMask                uint32
	ColorAttachmentFormats  []Format
	DepthAttachmentFormat   Format
	StencilAttachmentFormat Format
}

func (s *PipelineRenderingCreateInfo) GetType() StructureType {
	return StructureTypePipelineRenderingCreateInfo
}

func (s *PipelineRenderingCreateInfo) toC() (*C.VkPipelineRenderingCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineRenderingCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineRenderingCreateInfo)))
	*p = C.VkPipelineRenderingCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.uint32_t(s.ViewMask)
	p.viewMask = val0
	len1 := len(s.ColorAttachmentFormats)

	var arr2 *C.VkFormat
	if len1 > 0 {
		arr2 = (*C.VkFormat)(C.malloc(C.size_t(len1) * C.size_t(unsafe.Sizeof(*new(C.VkFormat)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr2)) })
	}
	for i3, elem4 := range s.ColorAttachmentFormats {
		val5 := C.VkFormat(elem4)
		(*[1 << 30]C.VkFormat)(unsafe.Pointer(arr2))[i3] = val5
	}
	p.pColorAttachmentFormats = arr2
	p.colorAttachmentCount = C.uint32_t(len(s.ColorAttachmentFormats))
	val6 := C.VkFormat(s.DepthAttachmentFormat)
	p.depthAttachmentFormat = val6
	val7 := C.VkFormat(s.StencilAttachmentFormat)
	p.stencilAttachmentFormat = val7
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineShaderStageCreateInfo struct {
	Next               Structure
	Flags              PipelineShaderStageCreateFlags
	Stage              ShaderStageFlagBits
	Module             *ShaderModule
	Name               string
	SpecializationInfo *SpecializationInfo
}

func (s *PipelineShaderStageCreateInfo) GetType() StructureType {
	return StructureTypePipelineShaderStageCreateInfo
}

func (s *PipelineShaderStageCreateInfo) toC() (*C.VkPipelineShaderStageCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineShaderStageCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineShaderStageCreateInfo)))
	*p = C.VkPipelineShaderStageCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPipelineShaderStageCreateFlags(s.Flags)
	p.flags = val0
	val1 := C.VkShaderStageFlagBits(s.Stage)
	p.stage = val1
	h2 := C.VkShaderModule(unsafe.Pointer(s.Module.handle))
	p.module = h2
	cstr3 := C.CString(s.Name)
	cancels = append(cancels, func() { C.free(unsafe.Pointer(cstr3)) })
	p.pName = cstr3
	var ptr4 *C.VkSpecializationInfo
	if s.SpecializationInfo != nil {
		val5, cancel6 := s.SpecializationInfo.toC()
		cancels = append(cancels, cancel6)
		ptr4 = val5
	}
	p.pSpecializationInfo = ptr4
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineTessellationDomainOriginStateCreateInfo struct {
	Next         Structure
	DomainOrigin TessellationDomainOrigin
}

func (s *PipelineTessellationDomainOriginStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineTessellationDomainOriginStateCreateInfo
}

func (s *PipelineTessellationDomainOriginStateCreateInfo) toC() (*C.VkPipelineTessellationDomainOriginStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineTessellationDomainOriginStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineTessellationDomainOriginStateCreateInfo)))
	*p = C.VkPipelineTessellationDomainOriginStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkTessellationDomainOrigin(s.DomainOrigin)
	p.domainOrigin = val0
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineTessellationStateCreateInfo struct {
	Next               Structure
	Flags              PipelineTessellationStateCreateFlags
	PatchControlPoints uint32
}

func (s *PipelineTessellationStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineTessellationStateCreateInfo
}

func (s *PipelineTessellationStateCreateInfo) toC() (*C.VkPipelineTessellationStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineTessellationStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineTessellationStateCreateInfo)))
	*p = C.VkPipelineTessellationStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPipelineTessellationStateCreateFlags(s.Flags)
	p.flags = val0
	val1 := C.uint32_t(s.PatchControlPoints)
	p.patchControlPoints = val1
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineVertexInputDivisorStateCreateInfo struct {
	Next                  Structure
	VertexBindingDivisors []VertexInputBindingDivisorDescription
}

func (s *PipelineVertexInputDivisorStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineVertexInputDivisorStateCreateInfo
}

func (s *PipelineVertexInputDivisorStateCreateInfo) toC() (*C.VkPipelineVertexInputDivisorStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineVertexInputDivisorStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineVertexInputDivisorStateCreateInfo)))
	*p = C.VkPipelineVertexInputDivisorStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	len0 := len(s.VertexBindingDivisors)

	var arr1 *C.VkVertexInputBindingDivisorDescription
	if len0 > 0 {
		arr1 = (*C.VkVertexInputBindingDivisorDescription)(C.malloc(C.size_t(len0) * C.size_t(unsafe.Sizeof(*new(C.VkVertexInputBindingDivisorDescription)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr1)) })
	}
	for i2, elem3 := range s.VertexBindingDivisors {
		val4, cancel5 := elem3.toC()
		cancels = append(cancels, cancel5)
		(*[1 << 30]C.VkVertexInputBindingDivisorDescription)(unsafe.Pointer(arr1))[i2] = *val4
	}
	p.pVertexBindingDivisors = arr1
	p.vertexBindingDivisorCount = C.uint32_t(len(s.VertexBindingDivisors))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineVertexInputStateCreateInfo struct {
	Next                        Structure
	Flags                       PipelineVertexInputStateCreateFlags
	VertexBindingDescriptions   []VertexInputBindingDescription
	VertexAttributeDescriptions []VertexInputAttributeDescription
}

func (s *PipelineVertexInputStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineVertexInputStateCreateInfo
}

func (s *PipelineVertexInputStateCreateInfo) toC() (*C.VkPipelineVertexInputStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineVertexInputStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineVertexInputStateCreateInfo)))
	*p = C.VkPipelineVertexInputStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPipelineVertexInputStateCreateFlags(s.Flags)
	p.flags = val0
	len1 := len(s.VertexBindingDescriptions)

	var arr2 *C.VkVertexInputBindingDescription
	if len1 > 0 {
		arr2 = (*C.VkVertexInputBindingDescription)(C.malloc(C.size_t(len1) * C.size_t(unsafe.Sizeof(*new(C.VkVertexInputBindingDescription)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr2)) })
	}
	for i3, elem4 := range s.VertexBindingDescriptions {
		val5, cancel6 := elem4.toC()
		cancels = append(cancels, cancel6)
		(*[1 << 30]C.VkVertexInputBindingDescription)(unsafe.Pointer(arr2))[i3] = *val5
	}
	p.pVertexBindingDescriptions = arr2
	p.vertexBindingDescriptionCount = C.uint32_t(len(s.VertexBindingDescriptions))
	len7 := len(s.VertexAttributeDescriptions)

	var arr8 *C.VkVertexInputAttributeDescription
	if len7 > 0 {
		arr8 = (*C.VkVertexInputAttributeDescription)(C.malloc(C.size_t(len7) * C.size_t(unsafe.Sizeof(*new(C.VkVertexInputAttributeDescription)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr8)) })
	}
	for i9, elem10 := range s.VertexAttributeDescriptions {
		val11, cancel12 := elem10.toC()
		cancels = append(cancels, cancel12)
		(*[1 << 30]C.VkVertexInputAttributeDescription)(unsafe.Pointer(arr8))[i9] = *val11
	}
	p.pVertexAttributeDescriptions = arr8
	p.vertexAttributeDescriptionCount = C.uint32_t(len(s.VertexAttributeDescriptions))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PipelineViewportStateCreateInfo struct {
	Next      Structure
	Flags     PipelineViewportStateCreateFlags
	Viewports []Viewport
	Scissors  []Rect2D
}

func (s *PipelineViewportStateCreateInfo) GetType() StructureType {
	return StructureTypePipelineViewportStateCreateInfo
}

func (s *PipelineViewportStateCreateInfo) toC() (*C.VkPipelineViewportStateCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPipelineViewportStateCreateInfo)(C.malloc(C.size_t(C.sizeof_VkPipelineViewportStateCreateInfo)))
	*p = C.VkPipelineViewportStateCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkPipelineViewportStateCreateFlags(s.Flags)
	p.flags = val0
	len1 := len(s.Viewports)

	var arr2 *C.VkViewport
	if len1 > 0 {
		arr2 = (*C.VkViewport)(C.malloc(C.size_t(len1) * C.size_t(unsafe.Sizeof(*new(C.VkViewport)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr2)) })
	}
	for i3, elem4 := range s.Viewports {
		val5, cancel6 := elem4.toC()
		cancels = append(cancels, cancel6)
		(*[1 << 30]C.VkViewport)(unsafe.Pointer(arr2))[i3] = *val5
	}
	p.pViewports = arr2
	p.viewportCount = C.uint32_t(len(s.Viewports))
	len7 := len(s.Scissors)

	var arr8 *C.VkRect2D
	if len7 > 0 {
		arr8 = (*C.VkRect2D)(C.malloc(C.size_t(len7) * C.size_t(unsafe.Sizeof(*new(C.VkRect2D)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr8)) })
	}
	for i9, elem10 := range s.Scissors {
		val11, cancel12 := elem10.toC()
		cancels = append(cancels, cancel12)
		(*[1 << 30]C.VkRect2D)(unsafe.Pointer(arr8))[i9] = *val11
	}
	p.pScissors = arr8
	p.scissorCount = C.uint32_t(len(s.Scissors))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type Rect2D struct {
	Offset Offset2D
	Extent Extent2D
}

func (s *Rect2D) toC() (*C.VkRect2D, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRect2D)(C.malloc(C.size_t(C.sizeof_VkRect2D)))
	*p = C.VkRect2D{}
	val0, cancel1 := s.Offset.toC()
	cancels = append(cancels, cancel1)
	p.offset = *val0
	val2, cancel3 := s.Extent.toC()
	cancels = append(cancels, cancel3)
	p.extent = *val2
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type RenderPassAttachmentBeginInfo struct {
	Next        Structure
	Attachments []*ImageView
}

func (s *RenderPassAttachmentBeginInfo) GetType() StructureType {
	return StructureTypeRenderPassAttachmentBeginInfo
}

func (s *RenderPassAttachmentBeginInfo) toC() (*C.VkRenderPassAttachmentBeginInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRenderPassAttachmentBeginInfo)(C.malloc(C.size_t(C.sizeof_VkRenderPassAttachmentBeginInfo)))
	*p = C.VkRenderPassAttachmentBeginInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	len0 := len(s.Attachments)

	var arr1 *C.VkImageView
	if len0 > 0 {
		arr1 = (*C.VkImageView)(C.malloc(C.size_t(len0) * C.size_t(unsafe.Sizeof(*new(C.VkImageView)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr1)) })
	}
	for i2, elem3 := range s.Attachments {
		h4 := C.VkImageView(unsafe.Pointer(elem3.handle))
		(*[1 << 30]C.VkImageView)(unsafe.Pointer(arr1))[i2] = h4
	}
	p.pAttachments = arr1
	p.attachmentCount = C.uint32_t(len(s.Attachments))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type RenderPassBeginInfo struct {
	Next        Structure
	RenderPass  *RenderPass
	Framebuffer *Framebuffer
	RenderArea  Rect2D
	ClearValues []ClearValue
}

func (s *RenderPassBeginInfo) GetType() StructureType {
	return StructureTypeRenderPassBeginInfo
}

func (s *RenderPassBeginInfo) toC() (*C.VkRenderPassBeginInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRenderPassBeginInfo)(C.malloc(C.size_t(C.sizeof_VkRenderPassBeginInfo)))
	*p = C.VkRenderPassBeginInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	h0 := C.VkRenderPass(unsafe.Pointer(s.RenderPass.handle))
	p.renderPass = h0
	h1 := C.VkFramebuffer(unsafe.Pointer(s.Framebuffer.handle))
	p.framebuffer = h1
	val2, cancel3 := s.RenderArea.toC()
	cancels = append(cancels, cancel3)
	p.renderArea = *val2
	len4 := len(s.ClearValues)

	var arr5 *C.VkClearValue
	if len4 > 0 {
		arr5 = (*C.VkClearValue)(C.malloc(C.size_t(len4) * C.size_t(unsafe.Sizeof(*new(C.VkClearValue)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr5)) })
	}
	for i6, elem7 := range s.ClearValues {
		val8, cancel9 := elem7.toC()
		cancels = append(cancels, cancel9)
		(*[1 << 30]C.VkClearValue)(unsafe.Pointer(arr5))[i6] = *val8
	}
	p.pClearValues = arr5
	p.clearValueCount = C.uint32_t(len(s.ClearValues))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type RenderPassCreateInfo struct {
	Next         Structure
	Flags        RenderPassCreateFlags
	Attachments  []AttachmentDescription
	Subpasses    []SubpassDescription
	Dependencies []SubpassDependency
}

func (s *RenderPassCreateInfo) GetType() StructureType {
	return StructureTypeRenderPassCreateInfo
}

func (s *RenderPassCreateInfo) toC() (*C.VkRenderPassCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRenderPassCreateInfo)(C.malloc(C.size_t(C.sizeof_VkRenderPassCreateInfo)))
	*p = C.VkRenderPassCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkRenderPassCreateFlags(s.Flags)
	p.flags = val0
	len1 := len(s.Attachments)

	var arr2 *C.VkAttachmentDescription
	if len1 > 0 {
		arr2 = (*C.VkAttachmentDescription)(C.malloc(C.size_t(len1) * C.size_t(unsafe.Sizeof(*new(C.VkAttachmentDescription)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr2)) })
	}
	for i3, elem4 := range s.Attachments {
		val5, cancel6 := elem4.toC()
		cancels = append(cancels, cancel6)
		(*[1 << 30]C.VkAttachmentDescription)(unsafe.Pointer(arr2))[i3] = *val5
	}
	p.pAttachments = arr2
	p.attachmentCount = C.uint32_t(len(s.Attachments))
	len7 := len(s.Subpasses)

	var arr8 *C.VkSubpassDescription
	if len7 > 0 {
		arr8 = (*C.VkSubpassDescription)(C.malloc(C.size_t(len7) * C.size_t(unsafe.Sizeof(*new(C.VkSubpassDescription)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr8)) })
	}
	for i9, elem10 := range s.Subpasses {
		val11, cancel12 := elem10.toC()
		cancels = append(cancels, cancel12)
		(*[1 << 30]C.VkSubpassDescription)(unsafe.Pointer(arr8))[i9] = *val11
	}
	p.pSubpasses = arr8
	p.subpassCount = C.uint32_t(len(s.Subpasses))
	len13 := len(s.Dependencies)

	var arr14 *C.VkSubpassDependency
	if len13 > 0 {
		arr14 = (*C.VkSubpassDependency)(C.malloc(C.size_t(len13) * C.size_t(unsafe.Sizeof(*new(C.VkSubpassDependency)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr14)) })
	}
	for i15, elem16 := range s.Dependencies {
		val17, cancel18 := elem16.toC()
		cancels = append(cancels, cancel18)
		(*[1 << 30]C.VkSubpassDependency)(unsafe.Pointer(arr14))[i15] = *val17
	}
	p.pDependencies = arr14
	p.dependencyCount = C.uint32_t(len(s.Dependencies))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type RenderPassCreateInfo2 struct {
	Next                Structure
	Flags               RenderPassCreateFlags
	Attachments         []AttachmentDescription2
	Subpasses           []SubpassDescription2
	Dependencies        []SubpassDependency2
	CorrelatedViewMasks []uint32
}

func (s *RenderPassCreateInfo2) GetType() StructureType {
	return StructureTypeRenderPassCreateInfo2
}

func (s *RenderPassCreateInfo2) toC() (*C.VkRenderPassCreateInfo2, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRenderPassCreateInfo2)(C.malloc(C.size_t(C.sizeof_VkRenderPassCreateInfo2)))
	*p = C.VkRenderPassCreateInfo2{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkRenderPassCreateFlags(s.Flags)
	p.flags = val0
	len1 := len(s.Attachments)

	var arr2 *C.VkAttachmentDescription2
	if len1 > 0 {
		arr2 = (*C.VkAttachmentDescription2)(C.malloc(C.size_t(len1) * C.size_t(unsafe.Sizeof(*new(C.VkAttachmentDescription2)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr2)) })
	}
	for i3, elem4 := range s.Attachments {
		val5, cancel6 := elem4.toC()
		cancels = append(cancels, cancel6)
		(*[1 << 30]C.VkAttachmentDescription2)(unsafe.Pointer(arr2))[i3] = *val5
	}
	p.pAttachments = arr2
	p.attachmentCount = C.uint32_t(len(s.Attachments))
	len7 := len(s.Subpasses)

	var arr8 *C.VkSubpassDescription2
	if len7 > 0 {
		arr8 = (*C.VkSubpassDescription2)(C.malloc(C.size_t(len7) * C.size_t(unsafe.Sizeof(*new(C.VkSubpassDescription2)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr8)) })
	}
	for i9, elem10 := range s.Subpasses {
		val11, cancel12 := elem10.toC()
		cancels = append(cancels, cancel12)
		(*[1 << 30]C.VkSubpassDescription2)(unsafe.Pointer(arr8))[i9] = *val11
	}
	p.pSubpasses = arr8
	p.subpassCount = C.uint32_t(len(s.Subpasses))
	len13 := len(s.Dependencies)

	var arr14 *C.VkSubpassDependency2
	if len13 > 0 {
		arr14 = (*C.VkSubpassDependency2)(C.malloc(C.size_t(len13) * C.size_t(unsafe.Sizeof(*new(C.VkSubpassDependency2)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr14)) })
	}
	for i15, elem16 := range s.Dependencies {
		val17, cancel18 := elem16.toC()
		cancels = append(cancels, cancel18)
		(*[1 << 30]C.VkSubpassDependency2)(unsafe.Pointer(arr14))[i15] = *val17
	}
	p.pDependencies = arr14
	p.dependencyCount = C.uint32_t(len(s.Dependencies))
	len19 := len(s.CorrelatedViewMasks)

	var arr20 *C.uint32_t
	if len19 > 0 {
		arr20 = (*C.uint32_t)(C.malloc(C.size_t(len19) * C.size_t(unsafe.Sizeof(*new(C.uint32_t)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr20)) })
	}
	for i21, elem22 := range s.CorrelatedViewMasks {
		val23 := C.uint32_t(elem22)
		(*[1 << 30]C.uint32_t)(unsafe.Pointer(arr20))[i21] = val23
	}
	p.pCorrelatedViewMasks = arr20
	p.correlatedViewMaskCount = C.uint32_t(len(s.CorrelatedViewMasks))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type RenderPassInputAttachmentAspectCreateInfo struct {
	Next             Structure
	AspectReferences []InputAttachmentAspectReference
}

func (s *RenderPassInputAttachmentAspectCreateInfo) GetType() StructureType {
	return StructureTypeRenderPassInputAttachmentAspectCreateInfo
}

func (s *RenderPassInputAttachmentAspectCreateInfo) toC() (*C.VkRenderPassInputAttachmentAspectCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRenderPassInputAttachmentAspectCreateInfo)(C.malloc(C.size_t(C.sizeof_VkRenderPassInputAttachmentAspectCreateInfo)))
	*p = C.VkRenderPassInputAttachmentAspectCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	len0 := len(s.AspectReferences)

	var arr1 *C.VkInputAttachmentAspectReference
	if len0 > 0 {
		arr1 = (*C.VkInputAttachmentAspectReference)(C.malloc(C.size_t(len0) * C.size_t(unsafe.Sizeof(*new(C.VkInputAttachmentAspectReference)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr1)) })
	}
	for i2, elem3 := range s.AspectReferences {
		val4, cancel5 := elem3.toC()
		cancels = append(cancels, cancel5)
		(*[1 << 30]C.VkInputAttachmentAspectReference)(unsafe.Pointer(arr1))[i2] = *val4
	}
	p.pAspectReferences = arr1
	p.aspectReferenceCount = C.uint32_t(len(s.AspectReferences))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type RenderPassMultiviewCreateInfo struct {
	Next             Structure
	ViewMasks        []uint32
	ViewOffsets      []int32
	CorrelationMasks []uint32
}

func (s *RenderPassMultiviewCreateInfo) GetType() StructureType {
	return StructureTypeRenderPassMultiviewCreateInfo
}

func (s *RenderPassMultiviewCreateInfo) toC() (*C.VkRenderPassMultiviewCreateInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRenderPassMultiviewCreateInfo)(C.malloc(C.size_t(C.sizeof_VkRenderPassMultiviewCreateInfo)))
	*p = C.VkRenderPassMultiviewCreateInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	len0 := len(s.ViewMasks)

	var arr1 *C.uint32_t
	if len0 > 0 {
		arr1 = (*C.uint32_t)(C.malloc(C.size_t(len0) * C.size_t(unsafe.Sizeof(*new(C.uint32_t)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr1)) })
	}
	for i2, elem3 := range s.ViewMasks {
		val4 := C.uint32_t(elem3)
		(*[1 << 30]C.uint32_t)(unsafe.Pointer(arr1))[i2] = val4
	}
	p.pViewMasks = arr1
	p.subpassCount = C.uint32_t(len(s.ViewMasks))
	len5 := len(s.ViewOffsets)

	var arr6 *C.int32_t
	if len5 > 0 {
		arr6 = (*C.int32_t)(C.malloc(C.size_t(len5) * C.size_t(unsafe.Sizeof(*new(C.int32_t)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr6)) })
	}
	for i7, elem8 := range s.ViewOffsets {
		val9 := C.int32_t(elem8)
		(*[1 << 30]C.int32_t)(unsafe.Pointer(arr6))[i7] = val9
	}
	p.pViewOffsets = arr6
	p.dependencyCount = C.uint32_t(len(s.ViewOffsets))
	len10 := len(s.CorrelationMasks)

	var arr11 *C.uint32_t
	if len10 > 0 {
		arr11 = (*C.uint32_t)(C.malloc(C.size_t(len10) * C.size_t(unsafe.Sizeof(*new(C.uint32_t)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr11)) })
	}
	for i12, elem13 := range s.CorrelationMasks {
		val14 := C.uint32_t(elem13)
		(*[1 << 30]C.uint32_t)(unsafe.Pointer(arr11))[i12] = val14
	}
	p.pCorrelationMasks = arr11
	p.correlationMaskCount = C.uint32_t(len(s.CorrelationMasks))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type RenderingAreaInfo struct {
	Next                    Structure
	ViewMask                uint32
	ColorAttachmentFormats  []Format
	DepthAttachmentFormat   Format
	StencilAttachmentFormat Format
}

func (s *RenderingAreaInfo) GetType() StructureType {
	return StructureTypeRenderingAreaInfo
}

func (s *RenderingAreaInfo) toC() (*C.VkRenderingAreaInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRenderingAreaInfo)(C.malloc(C.size_t(C.sizeof_VkRenderingAreaInfo)))
	*p = C.VkRenderingAreaInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.uint32_t(s.ViewMask)
	p.viewMask = val0
	len1 := len(s.ColorAttachmentFormats)

	var arr2 *C.VkFormat
	if len1 > 0 {
		arr2 = (*C.VkFormat)(C.malloc(C.size_t(len1) * C.size_t(unsafe.Sizeof(*new(C.VkFormat)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr2)) })
	}
	for i3, elem4 := range s.ColorAttachmentFormats {
		val5 := C.VkFormat(elem4)
		(*[1 << 30]C.VkFormat)(unsafe.Pointer(arr2))[i3] = val5
	}
	p.pColorAttachmentFormats = arr2
	p.colorAttachmentCount = C.uint32_t(len(s.ColorAttachmentFormats))
	val6 := C.VkFormat(s.DepthAttachmentFormat)
	p.depthAttachmentFormat = val6
	val7 := C.VkFormat(s.StencilAttachmentFormat)
	p.stencilAttachmentFormat = val7
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type RenderingAttachmentInfo struct {
	Next               Structure
	ImageView          *ImageView
	ImageLayout        ImageLayout
	ResolveMode        ResolveModeFlagBits
	ResolveImageView   *ImageView
	ResolveImageLayout ImageLayout
	LoadOp             AttachmentLoadOp
	StoreOp            AttachmentStoreOp
	ClearValue         ClearValue
}

func (s *RenderingAttachmentInfo) GetType() StructureType {
	return StructureTypeRenderingAttachmentInfo
}

func (s *RenderingAttachmentInfo) toC() (*C.VkRenderingAttachmentInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRenderingAttachmentInfo)(C.malloc(C.size_t(C.sizeof_VkRenderingAttachmentInfo)))
	*p = C.VkRenderingAttachmentInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	h0 := C.VkImageView(unsafe.Pointer(s.ImageView.handle))
	p.imageView = h0
	val1 := C.VkImageLayout(s.ImageLayout)
	p.imageLayout = val1
	val2 := C.VkResolveModeFlagBits(s.ResolveMode)
	p.resolveMode = val2
	h3 := C.VkImageView(unsafe.Pointer(s.ResolveImageView.handle))
	p.resolveImageView = h3
	val4 := C.VkImageLayout(s.ResolveImageLayout)
	p.resolveImageLayout = val4
	val5 := C.VkAttachmentLoadOp(s.LoadOp)
	p.loadOp = val5
	val6 := C.VkAttachmentStoreOp(s.StoreOp)
	p.storeOp = val6
	val7, cancel8 := s.ClearValue.toC()
	cancels = append(cancels, cancel8)
	p.clearValue = *val7
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type RenderingAttachmentLocationInfo struct {
	Next                     Structure
	ColorAttachmentLocations []uint32
}

func (s *RenderingAttachmentLocationInfo) GetType() StructureType {
	return StructureTypeRenderingAttachmentLocationInfo
}

func (s *RenderingAttachmentLocationInfo) toC() (*C.VkRenderingAttachmentLocationInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRenderingAttachmentLocationInfo)(C.malloc(C.size_t(C.sizeof_VkRenderingAttachmentLocationInfo)))
	*p = C.VkRenderingAttachmentLocationInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	len0 := len(s.ColorAttachmentLocations)

	var arr1 *C.uint32_t
	if len0 > 0 {
		arr1 = (*C.uint32_t)(C.malloc(C.size_t(len0) * C.size_t(unsafe.Sizeof(*new(C.uint32_t)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr1)) })
	}
	for i2, elem3 := range s.ColorAttachmentLocations {
		val4 := C.uint32_t(elem3)
		(*[1 << 30]C.uint32_t)(unsafe.Pointer(arr1))[i2] = val4
	}
	p.pColorAttachmentLocations = arr1
	p.colorAttachmentCount = C.uint32_t(len(s.ColorAttachmentLocations))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type RenderingInfo struct {
	Next              Structure
	Flags             RenderingFlags
	RenderArea        Rect2D
	LayerCount        uint32
	ViewMask          uint32
	ColorAttachments  []RenderingAttachmentInfo
	DepthAttachment   *RenderingAttachmentInfo
	StencilAttachment *RenderingAttachmentInfo
}

func (s *RenderingInfo) GetType() StructureType {
	return StructureTypeRenderingInfo
}

func (s *RenderingInfo) toC() (*C.VkRenderingInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRenderingInfo)(C.malloc(C.size_t(C.sizeof_VkRenderingInfo)))
	*p = C.VkRenderingInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkRenderingFlags(s.Flags)
	p.flags = val0
	val1, cancel2 := s.RenderArea.toC()
	cancels = append(cancels, cancel2)
	p.renderArea = *val1
	val3 := C.uint32_t(s.LayerCount)
	p.layerCount = val3
	val4 := C.uint32_t(s.ViewMask)
	p.viewMask = val4
	len5 := len(s.ColorAttachments)

	var arr6 *C.VkRenderingAttachmentInfo
	if len5 > 0 {
		arr6 = (*C.VkRenderingAttachmentInfo)(C.malloc(C.size_t(len5) * C.size_t(unsafe.Sizeof(*new(C.VkRenderingAttachmentInfo)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr6)) })
	}
	for i7, elem8 := range s.ColorAttachments {
		val9, cancel10 := elem8.toC()
		cancels = append(cancels, cancel10)
		(*[1 << 30]C.VkRenderingAttachmentInfo)(unsafe.Pointer(arr6))[i7] = *val9
	}
	p.pColorAttachments = arr6
	p.colorAttachmentCount = C.uint32_t(len(s.ColorAttachments))
	var ptr11 *C.VkRenderingAttachmentInfo
	if s.DepthAttachment != nil {
		val12, cancel13 := s.DepthAttachment.toC()
		cancels = append(cancels, cancel13)
		ptr11 = val12
	}
	p.pDepthAttachment = ptr11
	var ptr14 *C.VkRenderingAttachmentInfo
	if s.StencilAttachment != nil {
		val15, cancel16 := s.StencilAttachment.toC()
		cancels = append(cancels, cancel16)
		ptr14 = val15
	}
	p.pStencilAttachment = ptr14
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type RenderingInputAttachmentIndexInfo struct {
	Next                        Structure
	ColorAttachmentInputIndices []uint32
	DepthInputAttachmentIndex   *uint32
	StencilInputAttachmentIndex *uint32
}

func (s *RenderingInputAttachmentIndexInfo) GetType() StructureType {
	return StructureTypeRenderingInputAttachmentIndexInfo
}

func (s *RenderingInputAttachmentIndexInfo) toC() (*C.VkRenderingInputAttachmentIndexInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRenderingInputAttachmentIndexInfo)(C.malloc(C.size_t(C.sizeof_VkRenderingInputAttachmentIndexInfo)))
	*p = C.VkRenderingInputAttachmentIndexInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	len0 := len(s.ColorAttachmentInputIndices)

	var arr1 *C.uint32_t
	if len0 > 0 {
		arr1 = (*C.uint32_t)(C.malloc(C.size_t(len0) * C.size_t(unsafe.Sizeof(*new(C.uint32_t)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr1)) })
	}
	for i2, elem3 := range s.ColorAttachmentInputIndices {
		val4 := C.uint32_t(elem3)
		(*[1 << 30]C.uint32_t)(unsafe.Pointer(arr1))[i2] = val4
	}
	p.pColorAttachmentInputIndices = arr1
	p.colorAttachmentCount = C.uint32_t(len(s.ColorAttachmentInputIndices))
	var ptr5 *C.uint32_t
	if s.DepthInputAttachmentIndex != nil {
		val6 := C.uint32_t(*s.DepthInputAttachmentIndex)
		ptr5 = &val6
	}
	p.pDepthInputAttachmentIndex = ptr5
	var ptr7 *C.uint32_t
	if s.StencilInputAttachmentIndex != nil {
		val8 := C.uint32_t(*s.StencilInputAttachmentIndex)
		ptr7 = &val8
	}
	p.pStencilInputAttachmentIndex = ptr7
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type ResolveImageInfo2 struct {
	Next           Structure
	SrcImage       *Image
	SrcImageLayout ImageLayout
	DstImage       *Image
	DstImageLayout ImageLayout
	Regions        []ImageResolve2
}

func (s *ResolveImageInfo2) GetType() StructureType {
	return StructureTypeResolveImageInfo2
}

func (s *ResolveImageInfo2) toC() (*C.VkResolveImageInfo2, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkResolveImageInfo2)(C.malloc(C.size_t(C.sizeof_VkResolveImageInfo2)))
	*p = C.VkResolveImageInfo2{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	h0 := C.VkImage(unsafe.Pointer(s.SrcImage.handle))
	p.srcImage = h0
	val1 := C.VkImageLayout(s.SrcImageLayout)
	p.srcImageLayout = val1
	h2 := C.VkImage(unsafe.Pointer(s.DstImage.handle))
	p.dstImage = h2
	val3 := C.VkImageLayout(s.DstImageLayout)
	p.dstImageLayout = val3
	len4 := len(s.Regions)

	var arr5 *C.VkImageResolve2
	if len4 > 0 {
		arr5 = (*C.VkImageResolve2)(C.malloc(C.size_t(len4) * C.size_t(unsafe.Sizeof(*new(C.VkImageResolve2)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr5)) })
	}
	for i6, elem7 := range s.Regions {
		val8, cancel9 := elem7.toC()
		cancels = append(cancels, cancel9)
		(*[1 << 30]C.VkImageResolve2)(unsafe.Pointer(arr5))[i6] = *val8
	}
	p.pRegions = arr5
	p.regionCount = C.uint32_t(len(s.Regions))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SpecializationInfo struct {
	MapEntries []SpecializationMapEntry
	Data       unsafe.Pointer
}

func (s *SpecializationInfo) toC() (*C.VkSpecializationInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSpecializationInfo)(C.malloc(C.size_t(C.sizeof_VkSpecializationInfo)))
	*p = C.VkSpecializationInfo{}
	len0 := len(s.MapEntries)

	var arr1 *C.VkSpecializationMapEntry
	if len0 > 0 {
		arr1 = (*C.VkSpecializationMapEntry)(C.malloc(C.size_t(len0) * C.size_t(unsafe.Sizeof(*new(C.VkSpecializationMapEntry)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr1)) })
	}
	for i2, elem3 := range s.MapEntries {
		val4, cancel5 := elem3.toC()
		cancels = append(cancels, cancel5)
		(*[1 << 30]C.VkSpecializationMapEntry)(unsafe.Pointer(arr1))[i2] = *val4
	}
	p.pMapEntries = arr1
	p.mapEntryCount = C.uint32_t(len(s.MapEntries))
	p.pData = s.Data
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SpecializationMapEntry struct {
	ConstantID uint32
	Offset     uint32
	Size       uintptr
}

func (s *SpecializationMapEntry) toC() (*C.VkSpecializationMapEntry, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSpecializationMapEntry)(C.malloc(C.size_t(C.sizeof_VkSpecializationMapEntry)))
	*p = C.VkSpecializationMapEntry{}
	val0 := C.uint32_t(s.ConstantID)
	p.constantID = val0
	val1 := C.uint32_t(s.Offset)
	p.offset = val1
	val2 := C.size_t(s.Size)
	p.size = val2
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type StencilOpState struct {
	FailOp      StencilOp
	PassOp      StencilOp
	DepthFailOp StencilOp
	CompareOp   CompareOp
	CompareMask uint32
	WriteMask   uint32
	Reference   uint32
}

func (s *StencilOpState) toC() (*C.VkStencilOpState, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkStencilOpState)(C.malloc(C.size_t(C.sizeof_VkStencilOpState)))
	*p = C.VkStencilOpState{}
	val0 := C.VkStencilOp(s.FailOp)
	p.failOp = val0
	val1 := C.VkStencilOp(s.PassOp)
	p.passOp = val1
	val2 := C.VkStencilOp(s.DepthFailOp)
	p.depthFailOp = val2
	val3 := C.VkCompareOp(s.CompareOp)
	p.compareOp = val3
	val4 := C.uint32_t(s.CompareMask)
	p.compareMask = val4
	val5 := C.uint32_t(s.WriteMask)
	p.writeMask = val5
	val6 := C.uint32_t(s.Reference)
	p.reference = val6
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SubpassBeginInfo struct {
	Next     Structure
	Contents SubpassContents
}

func (s *SubpassBeginInfo) GetType() StructureType {
	return StructureTypeSubpassBeginInfo
}

func (s *SubpassBeginInfo) toC() (*C.VkSubpassBeginInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSubpassBeginInfo)(C.malloc(C.size_t(C.sizeof_VkSubpassBeginInfo)))
	*p = C.VkSubpassBeginInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkSubpassContents(s.Contents)
	p.contents = val0
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SubpassDependency struct {
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    PipelineStageFlags
	DstStageMask    PipelineStageFlags
	SrcAccessMask   AccessFlags
	DstAccessMask   AccessFlags
	DependencyFlags DependencyFlags
}

func (s *SubpassDependency) toC() (*C.VkSubpassDependency, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSubpassDependency)(C.malloc(C.size_t(C.sizeof_VkSubpassDependency)))
	*p = C.VkSubpassDependency{}
	val0 := C.uint32_t(s.SrcSubpass)
	p.srcSubpass = val0
	val1 := C.uint32_t(s.DstSubpass)
	p.dstSubpass = val1
	val2 := C.VkPipelineStageFlags(s.SrcStageMask)
	p.srcStageMask = val2
	val3 := C.VkPipelineStageFlags(s.DstStageMask)
	p.dstStageMask = val3
	val4 := C.VkAccessFlags(s.SrcAccessMask)
	p.srcAccessMask = val4
	val5 := C.VkAccessFlags(s.DstAccessMask)
	p.dstAccessMask = val5
	val6 := C.VkDependencyFlags(s.DependencyFlags)
	p.dependencyFlags = val6
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SubpassDependency2 struct {
	Next            Structure
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    PipelineStageFlags
	DstStageMask    PipelineStageFlags
	SrcAccessMask   AccessFlags
	DstAccessMask   AccessFlags
	DependencyFlags DependencyFlags
	ViewOffset      int32
}

func (s *SubpassDependency2) GetType() StructureType {
	return StructureTypeSubpassDependency2
}

func (s *SubpassDependency2) toC() (*C.VkSubpassDependency2, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSubpassDependency2)(C.malloc(C.size_t(C.sizeof_VkSubpassDependency2)))
	*p = C.VkSubpassDependency2{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.uint32_t(s.SrcSubpass)
	p.srcSubpass = val0
	val1 := C.uint32_t(s.DstSubpass)
	p.dstSubpass = val1
	val2 := C.VkPipelineStageFlags(s.SrcStageMask)
	p.srcStageMask = val2
	val3 := C.VkPipelineStageFlags(s.DstStageMask)
	p.dstStageMask = val3
	val4 := C.VkAccessFlags(s.SrcAccessMask)
	p.srcAccessMask = val4
	val5 := C.VkAccessFlags(s.DstAccessMask)
	p.dstAccessMask = val5
	val6 := C.VkDependencyFlags(s.DependencyFlags)
	p.dependencyFlags = val6
	val7 := C.int32_t(s.ViewOffset)
	p.viewOffset = val7
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SubpassDescription struct {
	Flags                  SubpassDescriptionFlags
	PipelineBindPoint      PipelineBindPoint
	InputAttachments       []AttachmentReference
	ColorAttachments       []AttachmentReference
	ResolveAttachments     []AttachmentReference
	DepthStencilAttachment *AttachmentReference
	PreserveAttachments    []uint32
}

func (s *SubpassDescription) toC() (*C.VkSubpassDescription, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSubpassDescription)(C.malloc(C.size_t(C.sizeof_VkSubpassDescription)))
	*p = C.VkSubpassDescription{}
	val0 := C.VkSubpassDescriptionFlags(s.Flags)
	p.flags = val0
	val1 := C.VkPipelineBindPoint(s.PipelineBindPoint)
	p.pipelineBindPoint = val1
	len2 := len(s.InputAttachments)

	var arr3 *C.VkAttachmentReference
	if len2 > 0 {
		arr3 = (*C.VkAttachmentReference)(C.malloc(C.size_t(len2) * C.size_t(unsafe.Sizeof(*new(C.VkAttachmentReference)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr3)) })
	}
	for i4, elem5 := range s.InputAttachments {
		val6, cancel7 := elem5.toC()
		cancels = append(cancels, cancel7)
		(*[1 << 30]C.VkAttachmentReference)(unsafe.Pointer(arr3))[i4] = *val6
	}
	p.pInputAttachments = arr3
	p.inputAttachmentCount = C.uint32_t(len(s.InputAttachments))
	len8 := len(s.ColorAttachments)

	var arr9 *C.VkAttachmentReference
	if len8 > 0 {
		arr9 = (*C.VkAttachmentReference)(C.malloc(C.size_t(len8) * C.size_t(unsafe.Sizeof(*new(C.VkAttachmentReference)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr9)) })
	}
	for i10, elem11 := range s.ColorAttachments {
		val12, cancel13 := elem11.toC()
		cancels = append(cancels, cancel13)
		(*[1 << 30]C.VkAttachmentReference)(unsafe.Pointer(arr9))[i10] = *val12
	}
	p.pColorAttachments = arr9
	p.colorAttachmentCount = C.uint32_t(len(s.ColorAttachments))
	len14 := len(s.ResolveAttachments)

	var arr15 *C.VkAttachmentReference
	if len14 > 0 {
		arr15 = (*C.VkAttachmentReference)(C.malloc(C.size_t(len14) * C.size_t(unsafe.Sizeof(*new(C.VkAttachmentReference)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr15)) })
	}
	for i16, elem17 := range s.ResolveAttachments {
		val18, cancel19 := elem17.toC()
		cancels = append(cancels, cancel19)
		(*[1 << 30]C.VkAttachmentReference)(unsafe.Pointer(arr15))[i16] = *val18
	}
	p.pResolveAttachments = arr15
	p.colorAttachmentCount = C.uint32_t(len(s.ResolveAttachments))
	var ptr20 *C.VkAttachmentReference
	if s.DepthStencilAttachment != nil {
		val21, cancel22 := s.DepthStencilAttachment.toC()
		cancels = append(cancels, cancel22)
		ptr20 = val21
	}
	p.pDepthStencilAttachment = ptr20
	len23 := len(s.PreserveAttachments)

	var arr24 *C.uint32_t
	if len23 > 0 {
		arr24 = (*C.uint32_t)(C.malloc(C.size_t(len23) * C.size_t(unsafe.Sizeof(*new(C.uint32_t)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr24)) })
	}
	for i25, elem26 := range s.PreserveAttachments {
		val27 := C.uint32_t(elem26)
		(*[1 << 30]C.uint32_t)(unsafe.Pointer(arr24))[i25] = val27
	}
	p.pPreserveAttachments = arr24
	p.preserveAttachmentCount = C.uint32_t(len(s.PreserveAttachments))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SubpassDescription2 struct {
	Next                   Structure
	Flags                  SubpassDescriptionFlags
	PipelineBindPoint      PipelineBindPoint
	ViewMask               uint32
	InputAttachments       []AttachmentReference2
	ColorAttachments       []AttachmentReference2
	ResolveAttachments     []AttachmentReference2
	DepthStencilAttachment *AttachmentReference2
	PreserveAttachments    []uint32
}

func (s *SubpassDescription2) GetType() StructureType {
	return StructureTypeSubpassDescription2
}

func (s *SubpassDescription2) toC() (*C.VkSubpassDescription2, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSubpassDescription2)(C.malloc(C.size_t(C.sizeof_VkSubpassDescription2)))
	*p = C.VkSubpassDescription2{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkSubpassDescriptionFlags(s.Flags)
	p.flags = val0
	val1 := C.VkPipelineBindPoint(s.PipelineBindPoint)
	p.pipelineBindPoint = val1
	val2 := C.uint32_t(s.ViewMask)
	p.viewMask = val2
	len3 := len(s.InputAttachments)

	var arr4 *C.VkAttachmentReference2
	if len3 > 0 {
		arr4 = (*C.VkAttachmentReference2)(C.malloc(C.size_t(len3) * C.size_t(unsafe.Sizeof(*new(C.VkAttachmentReference2)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr4)) })
	}
	for i5, elem6 := range s.InputAttachments {
		val7, cancel8 := elem6.toC()
		cancels = append(cancels, cancel8)
		(*[1 << 30]C.VkAttachmentReference2)(unsafe.Pointer(arr4))[i5] = *val7
	}
	p.pInputAttachments = arr4
	p.inputAttachmentCount = C.uint32_t(len(s.InputAttachments))
	len9 := len(s.ColorAttachments)

	var arr10 *C.VkAttachmentReference2
	if len9 > 0 {
		arr10 = (*C.VkAttachmentReference2)(C.malloc(C.size_t(len9) * C.size_t(unsafe.Sizeof(*new(C.VkAttachmentReference2)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr10)) })
	}
	for i11, elem12 := range s.ColorAttachments {
		val13, cancel14 := elem12.toC()
		cancels = append(cancels, cancel14)
		(*[1 << 30]C.VkAttachmentReference2)(unsafe.Pointer(arr10))[i11] = *val13
	}
	p.pColorAttachments = arr10
	p.colorAttachmentCount = C.uint32_t(len(s.ColorAttachments))
	len15 := len(s.ResolveAttachments)

	var arr16 *C.VkAttachmentReference2
	if len15 > 0 {
		arr16 = (*C.VkAttachmentReference2)(C.malloc(C.size_t(len15) * C.size_t(unsafe.Sizeof(*new(C.VkAttachmentReference2)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr16)) })
	}
	for i17, elem18 := range s.ResolveAttachments {
		val19, cancel20 := elem18.toC()
		cancels = append(cancels, cancel20)
		(*[1 << 30]C.VkAttachmentReference2)(unsafe.Pointer(arr16))[i17] = *val19
	}
	p.pResolveAttachments = arr16
	p.colorAttachmentCount = C.uint32_t(len(s.ResolveAttachments))
	var ptr21 *C.VkAttachmentReference2
	if s.DepthStencilAttachment != nil {
		val22, cancel23 := s.DepthStencilAttachment.toC()
		cancels = append(cancels, cancel23)
		ptr21 = val22
	}
	p.pDepthStencilAttachment = ptr21
	len24 := len(s.PreserveAttachments)

	var arr25 *C.uint32_t
	if len24 > 0 {
		arr25 = (*C.uint32_t)(C.malloc(C.size_t(len24) * C.size_t(unsafe.Sizeof(*new(C.uint32_t)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr25)) })
	}
	for i26, elem27 := range s.PreserveAttachments {
		val28 := C.uint32_t(elem27)
		(*[1 << 30]C.uint32_t)(unsafe.Pointer(arr25))[i26] = val28
	}
	p.pPreserveAttachments = arr25
	p.preserveAttachmentCount = C.uint32_t(len(s.PreserveAttachments))
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SubpassDescriptionDepthStencilResolve struct {
	Next                          Structure
	DepthResolveMode              ResolveModeFlagBits
	StencilResolveMode            ResolveModeFlagBits
	DepthStencilResolveAttachment *AttachmentReference2
}

func (s *SubpassDescriptionDepthStencilResolve) GetType() StructureType {
	return StructureTypeSubpassDescriptionDepthStencilResolve
}

func (s *SubpassDescriptionDepthStencilResolve) toC() (*C.VkSubpassDescriptionDepthStencilResolve, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSubpassDescriptionDepthStencilResolve)(C.malloc(C.size_t(C.sizeof_VkSubpassDescriptionDepthStencilResolve)))
	*p = C.VkSubpassDescriptionDepthStencilResolve{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	val0 := C.VkResolveModeFlagBits(s.DepthResolveMode)
	p.depthResolveMode = val0
	val1 := C.VkResolveModeFlagBits(s.StencilResolveMode)
	p.stencilResolveMode = val1
	var ptr2 *C.VkAttachmentReference2
	if s.DepthStencilResolveAttachment != nil {
		val3, cancel4 := s.DepthStencilResolveAttachment.toC()
		cancels = append(cancels, cancel4)
		ptr2 = val3
	}
	p.pDepthStencilResolveAttachment = ptr2
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SubpassEndInfo struct {
	Next Structure
}

func (s *SubpassEndInfo) GetType() StructureType {
	return StructureTypeSubpassEndInfo
}

func (s *SubpassEndInfo) toC() (*C.VkSubpassEndInfo, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSubpassEndInfo)(C.malloc(C.size_t(C.sizeof_VkSubpassEndInfo)))
	*p = C.VkSubpassEndInfo{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = unsafe.Pointer(nextPtr)
	}
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type VertexInputAttributeDescription struct {
	Location uint32
	Binding  uint32
	Format   Format
	Offset   uint32
}

func (s *VertexInputAttributeDescription) toC() (*C.VkVertexInputAttributeDescription, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkVertexInputAttributeDescription)(C.malloc(C.size_t(C.sizeof_VkVertexInputAttributeDescription)))
	*p = C.VkVertexInputAttributeDescription{}
	val0 := C.uint32_t(s.Location)
	p.location = val0
	val1 := C.uint32_t(s.Binding)
	p.binding = val1
	val2 := C.VkFormat(s.Format)
	p.format = val2
	val3 := C.uint32_t(s.Offset)
	p.offset = val3
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type VertexInputBindingDescription struct {
	Binding   uint32
	Stride    uint32
	InputRate VertexInputRate
}

func (s *VertexInputBindingDescription) toC() (*C.VkVertexInputBindingDescription, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkVertexInputBindingDescription)(C.malloc(C.size_t(C.sizeof_VkVertexInputBindingDescription)))
	*p = C.VkVertexInputBindingDescription{}
	val0 := C.uint32_t(s.Binding)
	p.binding = val0
	val1 := C.uint32_t(s.Stride)
	p.stride = val1
	val2 := C.VkVertexInputRate(s.InputRate)
	p.inputRate = val2
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type VertexInputBindingDivisorDescription struct {
	Binding uint32
	Divisor uint32
}

func (s *VertexInputBindingDivisorDescription) toC() (*C.VkVertexInputBindingDivisorDescription, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkVertexInputBindingDivisorDescription)(C.malloc(C.size_t(C.sizeof_VkVertexInputBindingDivisorDescription)))
	*p = C.VkVertexInputBindingDivisorDescription{}
	val0 := C.uint32_t(s.Binding)
	p.binding = val0
	val1 := C.uint32_t(s.Divisor)
	p.divisor = val1
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type Viewport struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32
}

func (s *Viewport) toC() (*C.VkViewport, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkViewport)(C.malloc(C.size_t(C.sizeof_VkViewport)))
	*p = C.VkViewport{}
	val0 := C.float(s.X)
	p.x = val0
	val1 := C.float(s.Y)
	p.y = val1
	val2 := C.float(s.Width)
	p.width = val2
	val3 := C.float(s.Height)
	p.height = val3
	val4 := C.float(s.MinDepth)
	p.minDepth = val4
	val5 := C.float(s.MaxDepth)
	p.maxDepth = val5
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

func (h CommandBuffer) BeginRenderPass(
	renderPassBegin *RenderPassBeginInfo,
	contents SubpassContents,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param renderPassBegin
	var ptr1 *C.VkRenderPassBeginInfo
	if renderPassBegin != nil {
		val2, cancel3 := renderPassBegin.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	// param contents
	val5 := C.VkSubpassContents(contents)
	C.fn_vkCmdBeginRenderPass(C.VkCommandBuffer(unsafe.Pointer(h.handle)), ptr1, val5)
}

func (h CommandBuffer) BeginRenderPass2(
	renderPassBegin *RenderPassBeginInfo,
	subpassBeginInfo *SubpassBeginInfo,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param renderPassBegin
	var ptr1 *C.VkRenderPassBeginInfo
	if renderPassBegin != nil {
		val2, cancel3 := renderPassBegin.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	// param subpassBeginInfo
	var ptr5 *C.VkSubpassBeginInfo
	if subpassBeginInfo != nil {
		val6, cancel7 := subpassBeginInfo.toC()
		cancels = append(cancels, cancel7)
		ptr5 = val6
	}
	C.fn_vkCmdBeginRenderPass2(C.VkCommandBuffer(unsafe.Pointer(h.handle)), ptr1, ptr5)
}

func (h CommandBuffer) BeginRendering(
	renderingInfo *RenderingInfo,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param renderingInfo
	var ptr1 *C.VkRenderingInfo
	if renderingInfo != nil {
		val2, cancel3 := renderingInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	C.fn_vkCmdBeginRendering(C.VkCommandBuffer(unsafe.Pointer(h.handle)), ptr1)
}

func (h CommandBuffer) BindIndexBuffer(
	buffer *Buffer,
	offset uint64,
	indexType IndexType,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param buffer
	h1 := C.VkBuffer(unsafe.Pointer(buffer.handle))
	// param offset
	val3 := C.VkDeviceSize(offset)
	// param indexType
	val5 := C.VkIndexType(indexType)
	C.fn_vkCmdBindIndexBuffer(C.VkCommandBuffer(unsafe.Pointer(h.handle)), h1, val3, val5)
}

func (h CommandBuffer) BindIndexBuffer2(
	buffer *Buffer,
	offset uint64,
	size uint64,
	indexType IndexType,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param buffer
	h1 := C.VkBuffer(unsafe.Pointer(buffer.handle))
	// param offset
	val3 := C.VkDeviceSize(offset)
	// param size
	val5 := C.VkDeviceSize(size)
	// param indexType
	val7 := C.VkIndexType(indexType)
	C.fn_vkCmdBindIndexBuffer2(C.VkCommandBuffer(unsafe.Pointer(h.handle)), h1, val3, val5, val7)
}

func (h CommandBuffer) BindVertexBuffers(
	firstBinding uint32,
	bindingCount uint32,
	buffers []*Buffer,
	offsets []uint64,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param firstBinding
	val1 := C.uint32_t(firstBinding)
	// param bindingCount
	val3 := C.uint32_t(bindingCount)
	// param buffers
	len5 := len(buffers)

	var arr6 *C.VkBuffer
	if len5 > 0 {
		arr6 = (*C.VkBuffer)(C.malloc(C.size_t(len5) * C.size_t(unsafe.Sizeof(*new(C.VkBuffer)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr6)) })
	}
	for i7, elem8 := range buffers {
		h9 := C.VkBuffer(unsafe.Pointer(elem8.handle))
		(*[1 << 30]C.VkBuffer)(unsafe.Pointer(arr6))[i7] = h9
	}
	// param offsets
	len11 := len(offsets)

	var arr12 *C.VkDeviceSize
	if len11 > 0 {
		arr12 = (*C.VkDeviceSize)(C.malloc(C.size_t(len11) * C.size_t(unsafe.Sizeof(*new(C.VkDeviceSize)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr12)) })
	}
	for i13, elem14 := range offsets {
		val15 := C.VkDeviceSize(elem14)
		(*[1 << 30]C.VkDeviceSize)(unsafe.Pointer(arr12))[i13] = val15
	}
	C.fn_vkCmdBindVertexBuffers(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3, arr6, arr12)
}

func (h CommandBuffer) BindVertexBuffers2(
	firstBinding uint32,
	bindingCount uint32,
	buffers []*Buffer,
	offsets []uint64,
	sizes []uint64,
	strides []uint64,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param firstBinding
	val1 := C.uint32_t(firstBinding)
	// param bindingCount
	val3 := C.uint32_t(bindingCount)
	// param buffers
	len5 := len(buffers)

	var arr6 *C.VkBuffer
	if len5 > 0 {
		arr6 = (*C.VkBuffer)(C.malloc(C.size_t(len5) * C.size_t(unsafe.Sizeof(*new(C.VkBuffer)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr6)) })
	}
	for i7, elem8 := range buffers {
		h9 := C.VkBuffer(unsafe.Pointer(elem8.handle))
		(*[1 << 30]C.VkBuffer)(unsafe.Pointer(arr6))[i7] = h9
	}
	// param offsets
	len11 := len(offsets)

	var arr12 *C.VkDeviceSize
	if len11 > 0 {
		arr12 = (*C.VkDeviceSize)(C.malloc(C.size_t(len11) * C.size_t(unsafe.Sizeof(*new(C.VkDeviceSize)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr12)) })
	}
	for i13, elem14 := range offsets {
		val15 := C.VkDeviceSize(elem14)
		(*[1 << 30]C.VkDeviceSize)(unsafe.Pointer(arr12))[i13] = val15
	}
	// param sizes
	len17 := len(sizes)

	var arr18 *C.VkDeviceSize
	if len17 > 0 {
		arr18 = (*C.VkDeviceSize)(C.malloc(C.size_t(len17) * C.size_t(unsafe.Sizeof(*new(C.VkDeviceSize)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr18)) })
	}
	for i19, elem20 := range sizes {
		val21 := C.VkDeviceSize(elem20)
		(*[1 << 30]C.VkDeviceSize)(unsafe.Pointer(arr18))[i19] = val21
	}
	// param strides
	len23 := len(strides)

	var arr24 *C.VkDeviceSize
	if len23 > 0 {
		arr24 = (*C.VkDeviceSize)(C.malloc(C.size_t(len23) * C.size_t(unsafe.Sizeof(*new(C.VkDeviceSize)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr24)) })
	}
	for i25, elem26 := range strides {
		val27 := C.VkDeviceSize(elem26)
		(*[1 << 30]C.VkDeviceSize)(unsafe.Pointer(arr24))[i25] = val27
	}
	C.fn_vkCmdBindVertexBuffers2(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3, arr6, arr12, arr18, arr24)
}

func (h CommandBuffer) BlitImage(
	srcImage *Image,
	srcImageLayout ImageLayout,
	dstImage *Image,
	dstImageLayout ImageLayout,
	regionCount uint32,
	regions []ImageBlit,
	filter Filter,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param srcImage
	h1 := C.VkImage(unsafe.Pointer(srcImage.handle))
	// param srcImageLayout
	val3 := C.VkImageLayout(srcImageLayout)
	// param dstImage
	h5 := C.VkImage(unsafe.Pointer(dstImage.handle))
	// param dstImageLayout
	val7 := C.VkImageLayout(dstImageLayout)
	// param regionCount
	val9 := C.uint32_t(regionCount)
	// param regions
	len11 := len(regions)

	var arr12 *C.VkImageBlit
	if len11 > 0 {
		arr12 = (*C.VkImageBlit)(C.malloc(C.size_t(len11) * C.size_t(unsafe.Sizeof(*new(C.VkImageBlit)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr12)) })
	}
	for i13, elem14 := range regions {
		val15, cancel16 := elem14.toC()
		cancels = append(cancels, cancel16)
		(*[1 << 30]C.VkImageBlit)(unsafe.Pointer(arr12))[i13] = *val15
	}
	// param filter
	val18 := C.VkFilter(filter)
	C.fn_vkCmdBlitImage(C.VkCommandBuffer(unsafe.Pointer(h.handle)), h1, val3, h5, val7, val9, arr12, val18)
}

func (h CommandBuffer) BlitImage2(
	blitImageInfo *BlitImageInfo2,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param blitImageInfo
	var ptr1 *C.VkBlitImageInfo2
	if blitImageInfo != nil {
		val2, cancel3 := blitImageInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	C.fn_vkCmdBlitImage2(C.VkCommandBuffer(unsafe.Pointer(h.handle)), ptr1)
}

func (h CommandBuffer) ClearAttachments(
	attachmentCount uint32,
	attachments []ClearAttachment,
	rectCount uint32,
	rects []ClearRect,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param attachmentCount
	val1 := C.uint32_t(attachmentCount)
	// param attachments
	len3 := len(attachments)

	var arr4 *C.VkClearAttachment
	if len3 > 0 {
		arr4 = (*C.VkClearAttachment)(C.malloc(C.size_t(len3) * C.size_t(unsafe.Sizeof(*new(C.VkClearAttachment)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr4)) })
	}
	for i5, elem6 := range attachments {
		val7, cancel8 := elem6.toC()
		cancels = append(cancels, cancel8)
		(*[1 << 30]C.VkClearAttachment)(unsafe.Pointer(arr4))[i5] = *val7
	}
	// param rectCount
	val10 := C.uint32_t(rectCount)
	// param rects
	len12 := len(rects)

	var arr13 *C.VkClearRect
	if len12 > 0 {
		arr13 = (*C.VkClearRect)(C.malloc(C.size_t(len12) * C.size_t(unsafe.Sizeof(*new(C.VkClearRect)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr13)) })
	}
	for i14, elem15 := range rects {
		val16, cancel17 := elem15.toC()
		cancels = append(cancels, cancel17)
		(*[1 << 30]C.VkClearRect)(unsafe.Pointer(arr13))[i14] = *val16
	}
	C.fn_vkCmdClearAttachments(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, arr4, val10, arr13)
}

func (h CommandBuffer) ClearDepthStencilImage(
	image *Image,
	imageLayout ImageLayout,
	depthStencil *ClearDepthStencilValue,
	rangeCount uint32,
	ranges []ImageSubresourceRange,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param image
	h1 := C.VkImage(unsafe.Pointer(image.handle))
	// param imageLayout
	val3 := C.VkImageLayout(imageLayout)
	// param depthStencil
	var ptr5 *C.VkClearDepthStencilValue
	if depthStencil != nil {
		val6, cancel7 := depthStencil.toC()
		cancels = append(cancels, cancel7)
		ptr5 = val6
	}
	// param rangeCount
	val9 := C.uint32_t(rangeCount)
	// param ranges
	len11 := len(ranges)

	var arr12 *C.VkImageSubresourceRange
	if len11 > 0 {
		arr12 = (*C.VkImageSubresourceRange)(C.malloc(C.size_t(len11) * C.size_t(unsafe.Sizeof(*new(C.VkImageSubresourceRange)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr12)) })
	}
	for i13, elem14 := range ranges {
		val15, cancel16 := elem14.toC()
		cancels = append(cancels, cancel16)
		(*[1 << 30]C.VkImageSubresourceRange)(unsafe.Pointer(arr12))[i13] = *val15
	}
	C.fn_vkCmdClearDepthStencilImage(C.VkCommandBuffer(unsafe.Pointer(h.handle)), h1, val3, ptr5, val9, arr12)
}

func (h CommandBuffer) Draw(
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
	firstInstance uint32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param vertexCount
	val1 := C.uint32_t(vertexCount)
	// param instanceCount
	val3 := C.uint32_t(instanceCount)
	// param firstVertex
	val5 := C.uint32_t(firstVertex)
	// param firstInstance
	val7 := C.uint32_t(firstInstance)
	C.fn_vkCmdDraw(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3, val5, val7)
}

func (h CommandBuffer) DrawIndexed(
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	vertexOffset int32,
	firstInstance uint32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param indexCount
	val1 := C.uint32_t(indexCount)
	// param instanceCount
	val3 := C.uint32_t(instanceCount)
	// param firstIndex
	val5 := C.uint32_t(firstIndex)
	// param vertexOffset
	val7 := C.int32_t(vertexOffset)
	// param firstInstance
	val9 := C.uint32_t(firstInstance)
	C.fn_vkCmdDrawIndexed(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3, val5, val7, val9)
}

func (h CommandBuffer) DrawIndexedIndirect(
	buffer *Buffer,
	offset uint64,
	drawCount uint32,
	stride uint32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param buffer
	h1 := C.VkBuffer(unsafe.Pointer(buffer.handle))
	// param offset
	val3 := C.VkDeviceSize(offset)
	// param drawCount
	val5 := C.uint32_t(drawCount)
	// param stride
	val7 := C.uint32_t(stride)
	C.fn_vkCmdDrawIndexedIndirect(C.VkCommandBuffer(unsafe.Pointer(h.handle)), h1, val3, val5, val7)
}

func (h CommandBuffer) DrawIndexedIndirectCount(
	buffer *Buffer,
	offset uint64,
	countBuffer *Buffer,
	countBufferOffset uint64,
	maxDrawCount uint32,
	stride uint32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param buffer
	h1 := C.VkBuffer(unsafe.Pointer(buffer.handle))
	// param offset
	val3 := C.VkDeviceSize(offset)
	// param countBuffer
	h5 := C.VkBuffer(unsafe.Pointer(countBuffer.handle))
	// param countBufferOffset
	val7 := C.VkDeviceSize(countBufferOffset)
	// param maxDrawCount
	val9 := C.uint32_t(maxDrawCount)
	// param stride
	val11 := C.uint32_t(stride)
	C.fn_vkCmdDrawIndexedIndirectCount(C.VkCommandBuffer(unsafe.Pointer(h.handle)), h1, val3, h5, val7, val9, val11)
}

func (h CommandBuffer) DrawIndirect(
	buffer *Buffer,
	offset uint64,
	drawCount uint32,
	stride uint32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param buffer
	h1 := C.VkBuffer(unsafe.Pointer(buffer.handle))
	// param offset
	val3 := C.VkDeviceSize(offset)
	// param drawCount
	val5 := C.uint32_t(drawCount)
	// param stride
	val7 := C.uint32_t(stride)
	C.fn_vkCmdDrawIndirect(C.VkCommandBuffer(unsafe.Pointer(h.handle)), h1, val3, val5, val7)
}

func (h CommandBuffer) DrawIndirectCount(
	buffer *Buffer,
	offset uint64,
	countBuffer *Buffer,
	countBufferOffset uint64,
	maxDrawCount uint32,
	stride uint32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param buffer
	h1 := C.VkBuffer(unsafe.Pointer(buffer.handle))
	// param offset
	val3 := C.VkDeviceSize(offset)
	// param countBuffer
	h5 := C.VkBuffer(unsafe.Pointer(countBuffer.handle))
	// param countBufferOffset
	val7 := C.VkDeviceSize(countBufferOffset)
	// param maxDrawCount
	val9 := C.uint32_t(maxDrawCount)
	// param stride
	val11 := C.uint32_t(stride)
	C.fn_vkCmdDrawIndirectCount(C.VkCommandBuffer(unsafe.Pointer(h.handle)), h1, val3, h5, val7, val9, val11)
}

func (h CommandBuffer) EndRenderPass() {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	C.fn_vkCmdEndRenderPass(C.VkCommandBuffer(unsafe.Pointer(h.handle)))
}

func (h CommandBuffer) EndRenderPass2(
	subpassEndInfo *SubpassEndInfo,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param subpassEndInfo
	var ptr1 *C.VkSubpassEndInfo
	if subpassEndInfo != nil {
		val2, cancel3 := subpassEndInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	C.fn_vkCmdEndRenderPass2(C.VkCommandBuffer(unsafe.Pointer(h.handle)), ptr1)
}

func (h CommandBuffer) EndRendering() {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	C.fn_vkCmdEndRendering(C.VkCommandBuffer(unsafe.Pointer(h.handle)))
}

func (h CommandBuffer) NextSubpass(
	contents SubpassContents,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param contents
	val1 := C.VkSubpassContents(contents)
	C.fn_vkCmdNextSubpass(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) NextSubpass2(
	subpassBeginInfo *SubpassBeginInfo,
	subpassEndInfo *SubpassEndInfo,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param subpassBeginInfo
	var ptr1 *C.VkSubpassBeginInfo
	if subpassBeginInfo != nil {
		val2, cancel3 := subpassBeginInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	// param subpassEndInfo
	var ptr5 *C.VkSubpassEndInfo
	if subpassEndInfo != nil {
		val6, cancel7 := subpassEndInfo.toC()
		cancels = append(cancels, cancel7)
		ptr5 = val6
	}
	C.fn_vkCmdNextSubpass2(C.VkCommandBuffer(unsafe.Pointer(h.handle)), ptr1, ptr5)
}

func (h CommandBuffer) ResolveImage(
	srcImage *Image,
	srcImageLayout ImageLayout,
	dstImage *Image,
	dstImageLayout ImageLayout,
	regionCount uint32,
	regions []ImageResolve,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param srcImage
	h1 := C.VkImage(unsafe.Pointer(srcImage.handle))
	// param srcImageLayout
	val3 := C.VkImageLayout(srcImageLayout)
	// param dstImage
	h5 := C.VkImage(unsafe.Pointer(dstImage.handle))
	// param dstImageLayout
	val7 := C.VkImageLayout(dstImageLayout)
	// param regionCount
	val9 := C.uint32_t(regionCount)
	// param regions
	len11 := len(regions)

	var arr12 *C.VkImageResolve
	if len11 > 0 {
		arr12 = (*C.VkImageResolve)(C.malloc(C.size_t(len11) * C.size_t(unsafe.Sizeof(*new(C.VkImageResolve)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr12)) })
	}
	for i13, elem14 := range regions {
		val15, cancel16 := elem14.toC()
		cancels = append(cancels, cancel16)
		(*[1 << 30]C.VkImageResolve)(unsafe.Pointer(arr12))[i13] = *val15
	}
	C.fn_vkCmdResolveImage(C.VkCommandBuffer(unsafe.Pointer(h.handle)), h1, val3, h5, val7, val9, arr12)
}

func (h CommandBuffer) ResolveImage2(
	resolveImageInfo *ResolveImageInfo2,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param resolveImageInfo
	var ptr1 *C.VkResolveImageInfo2
	if resolveImageInfo != nil {
		val2, cancel3 := resolveImageInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	C.fn_vkCmdResolveImage2(C.VkCommandBuffer(unsafe.Pointer(h.handle)), ptr1)
}



func (h CommandBuffer) SetCullMode(
	cullMode CullModeFlags,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param cullMode
	val1 := C.VkCullModeFlags(cullMode)
	C.fn_vkCmdSetCullMode(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetDepthBias(
	depthBiasConstantFactor float32,
	depthBiasClamp float32,
	depthBiasSlopeFactor float32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param depthBiasConstantFactor
	val1 := C.float(depthBiasConstantFactor)
	// param depthBiasClamp
	val3 := C.float(depthBiasClamp)
	// param depthBiasSlopeFactor
	val5 := C.float(depthBiasSlopeFactor)
	C.fn_vkCmdSetDepthBias(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3, val5)
}

func (h CommandBuffer) SetDepthBiasEnable(
	depthBiasEnable bool,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param depthBiasEnable
	val1 := C.VkBool32(0)
	if depthBiasEnable {
		val1 = C.VkBool32(1)
	}
	C.fn_vkCmdSetDepthBiasEnable(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetDepthBounds(
	minDepthBounds float32,
	maxDepthBounds float32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param minDepthBounds
	val1 := C.float(minDepthBounds)
	// param maxDepthBounds
	val3 := C.float(maxDepthBounds)
	C.fn_vkCmdSetDepthBounds(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3)
}

func (h CommandBuffer) SetDepthBoundsTestEnable(
	depthBoundsTestEnable bool,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param depthBoundsTestEnable
	val1 := C.VkBool32(0)
	if depthBoundsTestEnable {
		val1 = C.VkBool32(1)
	}
	C.fn_vkCmdSetDepthBoundsTestEnable(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetDepthCompareOp(
	depthCompareOp CompareOp,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param depthCompareOp
	val1 := C.VkCompareOp(depthCompareOp)
	C.fn_vkCmdSetDepthCompareOp(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetDepthTestEnable(
	depthTestEnable bool,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param depthTestEnable
	val1 := C.VkBool32(0)
	if depthTestEnable {
		val1 = C.VkBool32(1)
	}
	C.fn_vkCmdSetDepthTestEnable(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetDepthWriteEnable(
	depthWriteEnable bool,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param depthWriteEnable
	val1 := C.VkBool32(0)
	if depthWriteEnable {
		val1 = C.VkBool32(1)
	}
	C.fn_vkCmdSetDepthWriteEnable(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetFrontFace(
	frontFace FrontFace,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param frontFace
	val1 := C.VkFrontFace(frontFace)
	C.fn_vkCmdSetFrontFace(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetLineStipple(
	lineStippleFactor uint32,
	lineStipplePattern uint16,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param lineStippleFactor
	val1 := C.uint32_t(lineStippleFactor)
	// param lineStipplePattern
	val3 := C.uint16_t(lineStipplePattern)
	C.fn_vkCmdSetLineStipple(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3)
}

func (h CommandBuffer) SetLineWidth(
	lineWidth float32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param lineWidth
	val1 := C.float(lineWidth)
	C.fn_vkCmdSetLineWidth(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetPrimitiveRestartEnable(
	primitiveRestartEnable bool,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param primitiveRestartEnable
	val1 := C.VkBool32(0)
	if primitiveRestartEnable {
		val1 = C.VkBool32(1)
	}
	C.fn_vkCmdSetPrimitiveRestartEnable(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetPrimitiveTopology(
	primitiveTopology PrimitiveTopology,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param primitiveTopology
	val1 := C.VkPrimitiveTopology(primitiveTopology)
	C.fn_vkCmdSetPrimitiveTopology(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetRasterizerDiscardEnable(
	rasterizerDiscardEnable bool,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param rasterizerDiscardEnable
	val1 := C.VkBool32(0)
	if rasterizerDiscardEnable {
		val1 = C.VkBool32(1)
	}
	C.fn_vkCmdSetRasterizerDiscardEnable(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetRenderingAttachmentLocations(
	locationInfo *RenderingAttachmentLocationInfo,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param locationInfo
	var ptr1 *C.VkRenderingAttachmentLocationInfo
	if locationInfo != nil {
		val2, cancel3 := locationInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	C.fn_vkCmdSetRenderingAttachmentLocations(C.VkCommandBuffer(unsafe.Pointer(h.handle)), ptr1)
}

func (h CommandBuffer) SetRenderingInputAttachmentIndices(
	inputAttachmentIndexInfo *RenderingInputAttachmentIndexInfo,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param inputAttachmentIndexInfo
	var ptr1 *C.VkRenderingInputAttachmentIndexInfo
	if inputAttachmentIndexInfo != nil {
		val2, cancel3 := inputAttachmentIndexInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	C.fn_vkCmdSetRenderingInputAttachmentIndices(C.VkCommandBuffer(unsafe.Pointer(h.handle)), ptr1)
}

func (h CommandBuffer) SetScissor(
	firstScissor uint32,
	scissorCount uint32,
	scissors []Rect2D,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param firstScissor
	val1 := C.uint32_t(firstScissor)
	// param scissorCount
	val3 := C.uint32_t(scissorCount)
	// param scissors
	len5 := len(scissors)

	var arr6 *C.VkRect2D
	if len5 > 0 {
		arr6 = (*C.VkRect2D)(C.malloc(C.size_t(len5) * C.size_t(unsafe.Sizeof(*new(C.VkRect2D)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr6)) })
	}
	for i7, elem8 := range scissors {
		val9, cancel10 := elem8.toC()
		cancels = append(cancels, cancel10)
		(*[1 << 30]C.VkRect2D)(unsafe.Pointer(arr6))[i7] = *val9
	}
	C.fn_vkCmdSetScissor(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3, arr6)
}

func (h CommandBuffer) SetScissorWithCount(
	scissorCount uint32,
	scissors []Rect2D,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param scissorCount
	val1 := C.uint32_t(scissorCount)
	// param scissors
	len3 := len(scissors)

	var arr4 *C.VkRect2D
	if len3 > 0 {
		arr4 = (*C.VkRect2D)(C.malloc(C.size_t(len3) * C.size_t(unsafe.Sizeof(*new(C.VkRect2D)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr4)) })
	}
	for i5, elem6 := range scissors {
		val7, cancel8 := elem6.toC()
		cancels = append(cancels, cancel8)
		(*[1 << 30]C.VkRect2D)(unsafe.Pointer(arr4))[i5] = *val7
	}
	C.fn_vkCmdSetScissorWithCount(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, arr4)
}

func (h CommandBuffer) SetStencilCompareMask(
	faceMask StencilFaceFlags,
	compareMask uint32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param faceMask
	val1 := C.VkStencilFaceFlags(faceMask)
	// param compareMask
	val3 := C.uint32_t(compareMask)
	C.fn_vkCmdSetStencilCompareMask(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3)
}

func (h CommandBuffer) SetStencilOp(
	faceMask StencilFaceFlags,
	failOp StencilOp,
	passOp StencilOp,
	depthFailOp StencilOp,
	compareOp CompareOp,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param faceMask
	val1 := C.VkStencilFaceFlags(faceMask)
	// param failOp
	val3 := C.VkStencilOp(failOp)
	// param passOp
	val5 := C.VkStencilOp(passOp)
	// param depthFailOp
	val7 := C.VkStencilOp(depthFailOp)
	// param compareOp
	val9 := C.VkCompareOp(compareOp)
	C.fn_vkCmdSetStencilOp(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3, val5, val7, val9)
}

func (h CommandBuffer) SetStencilReference(
	faceMask StencilFaceFlags,
	reference uint32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param faceMask
	val1 := C.VkStencilFaceFlags(faceMask)
	// param reference
	val3 := C.uint32_t(reference)
	C.fn_vkCmdSetStencilReference(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3)
}

func (h CommandBuffer) SetStencilTestEnable(
	stencilTestEnable bool,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param stencilTestEnable
	val1 := C.VkBool32(0)
	if stencilTestEnable {
		val1 = C.VkBool32(1)
	}
	C.fn_vkCmdSetStencilTestEnable(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1)
}

func (h CommandBuffer) SetStencilWriteMask(
	faceMask StencilFaceFlags,
	writeMask uint32,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param faceMask
	val1 := C.VkStencilFaceFlags(faceMask)
	// param writeMask
	val3 := C.uint32_t(writeMask)
	C.fn_vkCmdSetStencilWriteMask(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3)
}

func (h CommandBuffer) SetViewport(
	firstViewport uint32,
	viewportCount uint32,
	viewports []Viewport,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param firstViewport
	val1 := C.uint32_t(firstViewport)
	// param viewportCount
	val3 := C.uint32_t(viewportCount)
	// param viewports
	len5 := len(viewports)

	var arr6 *C.VkViewport
	if len5 > 0 {
		arr6 = (*C.VkViewport)(C.malloc(C.size_t(len5) * C.size_t(unsafe.Sizeof(*new(C.VkViewport)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr6)) })
	}
	for i7, elem8 := range viewports {
		val9, cancel10 := elem8.toC()
		cancels = append(cancels, cancel10)
		(*[1 << 30]C.VkViewport)(unsafe.Pointer(arr6))[i7] = *val9
	}
	C.fn_vkCmdSetViewport(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, val3, arr6)
}

func (h CommandBuffer) SetViewportWithCount(
	viewportCount uint32,
	viewports []Viewport,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param viewportCount
	val1 := C.uint32_t(viewportCount)
	// param viewports
	len3 := len(viewports)

	var arr4 *C.VkViewport
	if len3 > 0 {
		arr4 = (*C.VkViewport)(C.malloc(C.size_t(len3) * C.size_t(unsafe.Sizeof(*new(C.VkViewport)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr4)) })
	}
	for i5, elem6 := range viewports {
		val7, cancel8 := elem6.toC()
		cancels = append(cancels, cancel8)
		(*[1 << 30]C.VkViewport)(unsafe.Pointer(arr4))[i5] = *val7
	}
	C.fn_vkCmdSetViewportWithCount(C.VkCommandBuffer(unsafe.Pointer(h.handle)), val1, arr4)
}

func (h Device) CreateFramebuffer(
	createInfo *FramebufferCreateInfo,
	allocator *AllocationCallbacks,
) (*Framebuffer, error) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param createInfo
	var ptr1 *C.VkFramebufferCreateInfo
	if createInfo != nil {
		val2, cancel3 := createInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	// param allocator
	var ptr5 *C.VkAllocationCallbacks
	if allocator != nil {
		val6, cancel7 := allocator.toC()
		cancels = append(cancels, cancel7)
		ptr5 = val6
	}
	var framebufferOut C.VkFramebuffer
	_result := C.fn_vkCreateFramebuffer(C.VkDevice(unsafe.Pointer(h.handle)), ptr1, ptr5, &framebufferOut)
	if _result != C.VK_SUCCESS {
		return nil, vkError(_result)
	}
	h8 := &Framebuffer{handle: unsafe.Pointer(framebufferOut)}
	return h8, nil
}

func (h Device) CreateGraphicsPipelines(
	pipelineCache *PipelineCache,
	createInfoCount uint32,
	createInfos []GraphicsPipelineCreateInfo,
	allocator *AllocationCallbacks,
) (*Pipeline, error) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param pipelineCache
	h1 := C.VkPipelineCache(unsafe.Pointer(pipelineCache.handle))
	// param createInfoCount
	val3 := C.uint32_t(createInfoCount)
	// param createInfos
	len5 := len(createInfos)

	var arr6 *C.VkGraphicsPipelineCreateInfo
	if len5 > 0 {
		arr6 = (*C.VkGraphicsPipelineCreateInfo)(C.malloc(C.size_t(len5) * C.size_t(unsafe.Sizeof(*new(C.VkGraphicsPipelineCreateInfo)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(arr6)) })
	}
	for i7, elem8 := range createInfos {
		val9, cancel10 := elem8.toC()
		cancels = append(cancels, cancel10)
		(*[1 << 30]C.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(arr6))[i7] = *val9
	}
	// param allocator
	var ptr12 *C.VkAllocationCallbacks
	if allocator != nil {
		val13, cancel14 := allocator.toC()
		cancels = append(cancels, cancel14)
		ptr12 = val13
	}
	var pipelinesOut C.VkPipeline
	_result := C.fn_vkCreateGraphicsPipelines(C.VkDevice(unsafe.Pointer(h.handle)), h1, val3, arr6, ptr12, &pipelinesOut)
	if _result != C.VK_SUCCESS {
		return nil, vkError(_result)
	}
	h15 := &Pipeline{handle: unsafe.Pointer(pipelinesOut)}
	return h15, nil
}

func (h Device) CreateRenderPass(
	createInfo *RenderPassCreateInfo,
	allocator *AllocationCallbacks,
) (*RenderPass, error) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param createInfo
	var ptr1 *C.VkRenderPassCreateInfo
	if createInfo != nil {
		val2, cancel3 := createInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	// param allocator
	var ptr5 *C.VkAllocationCallbacks
	if allocator != nil {
		val6, cancel7 := allocator.toC()
		cancels = append(cancels, cancel7)
		ptr5 = val6
	}
	var renderPassOut C.VkRenderPass
	_result := C.fn_vkCreateRenderPass(C.VkDevice(unsafe.Pointer(h.handle)), ptr1, ptr5, &renderPassOut)
	if _result != C.VK_SUCCESS {
		return nil, vkError(_result)
	}
	h8 := &RenderPass{handle: unsafe.Pointer(renderPassOut)}
	return h8, nil
}

func (h Device) CreateRenderPass2(
	createInfo *RenderPassCreateInfo2,
	allocator *AllocationCallbacks,
) (*RenderPass, error) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param createInfo
	var ptr1 *C.VkRenderPassCreateInfo2
	if createInfo != nil {
		val2, cancel3 := createInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = val2
	}
	// param allocator
	var ptr5 *C.VkAllocationCallbacks
	if allocator != nil {
		val6, cancel7 := allocator.toC()
		cancels = append(cancels, cancel7)
		ptr5 = val6
	}
	var renderPassOut C.VkRenderPass
	_result := C.fn_vkCreateRenderPass2(C.VkDevice(unsafe.Pointer(h.handle)), ptr1, ptr5, &renderPassOut)
	if _result != C.VK_SUCCESS {
		return nil, vkError(_result)
	}
	h8 := &RenderPass{handle: unsafe.Pointer(renderPassOut)}
	return h8, nil
}

func (h Device) DestroyFramebuffer(
	framebuffer *Framebuffer,
	allocator *AllocationCallbacks,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param framebuffer
	h1 := C.VkFramebuffer(unsafe.Pointer(framebuffer.handle))
	// param allocator
	var ptr3 *C.VkAllocationCallbacks
	if allocator != nil {
		val4, cancel5 := allocator.toC()
		cancels = append(cancels, cancel5)
		ptr3 = val4
	}
	C.fn_vkDestroyFramebuffer(C.VkDevice(unsafe.Pointer(h.handle)), h1, ptr3)
}

func (h Device) DestroyRenderPass(
	renderPass *RenderPass,
	allocator *AllocationCallbacks,
) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param renderPass
	h1 := C.VkRenderPass(unsafe.Pointer(renderPass.handle))
	// param allocator
	var ptr3 *C.VkAllocationCallbacks
	if allocator != nil {
		val4, cancel5 := allocator.toC()
		cancels = append(cancels, cancel5)
		ptr3 = val4
	}
	C.fn_vkDestroyRenderPass(C.VkDevice(unsafe.Pointer(h.handle)), h1, ptr3)
}