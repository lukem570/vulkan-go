package vulkan

var (
	VkMaxPhysicalDeviceNameSize                              = 256
	VkUuidSize                                               = 16
	VkLuidSize                                               = 8
	VkMaxExtensionNameSize                                   = 256
	VkMaxDescriptionSize                                     = 256
	VkMaxMemoryTypes                                         = 32
	VkMaxMemoryHeaps                                         = 16
	VkLodClampNone                                           = 1000.0
	VkRemainingMipLevels                                     = ^0
	VkRemainingArrayLayers                                   = ^0
	VkRemaining3dSlicesExt                                   = ^0
	VkWholeSize                                              = ^0
	VkAttachmentUnused                                       = ^0
	VkTrue                                                   = 1
	VkFalse                                                  = 0
	VkQueueFamilyIgnored                                     = ^0
	VkQueueFamilyExternal                                    = ^1
	VkQueueFamilyForeignExt                                  = ^2
	VkSubpassExternal                                        = ^0
	VkMaxDeviceGroupSize                                     = 32
	VkMaxDriverNameSize                                      = 256
	VkMaxDriverInfoSize                                      = 256
	VkShaderUnusedKhr                                        = ^0
	VkMaxGlobalPrioritySize                                  = 16
	VkMaxShaderModuleIdentifierSizeExt                       = 32
	VkMaxPipelineBinaryKeySizeKhr                            = 32
	VkMaxVideoAv1ReferencesPerFrameKhr                       = 7
	VkMaxVideoVp9ReferencesPerFrameKhr                       = 3
	VkShaderIndexUnusedAmdx                                  = ^0
	VkPartitionedAccelerationStructurePartitionIndexGlobalNv = ^0
	VkCompressedTriangleFormatDgf1ByteAlignmentAmdx          = 128
	VkCompressedTriangleFormatDgf1ByteStrideAmdx             = 128
	VkMaxPhysicalDeviceDataGraphOperationSetNameSizeArm      = 128
	VkDataGraphModelToolchainVersionLengthQcom               = 3
	VkComputeOccupancyPriorityLowNv                          = 0.25
	VkComputeOccupancyPriorityNormalNv                       = 0.50
	VkComputeOccupancyPriorityHighNv                         = 0.75
)

type VkImageLayout int64

var (
	VkImageLayoutUndefined                     VkImageLayout = 0
	VkImageLayoutGeneral                       VkImageLayout = 1
	VkImageLayoutColorAttachmentOptimal        VkImageLayout = 2
	VkImageLayoutDepthStencilAttachmentOptimal VkImageLayout = 3
	VkImageLayoutDepthStencilReadOnlyOptimal   VkImageLayout = 4
	VkImageLayoutShaderReadOnlyOptimal         VkImageLayout = 5
	VkImageLayoutTransferSrcOptimal            VkImageLayout = 6
	VkImageLayoutTransferDstOptimal            VkImageLayout = 7
	VkImageLayoutPreinitialized                VkImageLayout = 8
)

type VkAttachmentLoadOp int64

var (
	VkAttachmentLoadOpLoad     VkAttachmentLoadOp = 0
	VkAttachmentLoadOpClear    VkAttachmentLoadOp = 1
	VkAttachmentLoadOpDontCare VkAttachmentLoadOp = 2
)

type VkAttachmentStoreOp int64

var (
	VkAttachmentStoreOpStore    VkAttachmentStoreOp = 0
	VkAttachmentStoreOpDontCare VkAttachmentStoreOp = 1
)

type VkImageType int64

var (
	VkImageType1d VkImageType = 0
	VkImageType2d VkImageType = 1
	VkImageType3d VkImageType = 2
)

type VkImageTiling int64

var (
	VkImageTilingOptimal VkImageTiling = 0
	VkImageTilingLinear  VkImageTiling = 1
)

type VkImageViewType int64

var (
	VkImageViewType1d        VkImageViewType = 0
	VkImageViewType2d        VkImageViewType = 1
	VkImageViewType3d        VkImageViewType = 2
	VkImageViewTypeCube      VkImageViewType = 3
	VkImageViewType1dArray   VkImageViewType = 4
	VkImageViewType2dArray   VkImageViewType = 5
	VkImageViewTypeCubeArray VkImageViewType = 6
)

type VkCommandBufferLevel int64

var (
	VkCommandBufferLevelPrimary   VkCommandBufferLevel = 0
	VkCommandBufferLevelSecondary VkCommandBufferLevel = 1
)

type VkComponentSwizzle int64

var (
	VkComponentSwizzleIdentity VkComponentSwizzle = 0
	VkComponentSwizzleZero     VkComponentSwizzle = 1
	VkComponentSwizzleOne      VkComponentSwizzle = 2
	VkComponentSwizzleR        VkComponentSwizzle = 3
	VkComponentSwizzleG        VkComponentSwizzle = 4
	VkComponentSwizzleB        VkComponentSwizzle = 5
	VkComponentSwizzleA        VkComponentSwizzle = 6
)

type VkDescriptorType int64

var (
	VkDescriptorTypeSampler              VkDescriptorType = 0
	VkDescriptorTypeCombinedImageSampler VkDescriptorType = 1
	VkDescriptorTypeSampledImage         VkDescriptorType = 2
	VkDescriptorTypeStorageImage         VkDescriptorType = 3
	VkDescriptorTypeUniformTexelBuffer   VkDescriptorType = 4
	VkDescriptorTypeStorageTexelBuffer   VkDescriptorType = 5
	VkDescriptorTypeUniformBuffer        VkDescriptorType = 6
	VkDescriptorTypeStorageBuffer        VkDescriptorType = 7
	VkDescriptorTypeUniformBufferDynamic VkDescriptorType = 8
	VkDescriptorTypeStorageBufferDynamic VkDescriptorType = 9
	VkDescriptorTypeInputAttachment      VkDescriptorType = 10
)

type VkQueryType int64

var (
	VkQueryTypeOcclusion          VkQueryType = 0
	VkQueryTypePipelineStatistics VkQueryType = 1
	VkQueryTypeTimestamp          VkQueryType = 2
)

type VkBorderColor int64

var (
	VkBorderColorFloatTransparentBlack VkBorderColor = 0
	VkBorderColorIntTransparentBlack   VkBorderColor = 1
	VkBorderColorFloatOpaqueBlack      VkBorderColor = 2
	VkBorderColorIntOpaqueBlack        VkBorderColor = 3
	VkBorderColorFloatOpaqueWhite      VkBorderColor = 4
	VkBorderColorIntOpaqueWhite        VkBorderColor = 5
)

type VkPipelineBindPoint int64

var (
	VkPipelineBindPointGraphics VkPipelineBindPoint = 0
	VkPipelineBindPointCompute  VkPipelineBindPoint = 1
)

type VkPipelineCacheHeaderVersion int64

var (
	VkPipelineCacheHeaderVersionOne VkPipelineCacheHeaderVersion = 1
)

type VkPipelineCacheCreateFlagBits int64

var ()

type VkPrimitiveTopology int64

var (
	VkPrimitiveTopologyPointList                  VkPrimitiveTopology = 0
	VkPrimitiveTopologyLineList                   VkPrimitiveTopology = 1
	VkPrimitiveTopologyLineStrip                  VkPrimitiveTopology = 2
	VkPrimitiveTopologyTriangleList               VkPrimitiveTopology = 3
	VkPrimitiveTopologyTriangleStrip              VkPrimitiveTopology = 4
	VkPrimitiveTopologyTriangleFan                VkPrimitiveTopology = 5
	VkPrimitiveTopologyLineListWithAdjacency      VkPrimitiveTopology = 6
	VkPrimitiveTopologyLineStripWithAdjacency     VkPrimitiveTopology = 7
	VkPrimitiveTopologyTriangleListWithAdjacency  VkPrimitiveTopology = 8
	VkPrimitiveTopologyTriangleStripWithAdjacency VkPrimitiveTopology = 9
	VkPrimitiveTopologyPatchList                  VkPrimitiveTopology = 10
)

type VkSharingMode int64

var (
	VkSharingModeExclusive  VkSharingMode = 0
	VkSharingModeConcurrent VkSharingMode = 1
)

type VkIndexType int64

var (
	VkIndexTypeUint16 VkIndexType = 0
	VkIndexTypeUint32 VkIndexType = 1
)

type VkFilter int64

var (
	VkFilterNearest VkFilter = 0
	VkFilterLinear  VkFilter = 1
)

type VkSamplerMipmapMode int64

var (
	VkSamplerMipmapModeNearest VkSamplerMipmapMode = 0
	VkSamplerMipmapModeLinear  VkSamplerMipmapMode = 1
)

type VkSamplerAddressMode int64

var (
	VkSamplerAddressModeRepeat         VkSamplerAddressMode = 0
	VkSamplerAddressModeMirroredRepeat VkSamplerAddressMode = 1
	VkSamplerAddressModeClampToEdge    VkSamplerAddressMode = 2
	VkSamplerAddressModeClampToBorder  VkSamplerAddressMode = 3
)

type VkCompareOp int64

var (
	VkCompareOpNever          VkCompareOp = 0
	VkCompareOpLess           VkCompareOp = 1
	VkCompareOpEqual          VkCompareOp = 2
	VkCompareOpLessOrEqual    VkCompareOp = 3
	VkCompareOpGreater        VkCompareOp = 4
	VkCompareOpNotEqual       VkCompareOp = 5
	VkCompareOpGreaterOrEqual VkCompareOp = 6
	VkCompareOpAlways         VkCompareOp = 7
)

type VkPolygonMode int64

var (
	VkPolygonModeFill  VkPolygonMode = 0
	VkPolygonModeLine  VkPolygonMode = 1
	VkPolygonModePoint VkPolygonMode = 2
)

type VkFrontFace int64

var (
	VkFrontFaceCounterClockwise VkFrontFace = 0
	VkFrontFaceClockwise        VkFrontFace = 1
)

type VkBlendFactor int64

var (
	VkBlendFactorZero                  VkBlendFactor = 0
	VkBlendFactorOne                   VkBlendFactor = 1
	VkBlendFactorSrcColor              VkBlendFactor = 2
	VkBlendFactorOneMinusSrcColor      VkBlendFactor = 3
	VkBlendFactorDstColor              VkBlendFactor = 4
	VkBlendFactorOneMinusDstColor      VkBlendFactor = 5
	VkBlendFactorSrcAlpha              VkBlendFactor = 6
	VkBlendFactorOneMinusSrcAlpha      VkBlendFactor = 7
	VkBlendFactorDstAlpha              VkBlendFactor = 8
	VkBlendFactorOneMinusDstAlpha      VkBlendFactor = 9
	VkBlendFactorConstantColor         VkBlendFactor = 10
	VkBlendFactorOneMinusConstantColor VkBlendFactor = 11
	VkBlendFactorConstantAlpha         VkBlendFactor = 12
	VkBlendFactorOneMinusConstantAlpha VkBlendFactor = 13
	VkBlendFactorSrcAlphaSaturate      VkBlendFactor = 14
	VkBlendFactorSrc1Color             VkBlendFactor = 15
	VkBlendFactorOneMinusSrc1Color     VkBlendFactor = 16
	VkBlendFactorSrc1Alpha             VkBlendFactor = 17
	VkBlendFactorOneMinusSrc1Alpha     VkBlendFactor = 18
)

type VkBlendOp int64

var (
	VkBlendOpAdd             VkBlendOp = 0
	VkBlendOpSubtract        VkBlendOp = 1
	VkBlendOpReverseSubtract VkBlendOp = 2
	VkBlendOpMin             VkBlendOp = 3
	VkBlendOpMax             VkBlendOp = 4
)

type VkStencilOp int64

var (
	VkStencilOpKeep              VkStencilOp = 0
	VkStencilOpZero              VkStencilOp = 1
	VkStencilOpReplace           VkStencilOp = 2
	VkStencilOpIncrementAndClamp VkStencilOp = 3
	VkStencilOpDecrementAndClamp VkStencilOp = 4
	VkStencilOpInvert            VkStencilOp = 5
	VkStencilOpIncrementAndWrap  VkStencilOp = 6
	VkStencilOpDecrementAndWrap  VkStencilOp = 7
)

type VkLogicOp int64

var (
	VkLogicOpClear        VkLogicOp = 0
	VkLogicOpAnd          VkLogicOp = 1
	VkLogicOpAndReverse   VkLogicOp = 2
	VkLogicOpCopy         VkLogicOp = 3
	VkLogicOpAndInverted  VkLogicOp = 4
	VkLogicOpNoOp         VkLogicOp = 5
	VkLogicOpXor          VkLogicOp = 6
	VkLogicOpOr           VkLogicOp = 7
	VkLogicOpNor          VkLogicOp = 8
	VkLogicOpEquivalent   VkLogicOp = 9
	VkLogicOpInvert       VkLogicOp = 10
	VkLogicOpOrReverse    VkLogicOp = 11
	VkLogicOpCopyInverted VkLogicOp = 12
	VkLogicOpOrInverted   VkLogicOp = 13
	VkLogicOpNand         VkLogicOp = 14
	VkLogicOpSet          VkLogicOp = 15
)

type VkInternalAllocationType int64

var (
	VkInternalAllocationTypeExecutable VkInternalAllocationType = 0
)

type VkSystemAllocationScope int64

var (
	VkSystemAllocationScopeCommand  VkSystemAllocationScope = 0
	VkSystemAllocationScopeObject   VkSystemAllocationScope = 1
	VkSystemAllocationScopeCache    VkSystemAllocationScope = 2
	VkSystemAllocationScopeDevice   VkSystemAllocationScope = 3
	VkSystemAllocationScopeInstance VkSystemAllocationScope = 4
)

type VkPhysicalDeviceType int64

var (
	VkPhysicalDeviceTypeOther         VkPhysicalDeviceType = 0
	VkPhysicalDeviceTypeIntegratedGpu VkPhysicalDeviceType = 1
	VkPhysicalDeviceTypeDiscreteGpu   VkPhysicalDeviceType = 2
	VkPhysicalDeviceTypeVirtualGpu    VkPhysicalDeviceType = 3
	VkPhysicalDeviceTypeCpu           VkPhysicalDeviceType = 4
)

type VkVertexInputRate int64

var (
	VkVertexInputRateVertex   VkVertexInputRate = 0
	VkVertexInputRateInstance VkVertexInputRate = 1
)

type VkFormat int64

var (
	VkFormatUndefined                VkFormat = 0
	VkFormatR4g4UnormPack8           VkFormat = 1
	VkFormatR4g4b4a4UnormPack16      VkFormat = 2
	VkFormatB4g4r4a4UnormPack16      VkFormat = 3
	VkFormatR5g6b5UnormPack16        VkFormat = 4
	VkFormatB5g6r5UnormPack16        VkFormat = 5
	VkFormatR5g5b5a1UnormPack16      VkFormat = 6
	VkFormatB5g5r5a1UnormPack16      VkFormat = 7
	VkFormatA1r5g5b5UnormPack16      VkFormat = 8
	VkFormatR8Unorm                  VkFormat = 9
	VkFormatR8Snorm                  VkFormat = 10
	VkFormatR8Uscaled                VkFormat = 11
	VkFormatR8Sscaled                VkFormat = 12
	VkFormatR8Uint                   VkFormat = 13
	VkFormatR8Sint                   VkFormat = 14
	VkFormatR8Srgb                   VkFormat = 15
	VkFormatR8g8Unorm                VkFormat = 16
	VkFormatR8g8Snorm                VkFormat = 17
	VkFormatR8g8Uscaled              VkFormat = 18
	VkFormatR8g8Sscaled              VkFormat = 19
	VkFormatR8g8Uint                 VkFormat = 20
	VkFormatR8g8Sint                 VkFormat = 21
	VkFormatR8g8Srgb                 VkFormat = 22
	VkFormatR8g8b8Unorm              VkFormat = 23
	VkFormatR8g8b8Snorm              VkFormat = 24
	VkFormatR8g8b8Uscaled            VkFormat = 25
	VkFormatR8g8b8Sscaled            VkFormat = 26
	VkFormatR8g8b8Uint               VkFormat = 27
	VkFormatR8g8b8Sint               VkFormat = 28
	VkFormatR8g8b8Srgb               VkFormat = 29
	VkFormatB8g8r8Unorm              VkFormat = 30
	VkFormatB8g8r8Snorm              VkFormat = 31
	VkFormatB8g8r8Uscaled            VkFormat = 32
	VkFormatB8g8r8Sscaled            VkFormat = 33
	VkFormatB8g8r8Uint               VkFormat = 34
	VkFormatB8g8r8Sint               VkFormat = 35
	VkFormatB8g8r8Srgb               VkFormat = 36
	VkFormatR8g8b8a8Unorm            VkFormat = 37
	VkFormatR8g8b8a8Snorm            VkFormat = 38
	VkFormatR8g8b8a8Uscaled          VkFormat = 39
	VkFormatR8g8b8a8Sscaled          VkFormat = 40
	VkFormatR8g8b8a8Uint             VkFormat = 41
	VkFormatR8g8b8a8Sint             VkFormat = 42
	VkFormatR8g8b8a8Srgb             VkFormat = 43
	VkFormatB8g8r8a8Unorm            VkFormat = 44
	VkFormatB8g8r8a8Snorm            VkFormat = 45
	VkFormatB8g8r8a8Uscaled          VkFormat = 46
	VkFormatB8g8r8a8Sscaled          VkFormat = 47
	VkFormatB8g8r8a8Uint             VkFormat = 48
	VkFormatB8g8r8a8Sint             VkFormat = 49
	VkFormatB8g8r8a8Srgb             VkFormat = 50
	VkFormatA8b8g8r8UnormPack32      VkFormat = 51
	VkFormatA8b8g8r8SnormPack32      VkFormat = 52
	VkFormatA8b8g8r8UscaledPack32    VkFormat = 53
	VkFormatA8b8g8r8SscaledPack32    VkFormat = 54
	VkFormatA8b8g8r8UintPack32       VkFormat = 55
	VkFormatA8b8g8r8SintPack32       VkFormat = 56
	VkFormatA8b8g8r8SrgbPack32       VkFormat = 57
	VkFormatA2r10g10b10UnormPack32   VkFormat = 58
	VkFormatA2r10g10b10SnormPack32   VkFormat = 59
	VkFormatA2r10g10b10UscaledPack32 VkFormat = 60
	VkFormatA2r10g10b10SscaledPack32 VkFormat = 61
	VkFormatA2r10g10b10UintPack32    VkFormat = 62
	VkFormatA2r10g10b10SintPack32    VkFormat = 63
	VkFormatA2b10g10r10UnormPack32   VkFormat = 64
	VkFormatA2b10g10r10SnormPack32   VkFormat = 65
	VkFormatA2b10g10r10UscaledPack32 VkFormat = 66
	VkFormatA2b10g10r10SscaledPack32 VkFormat = 67
	VkFormatA2b10g10r10UintPack32    VkFormat = 68
	VkFormatA2b10g10r10SintPack32    VkFormat = 69
	VkFormatR16Unorm                 VkFormat = 70
	VkFormatR16Snorm                 VkFormat = 71
	VkFormatR16Uscaled               VkFormat = 72
	VkFormatR16Sscaled               VkFormat = 73
	VkFormatR16Uint                  VkFormat = 74
	VkFormatR16Sint                  VkFormat = 75
	VkFormatR16Sfloat                VkFormat = 76
	VkFormatR16g16Unorm              VkFormat = 77
	VkFormatR16g16Snorm              VkFormat = 78
	VkFormatR16g16Uscaled            VkFormat = 79
	VkFormatR16g16Sscaled            VkFormat = 80
	VkFormatR16g16Uint               VkFormat = 81
	VkFormatR16g16Sint               VkFormat = 82
	VkFormatR16g16Sfloat             VkFormat = 83
	VkFormatR16g16b16Unorm           VkFormat = 84
	VkFormatR16g16b16Snorm           VkFormat = 85
	VkFormatR16g16b16Uscaled         VkFormat = 86
	VkFormatR16g16b16Sscaled         VkFormat = 87
	VkFormatR16g16b16Uint            VkFormat = 88
	VkFormatR16g16b16Sint            VkFormat = 89
	VkFormatR16g16b16Sfloat          VkFormat = 90
	VkFormatR16g16b16a16Unorm        VkFormat = 91
	VkFormatR16g16b16a16Snorm        VkFormat = 92
	VkFormatR16g16b16a16Uscaled      VkFormat = 93
	VkFormatR16g16b16a16Sscaled      VkFormat = 94
	VkFormatR16g16b16a16Uint         VkFormat = 95
	VkFormatR16g16b16a16Sint         VkFormat = 96
	VkFormatR16g16b16a16Sfloat       VkFormat = 97
	VkFormatR32Uint                  VkFormat = 98
	VkFormatR32Sint                  VkFormat = 99
	VkFormatR32Sfloat                VkFormat = 100
	VkFormatR32g32Uint               VkFormat = 101
	VkFormatR32g32Sint               VkFormat = 102
	VkFormatR32g32Sfloat             VkFormat = 103
	VkFormatR32g32b32Uint            VkFormat = 104
	VkFormatR32g32b32Sint            VkFormat = 105
	VkFormatR32g32b32Sfloat          VkFormat = 106
	VkFormatR32g32b32a32Uint         VkFormat = 107
	VkFormatR32g32b32a32Sint         VkFormat = 108
	VkFormatR32g32b32a32Sfloat       VkFormat = 109
	VkFormatR64Uint                  VkFormat = 110
	VkFormatR64Sint                  VkFormat = 111
	VkFormatR64Sfloat                VkFormat = 112
	VkFormatR64g64Uint               VkFormat = 113
	VkFormatR64g64Sint               VkFormat = 114
	VkFormatR64g64Sfloat             VkFormat = 115
	VkFormatR64g64b64Uint            VkFormat = 116
	VkFormatR64g64b64Sint            VkFormat = 117
	VkFormatR64g64b64Sfloat          VkFormat = 118
	VkFormatR64g64b64a64Uint         VkFormat = 119
	VkFormatR64g64b64a64Sint         VkFormat = 120
	VkFormatR64g64b64a64Sfloat       VkFormat = 121
	VkFormatB10g11r11UfloatPack32    VkFormat = 122
	VkFormatE5b9g9r9UfloatPack32     VkFormat = 123
	VkFormatD16Unorm                 VkFormat = 124
	VkFormatX8D24UnormPack32         VkFormat = 125
	VkFormatD32Sfloat                VkFormat = 126
	VkFormatS8Uint                   VkFormat = 127
	VkFormatD16UnormS8Uint           VkFormat = 128
	VkFormatD24UnormS8Uint           VkFormat = 129
	VkFormatD32SfloatS8Uint          VkFormat = 130
	VkFormatBc1RgbUnormBlock         VkFormat = 131
	VkFormatBc1RgbSrgbBlock          VkFormat = 132
	VkFormatBc1RgbaUnormBlock        VkFormat = 133
	VkFormatBc1RgbaSrgbBlock         VkFormat = 134
	VkFormatBc2UnormBlock            VkFormat = 135
	VkFormatBc2SrgbBlock             VkFormat = 136
	VkFormatBc3UnormBlock            VkFormat = 137
	VkFormatBc3SrgbBlock             VkFormat = 138
	VkFormatBc4UnormBlock            VkFormat = 139
	VkFormatBc4SnormBlock            VkFormat = 140
	VkFormatBc5UnormBlock            VkFormat = 141
	VkFormatBc5SnormBlock            VkFormat = 142
	VkFormatBc6hUfloatBlock          VkFormat = 143
	VkFormatBc6hSfloatBlock          VkFormat = 144
	VkFormatBc7UnormBlock            VkFormat = 145
	VkFormatBc7SrgbBlock             VkFormat = 146
	VkFormatEtc2R8g8b8UnormBlock     VkFormat = 147
	VkFormatEtc2R8g8b8SrgbBlock      VkFormat = 148
	VkFormatEtc2R8g8b8a1UnormBlock   VkFormat = 149
	VkFormatEtc2R8g8b8a1SrgbBlock    VkFormat = 150
	VkFormatEtc2R8g8b8a8UnormBlock   VkFormat = 151
	VkFormatEtc2R8g8b8a8SrgbBlock    VkFormat = 152
	VkFormatEacR11UnormBlock         VkFormat = 153
	VkFormatEacR11SnormBlock         VkFormat = 154
	VkFormatEacR11g11UnormBlock      VkFormat = 155
	VkFormatEacR11g11SnormBlock      VkFormat = 156
	VkFormatAstc4x4UnormBlock        VkFormat = 157
	VkFormatAstc4x4SrgbBlock         VkFormat = 158
	VkFormatAstc5x4UnormBlock        VkFormat = 159
	VkFormatAstc5x4SrgbBlock         VkFormat = 160
	VkFormatAstc5x5UnormBlock        VkFormat = 161
	VkFormatAstc5x5SrgbBlock         VkFormat = 162
	VkFormatAstc6x5UnormBlock        VkFormat = 163
	VkFormatAstc6x5SrgbBlock         VkFormat = 164
	VkFormatAstc6x6UnormBlock        VkFormat = 165
	VkFormatAstc6x6SrgbBlock         VkFormat = 166
	VkFormatAstc8x5UnormBlock        VkFormat = 167
	VkFormatAstc8x5SrgbBlock         VkFormat = 168
	VkFormatAstc8x6UnormBlock        VkFormat = 169
	VkFormatAstc8x6SrgbBlock         VkFormat = 170
	VkFormatAstc8x8UnormBlock        VkFormat = 171
	VkFormatAstc8x8SrgbBlock         VkFormat = 172
	VkFormatAstc10x5UnormBlock       VkFormat = 173
	VkFormatAstc10x5SrgbBlock        VkFormat = 174
	VkFormatAstc10x6UnormBlock       VkFormat = 175
	VkFormatAstc10x6SrgbBlock        VkFormat = 176
	VkFormatAstc10x8UnormBlock       VkFormat = 177
	VkFormatAstc10x8SrgbBlock        VkFormat = 178
	VkFormatAstc10x10UnormBlock      VkFormat = 179
	VkFormatAstc10x10SrgbBlock       VkFormat = 180
	VkFormatAstc12x10UnormBlock      VkFormat = 181
	VkFormatAstc12x10SrgbBlock       VkFormat = 182
	VkFormatAstc12x12UnormBlock      VkFormat = 183
	VkFormatAstc12x12SrgbBlock       VkFormat = 184
)

type VkStructureType int64

var (
	VkStructureTypeApplicationInfo                      VkStructureType = 0
	VkStructureTypeInstanceCreateInfo                   VkStructureType = 1
	VkStructureTypeDeviceQueueCreateInfo                VkStructureType = 2
	VkStructureTypeDeviceCreateInfo                     VkStructureType = 3
	VkStructureTypeSubmitInfo                           VkStructureType = 4
	VkStructureTypeMemoryAllocateInfo                   VkStructureType = 5
	VkStructureTypeMappedMemoryRange                    VkStructureType = 6
	VkStructureTypeBindSparseInfo                       VkStructureType = 7
	VkStructureTypeFenceCreateInfo                      VkStructureType = 8
	VkStructureTypeSemaphoreCreateInfo                  VkStructureType = 9
	VkStructureTypeEventCreateInfo                      VkStructureType = 10
	VkStructureTypeQueryPoolCreateInfo                  VkStructureType = 11
	VkStructureTypeBufferCreateInfo                     VkStructureType = 12
	VkStructureTypeBufferViewCreateInfo                 VkStructureType = 13
	VkStructureTypeImageCreateInfo                      VkStructureType = 14
	VkStructureTypeImageViewCreateInfo                  VkStructureType = 15
	VkStructureTypeShaderModuleCreateInfo               VkStructureType = 16
	VkStructureTypePipelineCacheCreateInfo              VkStructureType = 17
	VkStructureTypePipelineShaderStageCreateInfo        VkStructureType = 18
	VkStructureTypePipelineVertexInputStateCreateInfo   VkStructureType = 19
	VkStructureTypePipelineInputAssemblyStateCreateInfo VkStructureType = 20
	VkStructureTypePipelineTessellationStateCreateInfo  VkStructureType = 21
	VkStructureTypePipelineViewportStateCreateInfo      VkStructureType = 22
	VkStructureTypePipelineRasterizationStateCreateInfo VkStructureType = 23
	VkStructureTypePipelineMultisampleStateCreateInfo   VkStructureType = 24
	VkStructureTypePipelineDepthStencilStateCreateInfo  VkStructureType = 25
	VkStructureTypePipelineColorBlendStateCreateInfo    VkStructureType = 26
	VkStructureTypePipelineDynamicStateCreateInfo       VkStructureType = 27
	VkStructureTypeGraphicsPipelineCreateInfo           VkStructureType = 28
	VkStructureTypeComputePipelineCreateInfo            VkStructureType = 29
	VkStructureTypePipelineLayoutCreateInfo             VkStructureType = 30
	VkStructureTypeSamplerCreateInfo                    VkStructureType = 31
	VkStructureTypeDescriptorSetLayoutCreateInfo        VkStructureType = 32
	VkStructureTypeDescriptorPoolCreateInfo             VkStructureType = 33
	VkStructureTypeDescriptorSetAllocateInfo            VkStructureType = 34
	VkStructureTypeWriteDescriptorSet                   VkStructureType = 35
	VkStructureTypeCopyDescriptorSet                    VkStructureType = 36
	VkStructureTypeFramebufferCreateInfo                VkStructureType = 37
	VkStructureTypeRenderPassCreateInfo                 VkStructureType = 38
	VkStructureTypeCommandPoolCreateInfo                VkStructureType = 39
	VkStructureTypeCommandBufferAllocateInfo            VkStructureType = 40
	VkStructureTypeCommandBufferInheritanceInfo         VkStructureType = 41
	VkStructureTypeCommandBufferBeginInfo               VkStructureType = 42
	VkStructureTypeRenderPassBeginInfo                  VkStructureType = 43
	VkStructureTypeBufferMemoryBarrier                  VkStructureType = 44
	VkStructureTypeImageMemoryBarrier                   VkStructureType = 45
	VkStructureTypeMemoryBarrier                        VkStructureType = 46
	VkStructureTypeLoaderInstanceCreateInfo             VkStructureType = 47
	VkStructureTypeLoaderDeviceCreateInfo               VkStructureType = 48
)

type VkSubpassContents int64

var (
	VkSubpassContentsInline                  VkSubpassContents = 0
	VkSubpassContentsSecondaryCommandBuffers VkSubpassContents = 1
)

type VkResult int64

var (
	VkSuccess                   VkResult = 0
	VkNotReady                  VkResult = 1
	VkTimeout                   VkResult = 2
	VkEventSet                  VkResult = 3
	VkEventReset                VkResult = 4
	VkIncomplete                VkResult = 5
	VkErrorOutOfHostMemory      VkResult = -1
	VkErrorOutOfDeviceMemory    VkResult = -2
	VkErrorInitializationFailed VkResult = -3
	VkErrorDeviceLost           VkResult = -4
	VkErrorMemoryMapFailed      VkResult = -5
	VkErrorLayerNotPresent      VkResult = -6
	VkErrorExtensionNotPresent  VkResult = -7
	VkErrorFeatureNotPresent    VkResult = -8
	VkErrorIncompatibleDriver   VkResult = -9
	VkErrorTooManyObjects       VkResult = -10
	VkErrorFormatNotSupported   VkResult = -11
	VkErrorFragmentedPool       VkResult = -12
	VkErrorUnknown              VkResult = -13
)

type VkDynamicState int64

var (
	VkDynamicStateViewport           VkDynamicState = 0
	VkDynamicStateScissor            VkDynamicState = 1
	VkDynamicStateLineWidth          VkDynamicState = 2
	VkDynamicStateDepthBias          VkDynamicState = 3
	VkDynamicStateBlendConstants     VkDynamicState = 4
	VkDynamicStateDepthBounds        VkDynamicState = 5
	VkDynamicStateStencilCompareMask VkDynamicState = 6
	VkDynamicStateStencilWriteMask   VkDynamicState = 7
	VkDynamicStateStencilReference   VkDynamicState = 8
)

type VkDescriptorUpdateTemplateType int64

var (
	VkDescriptorUpdateTemplateTypeDescriptorSet VkDescriptorUpdateTemplateType = 0
)

type VkObjectType int64

var (
	VkObjectTypeUnknown             VkObjectType = 0
	VkObjectTypeInstance            VkObjectType = 1
	VkObjectTypePhysicalDevice      VkObjectType = 2
	VkObjectTypeDevice              VkObjectType = 3
	VkObjectTypeQueue               VkObjectType = 4
	VkObjectTypeSemaphore           VkObjectType = 5
	VkObjectTypeCommandBuffer       VkObjectType = 6
	VkObjectTypeFence               VkObjectType = 7
	VkObjectTypeDeviceMemory        VkObjectType = 8
	VkObjectTypeBuffer              VkObjectType = 9
	VkObjectTypeImage               VkObjectType = 10
	VkObjectTypeEvent               VkObjectType = 11
	VkObjectTypeQueryPool           VkObjectType = 12
	VkObjectTypeBufferView          VkObjectType = 13
	VkObjectTypeImageView           VkObjectType = 14
	VkObjectTypeShaderModule        VkObjectType = 15
	VkObjectTypePipelineCache       VkObjectType = 16
	VkObjectTypePipelineLayout      VkObjectType = 17
	VkObjectTypeRenderPass          VkObjectType = 18
	VkObjectTypePipeline            VkObjectType = 19
	VkObjectTypeDescriptorSetLayout VkObjectType = 20
	VkObjectTypeSampler             VkObjectType = 21
	VkObjectTypeDescriptorPool      VkObjectType = 22
	VkObjectTypeDescriptorSet       VkObjectType = 23
	VkObjectTypeFramebuffer         VkObjectType = 24
	VkObjectTypeCommandPool         VkObjectType = 25
)

type VkRayTracingInvocationReorderModeEXT int64

var (
	VkRayTracingInvocationReorderModeNoneExt    VkRayTracingInvocationReorderModeEXT = 0
	VkRayTracingInvocationReorderModeReorderExt VkRayTracingInvocationReorderModeEXT = 1
)

type VkRayTracingLssIndexingModeNV int64

var (
	VkRayTracingLssIndexingModeListNv       VkRayTracingLssIndexingModeNV = 0
	VkRayTracingLssIndexingModeSuccessiveNv VkRayTracingLssIndexingModeNV = 1
)

type VkRayTracingLssPrimitiveEndCapsModeNV int64

var (
	VkRayTracingLssPrimitiveEndCapsModeNoneNv    VkRayTracingLssPrimitiveEndCapsModeNV = 0
	VkRayTracingLssPrimitiveEndCapsModeChainedNv VkRayTracingLssPrimitiveEndCapsModeNV = 1
)

type VkDirectDriverLoadingModeLUNARG int64

var (
	VkDirectDriverLoadingModeExclusiveLunarg VkDirectDriverLoadingModeLUNARG = 0
	VkDirectDriverLoadingModeInclusiveLunarg VkDirectDriverLoadingModeLUNARG = 1
)

type VkAntiLagModeAMD int64

var (
	VkAntiLagModeDriverControlAmd VkAntiLagModeAMD = 0
	VkAntiLagModeOnAmd            VkAntiLagModeAMD = 1
	VkAntiLagModeOffAmd           VkAntiLagModeAMD = 2
)

type VkAntiLagStageAMD int64

var (
	VkAntiLagStageInputAmd   VkAntiLagStageAMD = 0
	VkAntiLagStagePresentAmd VkAntiLagStageAMD = 1
)

type VkQueueFlagBits int64

var (
	VkQueueGraphicsBit      VkQueueFlagBits = (1 << 0)
	VkQueueComputeBit       VkQueueFlagBits = (1 << 1)
	VkQueueTransferBit      VkQueueFlagBits = (1 << 2)
	VkQueueSparseBindingBit VkQueueFlagBits = (1 << 3)
)

type VkCullModeFlagBits int64

var (
	VkCullModeNone         VkCullModeFlagBits = (1 << 0)
	VkCullModeFrontBit     VkCullModeFlagBits = (1 << 0)
	VkCullModeBackBit      VkCullModeFlagBits = (1 << 1)
	VkCullModeFrontAndBack VkCullModeFlagBits = (1 << 0)
)

type VkRenderPassCreateFlagBits int64

var ()

type VkDeviceQueueCreateFlagBits int64

var ()

type VkMemoryPropertyFlagBits int64

var (
	VkMemoryPropertyDeviceLocalBit     VkMemoryPropertyFlagBits = (1 << 0)
	VkMemoryPropertyHostVisibleBit     VkMemoryPropertyFlagBits = (1 << 1)
	VkMemoryPropertyHostCoherentBit    VkMemoryPropertyFlagBits = (1 << 2)
	VkMemoryPropertyHostCachedBit      VkMemoryPropertyFlagBits = (1 << 3)
	VkMemoryPropertyLazilyAllocatedBit VkMemoryPropertyFlagBits = (1 << 4)
)

type VkMemoryHeapFlagBits int64

var (
	VkMemoryHeapDeviceLocalBit VkMemoryHeapFlagBits = (1 << 0)
)

type VkAccessFlagBits int64

var (
	VkAccessIndirectCommandReadBit         VkAccessFlagBits = (1 << 0)
	VkAccessIndexReadBit                   VkAccessFlagBits = (1 << 1)
	VkAccessVertexAttributeReadBit         VkAccessFlagBits = (1 << 2)
	VkAccessUniformReadBit                 VkAccessFlagBits = (1 << 3)
	VkAccessInputAttachmentReadBit         VkAccessFlagBits = (1 << 4)
	VkAccessShaderReadBit                  VkAccessFlagBits = (1 << 5)
	VkAccessShaderWriteBit                 VkAccessFlagBits = (1 << 6)
	VkAccessColorAttachmentReadBit         VkAccessFlagBits = (1 << 7)
	VkAccessColorAttachmentWriteBit        VkAccessFlagBits = (1 << 8)
	VkAccessDepthStencilAttachmentReadBit  VkAccessFlagBits = (1 << 9)
	VkAccessDepthStencilAttachmentWriteBit VkAccessFlagBits = (1 << 10)
	VkAccessTransferReadBit                VkAccessFlagBits = (1 << 11)
	VkAccessTransferWriteBit               VkAccessFlagBits = (1 << 12)
	VkAccessHostReadBit                    VkAccessFlagBits = (1 << 13)
	VkAccessHostWriteBit                   VkAccessFlagBits = (1 << 14)
	VkAccessMemoryReadBit                  VkAccessFlagBits = (1 << 15)
	VkAccessMemoryWriteBit                 VkAccessFlagBits = (1 << 16)
)

type VkBufferUsageFlagBits int64

var (
	VkBufferUsageTransferSrcBit        VkBufferUsageFlagBits = (1 << 0)
	VkBufferUsageTransferDstBit        VkBufferUsageFlagBits = (1 << 1)
	VkBufferUsageUniformTexelBufferBit VkBufferUsageFlagBits = (1 << 2)
	VkBufferUsageStorageTexelBufferBit VkBufferUsageFlagBits = (1 << 3)
	VkBufferUsageUniformBufferBit      VkBufferUsageFlagBits = (1 << 4)
	VkBufferUsageStorageBufferBit      VkBufferUsageFlagBits = (1 << 5)
	VkBufferUsageIndexBufferBit        VkBufferUsageFlagBits = (1 << 6)
	VkBufferUsageVertexBufferBit       VkBufferUsageFlagBits = (1 << 7)
	VkBufferUsageIndirectBufferBit     VkBufferUsageFlagBits = (1 << 8)
)

type VkBufferUsageFlagBits2 int64

var (
	VkBufferUsage2TransferSrcBit        VkBufferUsageFlagBits2 = (1 << 0)
	VkBufferUsage2TransferDstBit        VkBufferUsageFlagBits2 = (1 << 1)
	VkBufferUsage2UniformTexelBufferBit VkBufferUsageFlagBits2 = (1 << 2)
	VkBufferUsage2StorageTexelBufferBit VkBufferUsageFlagBits2 = (1 << 3)
	VkBufferUsage2UniformBufferBit      VkBufferUsageFlagBits2 = (1 << 4)
	VkBufferUsage2StorageBufferBit      VkBufferUsageFlagBits2 = (1 << 5)
	VkBufferUsage2IndexBufferBit        VkBufferUsageFlagBits2 = (1 << 6)
	VkBufferUsage2VertexBufferBit       VkBufferUsageFlagBits2 = (1 << 7)
	VkBufferUsage2IndirectBufferBit     VkBufferUsageFlagBits2 = (1 << 8)
)

type VkBufferCreateFlagBits int64

var (
	VkBufferCreateSparseBindingBit   VkBufferCreateFlagBits = (1 << 0)
	VkBufferCreateSparseResidencyBit VkBufferCreateFlagBits = (1 << 1)
	VkBufferCreateSparseAliasedBit   VkBufferCreateFlagBits = (1 << 2)
)

type VkShaderStageFlagBits int64

var (
	VkShaderStageVertexBit                 VkShaderStageFlagBits = (1 << 0)
	VkShaderStageTessellationControlBit    VkShaderStageFlagBits = (1 << 1)
	VkShaderStageTessellationEvaluationBit VkShaderStageFlagBits = (1 << 2)
	VkShaderStageGeometryBit               VkShaderStageFlagBits = (1 << 3)
	VkShaderStageFragmentBit               VkShaderStageFlagBits = (1 << 4)
	VkShaderStageComputeBit                VkShaderStageFlagBits = (1 << 5)
	VkShaderStageAllGraphics               VkShaderStageFlagBits = (1 << 0)
	VkShaderStageAll                       VkShaderStageFlagBits = (1 << 0)
)

type VkImageUsageFlagBits int64

var (
	VkImageUsageTransferSrcBit            VkImageUsageFlagBits = (1 << 0)
	VkImageUsageTransferDstBit            VkImageUsageFlagBits = (1 << 1)
	VkImageUsageSampledBit                VkImageUsageFlagBits = (1 << 2)
	VkImageUsageStorageBit                VkImageUsageFlagBits = (1 << 3)
	VkImageUsageColorAttachmentBit        VkImageUsageFlagBits = (1 << 4)
	VkImageUsageDepthStencilAttachmentBit VkImageUsageFlagBits = (1 << 5)
	VkImageUsageTransientAttachmentBit    VkImageUsageFlagBits = (1 << 6)
	VkImageUsageInputAttachmentBit        VkImageUsageFlagBits = (1 << 7)
)

type VkImageCreateFlagBits int64

var (
	VkImageCreateSparseBindingBit   VkImageCreateFlagBits = (1 << 0)
	VkImageCreateSparseResidencyBit VkImageCreateFlagBits = (1 << 1)
	VkImageCreateSparseAliasedBit   VkImageCreateFlagBits = (1 << 2)
	VkImageCreateMutableFormatBit   VkImageCreateFlagBits = (1 << 3)
	VkImageCreateCubeCompatibleBit  VkImageCreateFlagBits = (1 << 4)
)

type VkImageViewCreateFlagBits int64

var ()

type VkSamplerCreateFlagBits int64

var ()

type VkPipelineCreateFlagBits int64

var (
	VkPipelineCreateDisableOptimizationBit VkPipelineCreateFlagBits = (1 << 0)
	VkPipelineCreateAllowDerivativesBit    VkPipelineCreateFlagBits = (1 << 1)
	VkPipelineCreateDerivativeBit          VkPipelineCreateFlagBits = (1 << 2)
)

type VkPipelineCreateFlagBits2 int64

var (
	VkPipelineCreate2DisableOptimizationBit           VkPipelineCreateFlagBits2 = (1 << 0)
	VkPipelineCreate2AllowDerivativesBit              VkPipelineCreateFlagBits2 = (1 << 1)
	VkPipelineCreate2DerivativeBit                    VkPipelineCreateFlagBits2 = (1 << 2)
	VkPipelineCreate2ViewIndexFromDeviceIndexBit      VkPipelineCreateFlagBits2 = (1 << 3)
	VkPipelineCreate2DispatchBaseBit                  VkPipelineCreateFlagBits2 = (1 << 4)
	VkPipelineCreate2FailOnPipelineCompileRequiredBit VkPipelineCreateFlagBits2 = (1 << 8)
	VkPipelineCreate2EarlyReturnOnFailureBit          VkPipelineCreateFlagBits2 = (1 << 9)
	VkPipelineCreate2NoProtectedAccessBit             VkPipelineCreateFlagBits2 = (1 << 27)
	VkPipelineCreate2ProtectedAccessOnlyBit           VkPipelineCreateFlagBits2 = (1 << 30)
)

type VkPipelineShaderStageCreateFlagBits int64

var ()

type VkColorComponentFlagBits int64

var (
	VkColorComponentRBit VkColorComponentFlagBits = (1 << 0)
	VkColorComponentGBit VkColorComponentFlagBits = (1 << 1)
	VkColorComponentBBit VkColorComponentFlagBits = (1 << 2)
	VkColorComponentABit VkColorComponentFlagBits = (1 << 3)
)

type VkFenceCreateFlagBits int64

var (
	VkFenceCreateSignaledBit VkFenceCreateFlagBits = (1 << 0)
)

type VkSemaphoreCreateFlagBits int64

var ()

type VkFormatFeatureFlagBits int64

var (
	VkFormatFeatureSampledImageBit             VkFormatFeatureFlagBits = (1 << 0)
	VkFormatFeatureStorageImageBit             VkFormatFeatureFlagBits = (1 << 1)
	VkFormatFeatureStorageImageAtomicBit       VkFormatFeatureFlagBits = (1 << 2)
	VkFormatFeatureUniformTexelBufferBit       VkFormatFeatureFlagBits = (1 << 3)
	VkFormatFeatureStorageTexelBufferBit       VkFormatFeatureFlagBits = (1 << 4)
	VkFormatFeatureStorageTexelBufferAtomicBit VkFormatFeatureFlagBits = (1 << 5)
	VkFormatFeatureVertexBufferBit             VkFormatFeatureFlagBits = (1 << 6)
	VkFormatFeatureColorAttachmentBit          VkFormatFeatureFlagBits = (1 << 7)
	VkFormatFeatureColorAttachmentBlendBit     VkFormatFeatureFlagBits = (1 << 8)
	VkFormatFeatureDepthStencilAttachmentBit   VkFormatFeatureFlagBits = (1 << 9)
	VkFormatFeatureBlitSrcBit                  VkFormatFeatureFlagBits = (1 << 10)
	VkFormatFeatureBlitDstBit                  VkFormatFeatureFlagBits = (1 << 11)
	VkFormatFeatureSampledImageFilterLinearBit VkFormatFeatureFlagBits = (1 << 12)
)

type VkQueryControlFlagBits int64

var (
	VkQueryControlPreciseBit VkQueryControlFlagBits = (1 << 0)
)

type VkQueryResultFlagBits int64

var (
	VkQueryResult64Bit               VkQueryResultFlagBits = (1 << 0)
	VkQueryResultWaitBit             VkQueryResultFlagBits = (1 << 1)
	VkQueryResultWithAvailabilityBit VkQueryResultFlagBits = (1 << 2)
	VkQueryResultPartialBit          VkQueryResultFlagBits = (1 << 3)
)

type VkCommandBufferUsageFlagBits int64

var (
	VkCommandBufferUsageOneTimeSubmitBit      VkCommandBufferUsageFlagBits = (1 << 0)
	VkCommandBufferUsageRenderPassContinueBit VkCommandBufferUsageFlagBits = (1 << 1)
	VkCommandBufferUsageSimultaneousUseBit    VkCommandBufferUsageFlagBits = (1 << 2)
)

type VkQueryPipelineStatisticFlagBits int64

var (
	VkQueryPipelineStatisticInputAssemblyVerticesBit                   VkQueryPipelineStatisticFlagBits = (1 << 0)
	VkQueryPipelineStatisticInputAssemblyPrimitivesBit                 VkQueryPipelineStatisticFlagBits = (1 << 1)
	VkQueryPipelineStatisticVertexShaderInvocationsBit                 VkQueryPipelineStatisticFlagBits = (1 << 2)
	VkQueryPipelineStatisticGeometryShaderInvocationsBit               VkQueryPipelineStatisticFlagBits = (1 << 3)
	VkQueryPipelineStatisticGeometryShaderPrimitivesBit                VkQueryPipelineStatisticFlagBits = (1 << 4)
	VkQueryPipelineStatisticClippingInvocationsBit                     VkQueryPipelineStatisticFlagBits = (1 << 5)
	VkQueryPipelineStatisticClippingPrimitivesBit                      VkQueryPipelineStatisticFlagBits = (1 << 6)
	VkQueryPipelineStatisticFragmentShaderInvocationsBit               VkQueryPipelineStatisticFlagBits = (1 << 7)
	VkQueryPipelineStatisticTessellationControlShaderPatchesBit        VkQueryPipelineStatisticFlagBits = (1 << 8)
	VkQueryPipelineStatisticTessellationEvaluationShaderInvocationsBit VkQueryPipelineStatisticFlagBits = (1 << 9)
	VkQueryPipelineStatisticComputeShaderInvocationsBit                VkQueryPipelineStatisticFlagBits = (1 << 10)
)

type VkMemoryMapFlagBits int64

var ()

type VkImageAspectFlagBits int64

var (
	VkImageAspectColorBit    VkImageAspectFlagBits = (1 << 0)
	VkImageAspectDepthBit    VkImageAspectFlagBits = (1 << 1)
	VkImageAspectStencilBit  VkImageAspectFlagBits = (1 << 2)
	VkImageAspectMetadataBit VkImageAspectFlagBits = (1 << 3)
)

type VkSparseImageFormatFlagBits int64

var (
	VkSparseImageFormatSingleMiptailBit        VkSparseImageFormatFlagBits = (1 << 0)
	VkSparseImageFormatAlignedMipSizeBit       VkSparseImageFormatFlagBits = (1 << 1)
	VkSparseImageFormatNonstandardBlockSizeBit VkSparseImageFormatFlagBits = (1 << 2)
)

type VkSparseMemoryBindFlagBits int64

var (
	VkSparseMemoryBindMetadataBit VkSparseMemoryBindFlagBits = (1 << 0)
)

type VkPipelineStageFlagBits int64

var (
	VkPipelineStageTopOfPipeBit                    VkPipelineStageFlagBits = (1 << 0)
	VkPipelineStageDrawIndirectBit                 VkPipelineStageFlagBits = (1 << 1)
	VkPipelineStageVertexInputBit                  VkPipelineStageFlagBits = (1 << 2)
	VkPipelineStageVertexShaderBit                 VkPipelineStageFlagBits = (1 << 3)
	VkPipelineStageTessellationControlShaderBit    VkPipelineStageFlagBits = (1 << 4)
	VkPipelineStageTessellationEvaluationShaderBit VkPipelineStageFlagBits = (1 << 5)
	VkPipelineStageGeometryShaderBit               VkPipelineStageFlagBits = (1 << 6)
	VkPipelineStageFragmentShaderBit               VkPipelineStageFlagBits = (1 << 7)
	VkPipelineStageEarlyFragmentTestsBit           VkPipelineStageFlagBits = (1 << 8)
	VkPipelineStageLateFragmentTestsBit            VkPipelineStageFlagBits = (1 << 9)
	VkPipelineStageColorAttachmentOutputBit        VkPipelineStageFlagBits = (1 << 10)
	VkPipelineStageComputeShaderBit                VkPipelineStageFlagBits = (1 << 11)
	VkPipelineStageTransferBit                     VkPipelineStageFlagBits = (1 << 12)
	VkPipelineStageBottomOfPipeBit                 VkPipelineStageFlagBits = (1 << 13)
	VkPipelineStageHostBit                         VkPipelineStageFlagBits = (1 << 14)
	VkPipelineStageAllGraphicsBit                  VkPipelineStageFlagBits = (1 << 15)
	VkPipelineStageAllCommandsBit                  VkPipelineStageFlagBits = (1 << 16)
)

type VkCommandPoolCreateFlagBits int64

var (
	VkCommandPoolCreateTransientBit          VkCommandPoolCreateFlagBits = (1 << 0)
	VkCommandPoolCreateResetCommandBufferBit VkCommandPoolCreateFlagBits = (1 << 1)
)

type VkCommandPoolResetFlagBits int64

var (
	VkCommandPoolResetReleaseResourcesBit VkCommandPoolResetFlagBits = (1 << 0)
)

type VkCommandBufferResetFlagBits int64

var (
	VkCommandBufferResetReleaseResourcesBit VkCommandBufferResetFlagBits = (1 << 0)
)

type VkSampleCountFlagBits int64

var (
	VkSampleCount1Bit  VkSampleCountFlagBits = (1 << 0)
	VkSampleCount2Bit  VkSampleCountFlagBits = (1 << 1)
	VkSampleCount4Bit  VkSampleCountFlagBits = (1 << 2)
	VkSampleCount8Bit  VkSampleCountFlagBits = (1 << 3)
	VkSampleCount16Bit VkSampleCountFlagBits = (1 << 4)
	VkSampleCount32Bit VkSampleCountFlagBits = (1 << 5)
	VkSampleCount64Bit VkSampleCountFlagBits = (1 << 6)
)

type VkAttachmentDescriptionFlagBits int64

var (
	VkAttachmentDescriptionMayAliasBit VkAttachmentDescriptionFlagBits = (1 << 0)
)

type VkStencilFaceFlagBits int64

var (
	VkStencilFaceFrontBit     VkStencilFaceFlagBits = (1 << 0)
	VkStencilFaceBackBit      VkStencilFaceFlagBits = (1 << 1)
	VkStencilFaceFrontAndBack VkStencilFaceFlagBits = (1 << 0)
	VkStencilFrontAndBack     VkStencilFaceFlagBits = (1 << 0)
)

type VkDescriptorPoolCreateFlagBits int64

var (
	VkDescriptorPoolCreateFreeDescriptorSetBit VkDescriptorPoolCreateFlagBits = (1 << 0)
)

type VkDependencyFlagBits int64

var (
	VkDependencyByRegionBit VkDependencyFlagBits = (1 << 0)
)

type VkSemaphoreType int64

var (
	VkSemaphoreTypeBinary   VkSemaphoreType = 0
	VkSemaphoreTypeTimeline VkSemaphoreType = 1
)

type VkSemaphoreWaitFlagBits int64

var (
	VkSemaphoreWaitAnyBit VkSemaphoreWaitFlagBits = (1 << 0)
)

type VkPresentModeKHR int64

var (
	VkPresentModeImmediateKhr   VkPresentModeKHR = 0
	VkPresentModeMailboxKhr     VkPresentModeKHR = 1
	VkPresentModeFifoKhr        VkPresentModeKHR = 2
	VkPresentModeFifoRelaxedKhr VkPresentModeKHR = 3
)

type VkColorSpaceKHR int64

var (
	VkColorSpaceSrgbNonlinearKhr VkColorSpaceKHR = 0
	VkColorspaceSrgbNonlinearKhr VkColorSpaceKHR = 0
)

type VkDisplayPlaneAlphaFlagBitsKHR int64

var (
	VkDisplayPlaneAlphaOpaqueBitKhr                VkDisplayPlaneAlphaFlagBitsKHR = (1 << 0)
	VkDisplayPlaneAlphaGlobalBitKhr                VkDisplayPlaneAlphaFlagBitsKHR = (1 << 1)
	VkDisplayPlaneAlphaPerPixelBitKhr              VkDisplayPlaneAlphaFlagBitsKHR = (1 << 2)
	VkDisplayPlaneAlphaPerPixelPremultipliedBitKhr VkDisplayPlaneAlphaFlagBitsKHR = (1 << 3)
)

type VkCompositeAlphaFlagBitsKHR int64

var (
	VkCompositeAlphaOpaqueBitKhr         VkCompositeAlphaFlagBitsKHR = (1 << 0)
	VkCompositeAlphaPreMultipliedBitKhr  VkCompositeAlphaFlagBitsKHR = (1 << 1)
	VkCompositeAlphaPostMultipliedBitKhr VkCompositeAlphaFlagBitsKHR = (1 << 2)
	VkCompositeAlphaInheritBitKhr        VkCompositeAlphaFlagBitsKHR = (1 << 3)
)

type VkSurfaceTransformFlagBitsKHR int64

var (
	VkSurfaceTransformIdentityBitKhr                  VkSurfaceTransformFlagBitsKHR = (1 << 0)
	VkSurfaceTransformRotate90BitKhr                  VkSurfaceTransformFlagBitsKHR = (1 << 1)
	VkSurfaceTransformRotate180BitKhr                 VkSurfaceTransformFlagBitsKHR = (1 << 2)
	VkSurfaceTransformRotate270BitKhr                 VkSurfaceTransformFlagBitsKHR = (1 << 3)
	VkSurfaceTransformHorizontalMirrorBitKhr          VkSurfaceTransformFlagBitsKHR = (1 << 4)
	VkSurfaceTransformHorizontalMirrorRotate90BitKhr  VkSurfaceTransformFlagBitsKHR = (1 << 5)
	VkSurfaceTransformHorizontalMirrorRotate180BitKhr VkSurfaceTransformFlagBitsKHR = (1 << 6)
	VkSurfaceTransformHorizontalMirrorRotate270BitKhr VkSurfaceTransformFlagBitsKHR = (1 << 7)
	VkSurfaceTransformInheritBitKhr                   VkSurfaceTransformFlagBitsKHR = (1 << 8)
)

type VkDisplaySurfaceStereoTypeNV int64

var (
	VkDisplaySurfaceStereoTypeNoneNv              VkDisplaySurfaceStereoTypeNV = 0
	VkDisplaySurfaceStereoTypeOnboardDinNv        VkDisplaySurfaceStereoTypeNV = 1
	VkDisplaySurfaceStereoTypeHdmi3dNv            VkDisplaySurfaceStereoTypeNV = 2
	VkDisplaySurfaceStereoTypeInbandDisplayportNv VkDisplaySurfaceStereoTypeNV = 3
)

type VkSwapchainImageUsageFlagBitsANDROID int64

var (
	VkSwapchainImageUsageSharedBitAndroid VkSwapchainImageUsageFlagBitsANDROID = (1 << 0)
)

type VkTimeDomainKHR int64

var (
	VkTimeDomainDeviceKhr                  VkTimeDomainKHR = 0
	VkTimeDomainClockMonotonicKhr          VkTimeDomainKHR = 1
	VkTimeDomainClockMonotonicRawKhr       VkTimeDomainKHR = 2
	VkTimeDomainQueryPerformanceCounterKhr VkTimeDomainKHR = 3
)

type VkDebugReportFlagBitsEXT int64

var (
	VkDebugReportInformationBitExt        VkDebugReportFlagBitsEXT = (1 << 0)
	VkDebugReportWarningBitExt            VkDebugReportFlagBitsEXT = (1 << 1)
	VkDebugReportPerformanceWarningBitExt VkDebugReportFlagBitsEXT = (1 << 2)
	VkDebugReportErrorBitExt              VkDebugReportFlagBitsEXT = (1 << 3)
	VkDebugReportDebugBitExt              VkDebugReportFlagBitsEXT = (1 << 4)
)

type VkDebugReportObjectTypeEXT int64

var (
	VkDebugReportObjectTypeUnknownExt                VkDebugReportObjectTypeEXT = 0
	VkDebugReportObjectTypeInstanceExt               VkDebugReportObjectTypeEXT = 1
	VkDebugReportObjectTypePhysicalDeviceExt         VkDebugReportObjectTypeEXT = 2
	VkDebugReportObjectTypeDeviceExt                 VkDebugReportObjectTypeEXT = 3
	VkDebugReportObjectTypeQueueExt                  VkDebugReportObjectTypeEXT = 4
	VkDebugReportObjectTypeSemaphoreExt              VkDebugReportObjectTypeEXT = 5
	VkDebugReportObjectTypeCommandBufferExt          VkDebugReportObjectTypeEXT = 6
	VkDebugReportObjectTypeFenceExt                  VkDebugReportObjectTypeEXT = 7
	VkDebugReportObjectTypeDeviceMemoryExt           VkDebugReportObjectTypeEXT = 8
	VkDebugReportObjectTypeBufferExt                 VkDebugReportObjectTypeEXT = 9
	VkDebugReportObjectTypeImageExt                  VkDebugReportObjectTypeEXT = 10
	VkDebugReportObjectTypeEventExt                  VkDebugReportObjectTypeEXT = 11
	VkDebugReportObjectTypeQueryPoolExt              VkDebugReportObjectTypeEXT = 12
	VkDebugReportObjectTypeBufferViewExt             VkDebugReportObjectTypeEXT = 13
	VkDebugReportObjectTypeImageViewExt              VkDebugReportObjectTypeEXT = 14
	VkDebugReportObjectTypeShaderModuleExt           VkDebugReportObjectTypeEXT = 15
	VkDebugReportObjectTypePipelineCacheExt          VkDebugReportObjectTypeEXT = 16
	VkDebugReportObjectTypePipelineLayoutExt         VkDebugReportObjectTypeEXT = 17
	VkDebugReportObjectTypeRenderPassExt             VkDebugReportObjectTypeEXT = 18
	VkDebugReportObjectTypePipelineExt               VkDebugReportObjectTypeEXT = 19
	VkDebugReportObjectTypeDescriptorSetLayoutExt    VkDebugReportObjectTypeEXT = 20
	VkDebugReportObjectTypeSamplerExt                VkDebugReportObjectTypeEXT = 21
	VkDebugReportObjectTypeDescriptorPoolExt         VkDebugReportObjectTypeEXT = 22
	VkDebugReportObjectTypeDescriptorSetExt          VkDebugReportObjectTypeEXT = 23
	VkDebugReportObjectTypeFramebufferExt            VkDebugReportObjectTypeEXT = 24
	VkDebugReportObjectTypeCommandPoolExt            VkDebugReportObjectTypeEXT = 25
	VkDebugReportObjectTypeSurfaceKhrExt             VkDebugReportObjectTypeEXT = 26
	VkDebugReportObjectTypeSwapchainKhrExt           VkDebugReportObjectTypeEXT = 27
	VkDebugReportObjectTypeDebugReportCallbackExtExt VkDebugReportObjectTypeEXT = 28
	VkDebugReportObjectTypeDebugReportExt            VkDebugReportObjectTypeEXT = 0
	VkDebugReportObjectTypeDisplayKhrExt             VkDebugReportObjectTypeEXT = 29
	VkDebugReportObjectTypeDisplayModeKhrExt         VkDebugReportObjectTypeEXT = 30
	VkDebugReportObjectTypeValidationCacheExtExt     VkDebugReportObjectTypeEXT = 33
	VkDebugReportObjectTypeValidationCacheExt        VkDebugReportObjectTypeEXT = 0
)

type VkDeviceMemoryReportEventTypeEXT int64

var (
	VkDeviceMemoryReportEventTypeAllocateExt         VkDeviceMemoryReportEventTypeEXT = 0
	VkDeviceMemoryReportEventTypeFreeExt             VkDeviceMemoryReportEventTypeEXT = 1
	VkDeviceMemoryReportEventTypeImportExt           VkDeviceMemoryReportEventTypeEXT = 2
	VkDeviceMemoryReportEventTypeUnimportExt         VkDeviceMemoryReportEventTypeEXT = 3
	VkDeviceMemoryReportEventTypeAllocationFailedExt VkDeviceMemoryReportEventTypeEXT = 4
)

type VkRasterizationOrderAMD int64

var (
	VkRasterizationOrderStrictAmd  VkRasterizationOrderAMD = 0
	VkRasterizationOrderRelaxedAmd VkRasterizationOrderAMD = 1
)

type VkExternalMemoryHandleTypeFlagBitsNV int64

var (
	VkExternalMemoryHandleTypeOpaqueWin32BitNv    VkExternalMemoryHandleTypeFlagBitsNV = (1 << 0)
	VkExternalMemoryHandleTypeOpaqueWin32KmtBitNv VkExternalMemoryHandleTypeFlagBitsNV = (1 << 1)
	VkExternalMemoryHandleTypeD3d11ImageBitNv     VkExternalMemoryHandleTypeFlagBitsNV = (1 << 2)
	VkExternalMemoryHandleTypeD3d11ImageKmtBitNv  VkExternalMemoryHandleTypeFlagBitsNV = (1 << 3)
)

type VkExternalMemoryFeatureFlagBitsNV int64

var (
	VkExternalMemoryFeatureDedicatedOnlyBitNv VkExternalMemoryFeatureFlagBitsNV = (1 << 0)
	VkExternalMemoryFeatureExportableBitNv    VkExternalMemoryFeatureFlagBitsNV = (1 << 1)
	VkExternalMemoryFeatureImportableBitNv    VkExternalMemoryFeatureFlagBitsNV = (1 << 2)
)

type VkClusterAccelerationStructureIndexFormatFlagBitsNV int64

var (
	VkClusterAccelerationStructureIndexFormat8bitNv  VkClusterAccelerationStructureIndexFormatFlagBitsNV = (1 << 0)
	VkClusterAccelerationStructureIndexFormat16bitNv VkClusterAccelerationStructureIndexFormatFlagBitsNV = (1 << 1)
	VkClusterAccelerationStructureIndexFormat32bitNv VkClusterAccelerationStructureIndexFormatFlagBitsNV = (1 << 2)
)

type VkClusterAccelerationStructureTypeNV int64

var (
	VkClusterAccelerationStructureTypeClustersBottomLevelNv     VkClusterAccelerationStructureTypeNV = 0
	VkClusterAccelerationStructureTypeTriangleClusterNv         VkClusterAccelerationStructureTypeNV = 1
	VkClusterAccelerationStructureTypeTriangleClusterTemplateNv VkClusterAccelerationStructureTypeNV = 2
)

type VkClusterAccelerationStructureOpTypeNV int64

var (
	VkClusterAccelerationStructureOpTypeMoveObjectsNv                  VkClusterAccelerationStructureOpTypeNV = 0
	VkClusterAccelerationStructureOpTypeBuildClustersBottomLevelNv     VkClusterAccelerationStructureOpTypeNV = 1
	VkClusterAccelerationStructureOpTypeBuildTriangleClusterNv         VkClusterAccelerationStructureOpTypeNV = 2
	VkClusterAccelerationStructureOpTypeBuildTriangleClusterTemplateNv VkClusterAccelerationStructureOpTypeNV = 3
	VkClusterAccelerationStructureOpTypeInstantiateTriangleClusterNv   VkClusterAccelerationStructureOpTypeNV = 4
	VkClusterAccelerationStructureOpTypeGetClusterTemplateIndicesNv    VkClusterAccelerationStructureOpTypeNV = 5
)

type VkClusterAccelerationStructureOpModeNV int64

var (
	VkClusterAccelerationStructureOpModeImplicitDestinationsNv VkClusterAccelerationStructureOpModeNV = 0
	VkClusterAccelerationStructureOpModeExplicitDestinationsNv VkClusterAccelerationStructureOpModeNV = 1
	VkClusterAccelerationStructureOpModeComputeSizesNv         VkClusterAccelerationStructureOpModeNV = 2
)

type VkClusterAccelerationStructureClusterFlagBitsNV int64

var (
	VkClusterAccelerationStructureClusterAllowDisableOpacityMicromapsNv VkClusterAccelerationStructureClusterFlagBitsNV = (1 << 0)
)

type VkClusterAccelerationStructureGeometryFlagBitsNV int64

var (
	VkClusterAccelerationStructureGeometryCullDisableBitNv                 VkClusterAccelerationStructureGeometryFlagBitsNV = (1 << 0)
	VkClusterAccelerationStructureGeometryNoDuplicateAnyhitInvocationBitNv VkClusterAccelerationStructureGeometryFlagBitsNV = (1 << 1)
	VkClusterAccelerationStructureGeometryOpaqueBitNv                      VkClusterAccelerationStructureGeometryFlagBitsNV = (1 << 2)
)

type VkClusterAccelerationStructureAddressResolutionFlagBitsNV int64

var (
	VkClusterAccelerationStructureAddressResolutionNoneNv                         VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 0)
	VkClusterAccelerationStructureAddressResolutionIndirectedDstImplicitDataBitNv VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 0)
	VkClusterAccelerationStructureAddressResolutionIndirectedScratchDataBitNv     VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 1)
	VkClusterAccelerationStructureAddressResolutionIndirectedDstAddressArrayBitNv VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 2)
	VkClusterAccelerationStructureAddressResolutionIndirectedDstSizesArrayBitNv   VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 3)
	VkClusterAccelerationStructureAddressResolutionIndirectedSrcInfosArrayBitNv   VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 4)
	VkClusterAccelerationStructureAddressResolutionIndirectedSrcInfosCountBitNv   VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 5)
)

type VkValidationCheckEXT int64

var (
	VkValidationCheckAllExt     VkValidationCheckEXT = 0
	VkValidationCheckShadersExt VkValidationCheckEXT = 1
)

type VkValidationFeatureEnableEXT int64

var (
	VkValidationFeatureEnableGpuAssistedExt                   VkValidationFeatureEnableEXT = 0
	VkValidationFeatureEnableGpuAssistedReserveBindingSlotExt VkValidationFeatureEnableEXT = 1
	VkValidationFeatureEnableBestPracticesExt                 VkValidationFeatureEnableEXT = 2
	VkValidationFeatureEnableDebugPrintfExt                   VkValidationFeatureEnableEXT = 3
	VkValidationFeatureEnableSynchronizationValidationExt     VkValidationFeatureEnableEXT = 4
)

type VkValidationFeatureDisableEXT int64

var (
	VkValidationFeatureDisableAllExt                   VkValidationFeatureDisableEXT = 0
	VkValidationFeatureDisableShadersExt               VkValidationFeatureDisableEXT = 1
	VkValidationFeatureDisableThreadSafetyExt          VkValidationFeatureDisableEXT = 2
	VkValidationFeatureDisableApiParametersExt         VkValidationFeatureDisableEXT = 3
	VkValidationFeatureDisableObjectLifetimesExt       VkValidationFeatureDisableEXT = 4
	VkValidationFeatureDisableCoreChecksExt            VkValidationFeatureDisableEXT = 5
	VkValidationFeatureDisableUniqueHandlesExt         VkValidationFeatureDisableEXT = 6
	VkValidationFeatureDisableShaderValidationCacheExt VkValidationFeatureDisableEXT = 7
)

type VkLayerSettingTypeEXT int64

var (
	VkLayerSettingTypeBool32Ext  VkLayerSettingTypeEXT = 0
	VkLayerSettingTypeInt32Ext   VkLayerSettingTypeEXT = 1
	VkLayerSettingTypeInt64Ext   VkLayerSettingTypeEXT = 2
	VkLayerSettingTypeUint32Ext  VkLayerSettingTypeEXT = 3
	VkLayerSettingTypeUint64Ext  VkLayerSettingTypeEXT = 4
	VkLayerSettingTypeFloat32Ext VkLayerSettingTypeEXT = 5
	VkLayerSettingTypeFloat64Ext VkLayerSettingTypeEXT = 6
	VkLayerSettingTypeStringExt  VkLayerSettingTypeEXT = 7
)

type VkSubgroupFeatureFlagBits int64

var (
	VkSubgroupFeatureBasicBit           VkSubgroupFeatureFlagBits = (1 << 0)
	VkSubgroupFeatureVoteBit            VkSubgroupFeatureFlagBits = (1 << 1)
	VkSubgroupFeatureArithmeticBit      VkSubgroupFeatureFlagBits = (1 << 2)
	VkSubgroupFeatureBallotBit          VkSubgroupFeatureFlagBits = (1 << 3)
	VkSubgroupFeatureShuffleBit         VkSubgroupFeatureFlagBits = (1 << 4)
	VkSubgroupFeatureShuffleRelativeBit VkSubgroupFeatureFlagBits = (1 << 5)
	VkSubgroupFeatureClusteredBit       VkSubgroupFeatureFlagBits = (1 << 6)
	VkSubgroupFeatureQuadBit            VkSubgroupFeatureFlagBits = (1 << 7)
)

type VkIndirectCommandsLayoutUsageFlagBitsNV int64

var (
	VkIndirectCommandsLayoutUsageExplicitPreprocessBitNv VkIndirectCommandsLayoutUsageFlagBitsNV = (1 << 0)
	VkIndirectCommandsLayoutUsageIndexedSequencesBitNv   VkIndirectCommandsLayoutUsageFlagBitsNV = (1 << 1)
	VkIndirectCommandsLayoutUsageUnorderedSequencesBitNv VkIndirectCommandsLayoutUsageFlagBitsNV = (1 << 2)
)

type VkIndirectStateFlagBitsNV int64

var (
	VkIndirectStateFlagFrontfaceBitNv VkIndirectStateFlagBitsNV = (1 << 0)
)

type VkIndirectCommandsTokenTypeNV int64

var (
	VkIndirectCommandsTokenTypeShaderGroupNv  VkIndirectCommandsTokenTypeNV = 0
	VkIndirectCommandsTokenTypeStateFlagsNv   VkIndirectCommandsTokenTypeNV = 1
	VkIndirectCommandsTokenTypeIndexBufferNv  VkIndirectCommandsTokenTypeNV = 2
	VkIndirectCommandsTokenTypeVertexBufferNv VkIndirectCommandsTokenTypeNV = 3
	VkIndirectCommandsTokenTypePushConstantNv VkIndirectCommandsTokenTypeNV = 4
	VkIndirectCommandsTokenTypeDrawIndexedNv  VkIndirectCommandsTokenTypeNV = 5
	VkIndirectCommandsTokenTypeDrawNv         VkIndirectCommandsTokenTypeNV = 6
	VkIndirectCommandsTokenTypeDrawTasksNv    VkIndirectCommandsTokenTypeNV = 7
)

type VkPrivateDataSlotCreateFlagBits int64

var ()

type VkDescriptorSetLayoutCreateFlagBits int64

var ()

type VkExternalMemoryHandleTypeFlagBits int64

var (
	VkExternalMemoryHandleTypeOpaqueFdBit        VkExternalMemoryHandleTypeFlagBits = (1 << 0)
	VkExternalMemoryHandleTypeOpaqueWin32Bit     VkExternalMemoryHandleTypeFlagBits = (1 << 1)
	VkExternalMemoryHandleTypeOpaqueWin32KmtBit  VkExternalMemoryHandleTypeFlagBits = (1 << 2)
	VkExternalMemoryHandleTypeD3d11TextureBit    VkExternalMemoryHandleTypeFlagBits = (1 << 3)
	VkExternalMemoryHandleTypeD3d11TextureKmtBit VkExternalMemoryHandleTypeFlagBits = (1 << 4)
	VkExternalMemoryHandleTypeD3d12HeapBit       VkExternalMemoryHandleTypeFlagBits = (1 << 5)
	VkExternalMemoryHandleTypeD3d12ResourceBit   VkExternalMemoryHandleTypeFlagBits = (1 << 6)
)

type VkExternalMemoryFeatureFlagBits int64

var (
	VkExternalMemoryFeatureDedicatedOnlyBit VkExternalMemoryFeatureFlagBits = (1 << 0)
	VkExternalMemoryFeatureExportableBit    VkExternalMemoryFeatureFlagBits = (1 << 1)
	VkExternalMemoryFeatureImportableBit    VkExternalMemoryFeatureFlagBits = (1 << 2)
)

type VkExternalSemaphoreHandleTypeFlagBits int64

var (
	VkExternalSemaphoreHandleTypeOpaqueFdBit       VkExternalSemaphoreHandleTypeFlagBits = (1 << 0)
	VkExternalSemaphoreHandleTypeOpaqueWin32Bit    VkExternalSemaphoreHandleTypeFlagBits = (1 << 1)
	VkExternalSemaphoreHandleTypeOpaqueWin32KmtBit VkExternalSemaphoreHandleTypeFlagBits = (1 << 2)
	VkExternalSemaphoreHandleTypeD3d12FenceBit     VkExternalSemaphoreHandleTypeFlagBits = (1 << 3)
	VkExternalSemaphoreHandleTypeD3d11FenceBit     VkExternalSemaphoreHandleTypeFlagBits = (1 << 0)
	VkExternalSemaphoreHandleTypeSyncFdBit         VkExternalSemaphoreHandleTypeFlagBits = (1 << 4)
)

type VkExternalSemaphoreFeatureFlagBits int64

var (
	VkExternalSemaphoreFeatureExportableBit VkExternalSemaphoreFeatureFlagBits = (1 << 0)
	VkExternalSemaphoreFeatureImportableBit VkExternalSemaphoreFeatureFlagBits = (1 << 1)
)

type VkSemaphoreImportFlagBits int64

var (
	VkSemaphoreImportTemporaryBit VkSemaphoreImportFlagBits = (1 << 0)
)

type VkExternalFenceHandleTypeFlagBits int64

var (
	VkExternalFenceHandleTypeOpaqueFdBit       VkExternalFenceHandleTypeFlagBits = (1 << 0)
	VkExternalFenceHandleTypeOpaqueWin32Bit    VkExternalFenceHandleTypeFlagBits = (1 << 1)
	VkExternalFenceHandleTypeOpaqueWin32KmtBit VkExternalFenceHandleTypeFlagBits = (1 << 2)
	VkExternalFenceHandleTypeSyncFdBit         VkExternalFenceHandleTypeFlagBits = (1 << 3)
)

type VkExternalFenceFeatureFlagBits int64

var (
	VkExternalFenceFeatureExportableBit VkExternalFenceFeatureFlagBits = (1 << 0)
	VkExternalFenceFeatureImportableBit VkExternalFenceFeatureFlagBits = (1 << 1)
)

type VkFenceImportFlagBits int64

var (
	VkFenceImportTemporaryBit VkFenceImportFlagBits = (1 << 0)
)

type VkSurfaceCounterFlagBitsEXT int64

var (
	VkSurfaceCounterVblankBitExt VkSurfaceCounterFlagBitsEXT = (1 << 0)
	VkSurfaceCounterVblankExt    VkSurfaceCounterFlagBitsEXT = (1 << 0)
)

type VkDisplayPowerStateEXT int64

var (
	VkDisplayPowerStateOffExt     VkDisplayPowerStateEXT = 0
	VkDisplayPowerStateSuspendExt VkDisplayPowerStateEXT = 1
	VkDisplayPowerStateOnExt      VkDisplayPowerStateEXT = 2
)

type VkDeviceEventTypeEXT int64

var (
	VkDeviceEventTypeDisplayHotplugExt VkDeviceEventTypeEXT = 0
)

type VkDisplayEventTypeEXT int64

var (
	VkDisplayEventTypeFirstPixelOutExt VkDisplayEventTypeEXT = 0
)

type VkPeerMemoryFeatureFlagBits int64

var (
	VkPeerMemoryFeatureCopySrcBit    VkPeerMemoryFeatureFlagBits = (1 << 0)
	VkPeerMemoryFeatureCopyDstBit    VkPeerMemoryFeatureFlagBits = (1 << 1)
	VkPeerMemoryFeatureGenericSrcBit VkPeerMemoryFeatureFlagBits = (1 << 2)
	VkPeerMemoryFeatureGenericDstBit VkPeerMemoryFeatureFlagBits = (1 << 3)
)

type VkMemoryAllocateFlagBits int64

var (
	VkMemoryAllocateDeviceMaskBit VkMemoryAllocateFlagBits = (1 << 0)
)

type VkDeviceGroupPresentModeFlagBitsKHR int64

var (
	VkDeviceGroupPresentModeLocalBitKhr            VkDeviceGroupPresentModeFlagBitsKHR = (1 << 0)
	VkDeviceGroupPresentModeRemoteBitKhr           VkDeviceGroupPresentModeFlagBitsKHR = (1 << 1)
	VkDeviceGroupPresentModeSumBitKhr              VkDeviceGroupPresentModeFlagBitsKHR = (1 << 2)
	VkDeviceGroupPresentModeLocalMultiDeviceBitKhr VkDeviceGroupPresentModeFlagBitsKHR = (1 << 3)
)

type VkSwapchainCreateFlagBitsKHR int64

var ()

type VkViewportCoordinateSwizzleNV int64

var (
	VkViewportCoordinateSwizzlePositiveXNv VkViewportCoordinateSwizzleNV = 0
	VkViewportCoordinateSwizzleNegativeXNv VkViewportCoordinateSwizzleNV = 1
	VkViewportCoordinateSwizzlePositiveYNv VkViewportCoordinateSwizzleNV = 2
	VkViewportCoordinateSwizzleNegativeYNv VkViewportCoordinateSwizzleNV = 3
	VkViewportCoordinateSwizzlePositiveZNv VkViewportCoordinateSwizzleNV = 4
	VkViewportCoordinateSwizzleNegativeZNv VkViewportCoordinateSwizzleNV = 5
	VkViewportCoordinateSwizzlePositiveWNv VkViewportCoordinateSwizzleNV = 6
	VkViewportCoordinateSwizzleNegativeWNv VkViewportCoordinateSwizzleNV = 7
)

type VkDiscardRectangleModeEXT int64

var (
	VkDiscardRectangleModeInclusiveExt VkDiscardRectangleModeEXT = 0
	VkDiscardRectangleModeExclusiveExt VkDiscardRectangleModeEXT = 1
)

type VkSubpassDescriptionFlagBits int64

var ()

type VkPointClippingBehavior int64

var (
	VkPointClippingBehaviorAllClipPlanes      VkPointClippingBehavior = 0
	VkPointClippingBehaviorUserClipPlanesOnly VkPointClippingBehavior = 1
)

type VkSamplerReductionMode int64

var (
	VkSamplerReductionModeWeightedAverage VkSamplerReductionMode = 0
	VkSamplerReductionModeMin             VkSamplerReductionMode = 1
	VkSamplerReductionModeMax             VkSamplerReductionMode = 2
)

type VkTessellationDomainOrigin int64

var (
	VkTessellationDomainOriginUpperLeft VkTessellationDomainOrigin = 0
	VkTessellationDomainOriginLowerLeft VkTessellationDomainOrigin = 1
)

type VkSamplerYcbcrModelConversion int64

var (
	VkSamplerYcbcrModelConversionRgbIdentity   VkSamplerYcbcrModelConversion = 0
	VkSamplerYcbcrModelConversionYcbcrIdentity VkSamplerYcbcrModelConversion = 1
	VkSamplerYcbcrModelConversionYcbcr709      VkSamplerYcbcrModelConversion = 2
	VkSamplerYcbcrModelConversionYcbcr601      VkSamplerYcbcrModelConversion = 3
	VkSamplerYcbcrModelConversionYcbcr2020     VkSamplerYcbcrModelConversion = 4
)

type VkSamplerYcbcrRange int64

var (
	VkSamplerYcbcrRangeItuFull   VkSamplerYcbcrRange = 0
	VkSamplerYcbcrRangeItuNarrow VkSamplerYcbcrRange = 1
)

type VkChromaLocation int64

var (
	VkChromaLocationCositedEven VkChromaLocation = 0
	VkChromaLocationMidpoint    VkChromaLocation = 1
)

type VkBlendOverlapEXT int64

var (
	VkBlendOverlapUncorrelatedExt VkBlendOverlapEXT = 0
	VkBlendOverlapDisjointExt     VkBlendOverlapEXT = 1
	VkBlendOverlapConjointExt     VkBlendOverlapEXT = 2
)

type VkCoverageModulationModeNV int64

var (
	VkCoverageModulationModeNoneNv  VkCoverageModulationModeNV = 0
	VkCoverageModulationModeRgbNv   VkCoverageModulationModeNV = 1
	VkCoverageModulationModeAlphaNv VkCoverageModulationModeNV = 2
	VkCoverageModulationModeRgbaNv  VkCoverageModulationModeNV = 3
)

type VkCoverageReductionModeNV int64

var (
	VkCoverageReductionModeMergeNv    VkCoverageReductionModeNV = 0
	VkCoverageReductionModeTruncateNv VkCoverageReductionModeNV = 1
)

type VkValidationCacheHeaderVersionEXT int64

var (
	VkValidationCacheHeaderVersionOneExt VkValidationCacheHeaderVersionEXT = 1
)

type VkShaderInfoTypeAMD int64

var (
	VkShaderInfoTypeStatisticsAmd  VkShaderInfoTypeAMD = 0
	VkShaderInfoTypeBinaryAmd      VkShaderInfoTypeAMD = 1
	VkShaderInfoTypeDisassemblyAmd VkShaderInfoTypeAMD = 2
)

type VkQueueGlobalPriority int64

var (
	VkQueueGlobalPriorityLow      VkQueueGlobalPriority = 128
	VkQueueGlobalPriorityMedium   VkQueueGlobalPriority = 256
	VkQueueGlobalPriorityHigh     VkQueueGlobalPriority = 512
	VkQueueGlobalPriorityRealtime VkQueueGlobalPriority = 1024
)

type VkDebugUtilsMessageSeverityFlagBitsEXT int64

var (
	VkDebugUtilsMessageSeverityVerboseBitExt VkDebugUtilsMessageSeverityFlagBitsEXT = (1 << 0)
	VkDebugUtilsMessageSeverityInfoBitExt    VkDebugUtilsMessageSeverityFlagBitsEXT = (1 << 4)
	VkDebugUtilsMessageSeverityWarningBitExt VkDebugUtilsMessageSeverityFlagBitsEXT = (1 << 8)
	VkDebugUtilsMessageSeverityErrorBitExt   VkDebugUtilsMessageSeverityFlagBitsEXT = (1 << 12)
)

type VkDebugUtilsMessageTypeFlagBitsEXT int64

var (
	VkDebugUtilsMessageTypeGeneralBitExt     VkDebugUtilsMessageTypeFlagBitsEXT = (1 << 0)
	VkDebugUtilsMessageTypeValidationBitExt  VkDebugUtilsMessageTypeFlagBitsEXT = (1 << 1)
	VkDebugUtilsMessageTypePerformanceBitExt VkDebugUtilsMessageTypeFlagBitsEXT = (1 << 2)
)

type VkConservativeRasterizationModeEXT int64

var (
	VkConservativeRasterizationModeDisabledExt      VkConservativeRasterizationModeEXT = 0
	VkConservativeRasterizationModeOverestimateExt  VkConservativeRasterizationModeEXT = 1
	VkConservativeRasterizationModeUnderestimateExt VkConservativeRasterizationModeEXT = 2
)

type VkDescriptorBindingFlagBits int64

var (
	VkDescriptorBindingUpdateAfterBindBit          VkDescriptorBindingFlagBits = (1 << 0)
	VkDescriptorBindingUpdateUnusedWhilePendingBit VkDescriptorBindingFlagBits = (1 << 1)
	VkDescriptorBindingPartiallyBoundBit           VkDescriptorBindingFlagBits = (1 << 2)
	VkDescriptorBindingVariableDescriptorCountBit  VkDescriptorBindingFlagBits = (1 << 3)
)

type VkVendorId int64

var (
	VkVendorIdKhronos  VkVendorId = 0x10000
	VkVendorIdViv      VkVendorId = 0x10001
	VkVendorIdVsi      VkVendorId = 0x10002
	VkVendorIdKazan    VkVendorId = 0x10003
	VkVendorIdCodeplay VkVendorId = 0x10004
	VkVendorIdMesa     VkVendorId = 0x10005
	VkVendorIdPocl     VkVendorId = 0x10006
	VkVendorIdMobileye VkVendorId = 0x10007
)

type VkDriverId int64

var (
	VkDriverIdAmdProprietary            VkDriverId = 1
	VkDriverIdAmdOpenSource             VkDriverId = 2
	VkDriverIdMesaRadv                  VkDriverId = 3
	VkDriverIdNvidiaProprietary         VkDriverId = 4
	VkDriverIdIntelProprietaryWindows   VkDriverId = 5
	VkDriverIdIntelOpenSourceMesa       VkDriverId = 6
	VkDriverIdImaginationProprietary    VkDriverId = 7
	VkDriverIdQualcommProprietary       VkDriverId = 8
	VkDriverIdArmProprietary            VkDriverId = 9
	VkDriverIdGoogleSwiftshader         VkDriverId = 10
	VkDriverIdGgpProprietary            VkDriverId = 11
	VkDriverIdBroadcomProprietary       VkDriverId = 12
	VkDriverIdMesaLlvmpipe              VkDriverId = 13
	VkDriverIdMoltenvk                  VkDriverId = 14
	VkDriverIdCoreaviProprietary        VkDriverId = 15
	VkDriverIdJuiceProprietary          VkDriverId = 16
	VkDriverIdVerisiliconProprietary    VkDriverId = 17
	VkDriverIdMesaTurnip                VkDriverId = 18
	VkDriverIdMesaV3dv                  VkDriverId = 19
	VkDriverIdMesaPanvk                 VkDriverId = 20
	VkDriverIdSamsungProprietary        VkDriverId = 21
	VkDriverIdMesaVenus                 VkDriverId = 22
	VkDriverIdMesaDozen                 VkDriverId = 23
	VkDriverIdMesaNvk                   VkDriverId = 24
	VkDriverIdImaginationOpenSourceMesa VkDriverId = 25
	VkDriverIdMesaHoneykrisp            VkDriverId = 26
	VkDriverIdVulkanScEmulationOnVulkan VkDriverId = 27
	VkDriverIdMesaKosmickrisp           VkDriverId = 28
)

type VkConditionalRenderingFlagBitsEXT int64

var (
	VkConditionalRenderingInvertedBitExt VkConditionalRenderingFlagBitsEXT = (1 << 0)
)

type VkResolveModeFlagBits int64

var (
	VkResolveModeNone          VkResolveModeFlagBits = (1 << 0)
	VkResolveModeSampleZeroBit VkResolveModeFlagBits = (1 << 0)
	VkResolveModeAverageBit    VkResolveModeFlagBits = (1 << 1)
	VkResolveModeMinBit        VkResolveModeFlagBits = (1 << 2)
	VkResolveModeMaxBit        VkResolveModeFlagBits = (1 << 3)
)

type VkShadingRatePaletteEntryNV int64

var (
	VkShadingRatePaletteEntryNoInvocationsNv           VkShadingRatePaletteEntryNV = 0
	VkShadingRatePaletteEntry16InvocationsPerPixelNv   VkShadingRatePaletteEntryNV = 1
	VkShadingRatePaletteEntry8InvocationsPerPixelNv    VkShadingRatePaletteEntryNV = 2
	VkShadingRatePaletteEntry4InvocationsPerPixelNv    VkShadingRatePaletteEntryNV = 3
	VkShadingRatePaletteEntry2InvocationsPerPixelNv    VkShadingRatePaletteEntryNV = 4
	VkShadingRatePaletteEntry1InvocationPerPixelNv     VkShadingRatePaletteEntryNV = 5
	VkShadingRatePaletteEntry1InvocationPer2x1PixelsNv VkShadingRatePaletteEntryNV = 6
	VkShadingRatePaletteEntry1InvocationPer1x2PixelsNv VkShadingRatePaletteEntryNV = 7
	VkShadingRatePaletteEntry1InvocationPer2x2PixelsNv VkShadingRatePaletteEntryNV = 8
	VkShadingRatePaletteEntry1InvocationPer4x2PixelsNv VkShadingRatePaletteEntryNV = 9
	VkShadingRatePaletteEntry1InvocationPer2x4PixelsNv VkShadingRatePaletteEntryNV = 10
	VkShadingRatePaletteEntry1InvocationPer4x4PixelsNv VkShadingRatePaletteEntryNV = 11
)

type VkCoarseSampleOrderTypeNV int64

var (
	VkCoarseSampleOrderTypeDefaultNv     VkCoarseSampleOrderTypeNV = 0
	VkCoarseSampleOrderTypeCustomNv      VkCoarseSampleOrderTypeNV = 1
	VkCoarseSampleOrderTypePixelMajorNv  VkCoarseSampleOrderTypeNV = 2
	VkCoarseSampleOrderTypeSampleMajorNv VkCoarseSampleOrderTypeNV = 3
)

type VkGeometryInstanceFlagBitsKHR int64

var (
	VkGeometryInstanceTriangleFacingCullDisableBitKhr     VkGeometryInstanceFlagBitsKHR = (1 << 0)
	VkGeometryInstanceTriangleFlipFacingBitKhr            VkGeometryInstanceFlagBitsKHR = (1 << 1)
	VkGeometryInstanceForceOpaqueBitKhr                   VkGeometryInstanceFlagBitsKHR = (1 << 2)
	VkGeometryInstanceForceNoOpaqueBitKhr                 VkGeometryInstanceFlagBitsKHR = (1 << 3)
	VkGeometryInstanceTriangleFrontCounterclockwiseBitKhr VkGeometryInstanceFlagBitsKHR = (1 << 0)
)

type VkGeometryFlagBitsKHR int64

var (
	VkGeometryOpaqueBitKhr                      VkGeometryFlagBitsKHR = (1 << 0)
	VkGeometryNoDuplicateAnyHitInvocationBitKhr VkGeometryFlagBitsKHR = (1 << 1)
)

type VkBuildAccelerationStructureFlagBitsKHR int64

var (
	VkBuildAccelerationStructureAllowUpdateBitKhr     VkBuildAccelerationStructureFlagBitsKHR = (1 << 0)
	VkBuildAccelerationStructureAllowCompactionBitKhr VkBuildAccelerationStructureFlagBitsKHR = (1 << 1)
	VkBuildAccelerationStructurePreferFastTraceBitKhr VkBuildAccelerationStructureFlagBitsKHR = (1 << 2)
	VkBuildAccelerationStructurePreferFastBuildBitKhr VkBuildAccelerationStructureFlagBitsKHR = (1 << 3)
	VkBuildAccelerationStructureLowMemoryBitKhr       VkBuildAccelerationStructureFlagBitsKHR = (1 << 4)
)

type VkAccelerationStructureCreateFlagBitsKHR int64

var (
	VkAccelerationStructureCreateDeviceAddressCaptureReplayBitKhr VkAccelerationStructureCreateFlagBitsKHR = (1 << 0)
)

type VkCopyAccelerationStructureModeKHR int64

var (
	VkCopyAccelerationStructureModeCloneKhr   VkCopyAccelerationStructureModeKHR = 0
	VkCopyAccelerationStructureModeCompactKhr VkCopyAccelerationStructureModeKHR = 1
)

type VkBuildAccelerationStructureModeKHR int64

var (
	VkBuildAccelerationStructureModeBuildKhr  VkBuildAccelerationStructureModeKHR = 0
	VkBuildAccelerationStructureModeUpdateKhr VkBuildAccelerationStructureModeKHR = 1
)

type VkAccelerationStructureTypeKHR int64

var (
	VkAccelerationStructureTypeTopLevelKhr    VkAccelerationStructureTypeKHR = 0
	VkAccelerationStructureTypeBottomLevelKhr VkAccelerationStructureTypeKHR = 1
	VkAccelerationStructureTypeGenericKhr     VkAccelerationStructureTypeKHR = 2
)

type VkGeometryTypeKHR int64

var (
	VkGeometryTypeTrianglesKhr VkGeometryTypeKHR = 0
	VkGeometryTypeAabbsKhr     VkGeometryTypeKHR = 1
	VkGeometryTypeInstancesKhr VkGeometryTypeKHR = 2
)

type VkAccelerationStructureMemoryRequirementsTypeNV int64

var (
	VkAccelerationStructureMemoryRequirementsTypeObjectNv        VkAccelerationStructureMemoryRequirementsTypeNV = 0
	VkAccelerationStructureMemoryRequirementsTypeBuildScratchNv  VkAccelerationStructureMemoryRequirementsTypeNV = 1
	VkAccelerationStructureMemoryRequirementsTypeUpdateScratchNv VkAccelerationStructureMemoryRequirementsTypeNV = 2
)

type VkAccelerationStructureBuildTypeKHR int64

var (
	VkAccelerationStructureBuildTypeHostKhr         VkAccelerationStructureBuildTypeKHR = 0
	VkAccelerationStructureBuildTypeDeviceKhr       VkAccelerationStructureBuildTypeKHR = 1
	VkAccelerationStructureBuildTypeHostOrDeviceKhr VkAccelerationStructureBuildTypeKHR = 2
)

type VkRayTracingShaderGroupTypeKHR int64

var (
	VkRayTracingShaderGroupTypeGeneralKhr            VkRayTracingShaderGroupTypeKHR = 0
	VkRayTracingShaderGroupTypeTrianglesHitGroupKhr  VkRayTracingShaderGroupTypeKHR = 1
	VkRayTracingShaderGroupTypeProceduralHitGroupKhr VkRayTracingShaderGroupTypeKHR = 2
)

type VkAccelerationStructureCompatibilityKHR int64

var (
	VkAccelerationStructureCompatibilityCompatibleKhr   VkAccelerationStructureCompatibilityKHR = 0
	VkAccelerationStructureCompatibilityIncompatibleKhr VkAccelerationStructureCompatibilityKHR = 1
)

type VkShaderGroupShaderKHR int64

var (
	VkShaderGroupShaderGeneralKhr      VkShaderGroupShaderKHR = 0
	VkShaderGroupShaderClosestHitKhr   VkShaderGroupShaderKHR = 1
	VkShaderGroupShaderAnyHitKhr       VkShaderGroupShaderKHR = 2
	VkShaderGroupShaderIntersectionKhr VkShaderGroupShaderKHR = 3
)

type VkMemoryOverallocationBehaviorAMD int64

var (
	VkMemoryOverallocationBehaviorDefaultAmd    VkMemoryOverallocationBehaviorAMD = 0
	VkMemoryOverallocationBehaviorAllowedAmd    VkMemoryOverallocationBehaviorAMD = 1
	VkMemoryOverallocationBehaviorDisallowedAmd VkMemoryOverallocationBehaviorAMD = 2
)

type VkFramebufferCreateFlagBits int64

var ()

type VkQueryPoolCreateFlagBits int64

var ()

type VkDeviceDiagnosticsConfigFlagBitsNV int64

var (
	VkDeviceDiagnosticsConfigEnableShaderDebugInfoBitNv      VkDeviceDiagnosticsConfigFlagBitsNV = (1 << 0)
	VkDeviceDiagnosticsConfigEnableResourceTrackingBitNv     VkDeviceDiagnosticsConfigFlagBitsNV = (1 << 1)
	VkDeviceDiagnosticsConfigEnableAutomaticCheckpointsBitNv VkDeviceDiagnosticsConfigFlagBitsNV = (1 << 2)
	VkDeviceDiagnosticsConfigEnableShaderErrorReportingBitNv VkDeviceDiagnosticsConfigFlagBitsNV = (1 << 3)
)

type VkPipelineCreationFeedbackFlagBits int64

var (
	VkPipelineCreationFeedbackValidBit                       VkPipelineCreationFeedbackFlagBits = (1 << 0)
	VkPipelineCreationFeedbackApplicationPipelineCacheHitBit VkPipelineCreationFeedbackFlagBits = (1 << 1)
	VkPipelineCreationFeedbackBasePipelineAccelerationBit    VkPipelineCreationFeedbackFlagBits = (1 << 2)
)

type VkFullScreenExclusiveEXT int64

var (
	VkFullScreenExclusiveDefaultExt               VkFullScreenExclusiveEXT = 0
	VkFullScreenExclusiveAllowedExt               VkFullScreenExclusiveEXT = 1
	VkFullScreenExclusiveDisallowedExt            VkFullScreenExclusiveEXT = 2
	VkFullScreenExclusiveApplicationControlledExt VkFullScreenExclusiveEXT = 3
)

type VkPerformanceCounterScopeKHR int64

var (
	VkPerformanceCounterScopeCommandBufferKhr VkPerformanceCounterScopeKHR = 0
	VkPerformanceCounterScopeRenderPassKhr    VkPerformanceCounterScopeKHR = 1
	VkPerformanceCounterScopeCommandKhr       VkPerformanceCounterScopeKHR = 2
	VkQueryScopeCommandBufferKhr              VkPerformanceCounterScopeKHR = 0
	VkQueryScopeRenderPassKhr                 VkPerformanceCounterScopeKHR = 0
	VkQueryScopeCommandKhr                    VkPerformanceCounterScopeKHR = 0
)

type VkMemoryDecompressionMethodFlagBitsEXT int64

var (
	VkMemoryDecompressionMethodGdeflate10BitExt VkMemoryDecompressionMethodFlagBitsEXT = (1 << 0)
	VkMemoryDecompressionMethodGdeflate10BitNv  VkMemoryDecompressionMethodFlagBitsEXT = (1 << 0)
)

type VkPerformanceCounterUnitKHR int64

var (
	VkPerformanceCounterUnitGenericKhr        VkPerformanceCounterUnitKHR = 0
	VkPerformanceCounterUnitPercentageKhr     VkPerformanceCounterUnitKHR = 1
	VkPerformanceCounterUnitNanosecondsKhr    VkPerformanceCounterUnitKHR = 2
	VkPerformanceCounterUnitBytesKhr          VkPerformanceCounterUnitKHR = 3
	VkPerformanceCounterUnitBytesPerSecondKhr VkPerformanceCounterUnitKHR = 4
	VkPerformanceCounterUnitKelvinKhr         VkPerformanceCounterUnitKHR = 5
	VkPerformanceCounterUnitWattsKhr          VkPerformanceCounterUnitKHR = 6
	VkPerformanceCounterUnitVoltsKhr          VkPerformanceCounterUnitKHR = 7
	VkPerformanceCounterUnitAmpsKhr           VkPerformanceCounterUnitKHR = 8
	VkPerformanceCounterUnitHertzKhr          VkPerformanceCounterUnitKHR = 9
	VkPerformanceCounterUnitCyclesKhr         VkPerformanceCounterUnitKHR = 10
)

type VkPerformanceCounterStorageKHR int64

var (
	VkPerformanceCounterStorageInt32Khr   VkPerformanceCounterStorageKHR = 0
	VkPerformanceCounterStorageInt64Khr   VkPerformanceCounterStorageKHR = 1
	VkPerformanceCounterStorageUint32Khr  VkPerformanceCounterStorageKHR = 2
	VkPerformanceCounterStorageUint64Khr  VkPerformanceCounterStorageKHR = 3
	VkPerformanceCounterStorageFloat32Khr VkPerformanceCounterStorageKHR = 4
	VkPerformanceCounterStorageFloat64Khr VkPerformanceCounterStorageKHR = 5
)

type VkPerformanceCounterDescriptionFlagBitsKHR int64

var (
	VkPerformanceCounterDescriptionPerformanceImpactingBitKhr VkPerformanceCounterDescriptionFlagBitsKHR = (1 << 0)
	VkPerformanceCounterDescriptionPerformanceImpactingKhr    VkPerformanceCounterDescriptionFlagBitsKHR = (1 << 0)
	VkPerformanceCounterDescriptionConcurrentlyImpactedBitKhr VkPerformanceCounterDescriptionFlagBitsKHR = (1 << 1)
	VkPerformanceCounterDescriptionConcurrentlyImpactedKhr    VkPerformanceCounterDescriptionFlagBitsKHR = (1 << 0)
)

type VkAcquireProfilingLockFlagBitsKHR int64

var ()

type VkShaderCorePropertiesFlagBitsAMD int64

var ()

type VkRefreshObjectFlagBitsKHR int64

var ()

type VkPerformanceConfigurationTypeINTEL int64

var (
	VkPerformanceConfigurationTypeCommandQueueMetricsDiscoveryActivatedIntel VkPerformanceConfigurationTypeINTEL = 0
)

type VkQueryPoolSamplingModeINTEL int64

var (
	VkQueryPoolSamplingModeManualIntel VkQueryPoolSamplingModeINTEL = 0
)

type VkPerformanceOverrideTypeINTEL int64

var (
	VkPerformanceOverrideTypeNullHardwareIntel   VkPerformanceOverrideTypeINTEL = 0
	VkPerformanceOverrideTypeFlushGpuCachesIntel VkPerformanceOverrideTypeINTEL = 1
)

type VkPerformanceParameterTypeINTEL int64

var (
	VkPerformanceParameterTypeHwCountersSupportedIntel   VkPerformanceParameterTypeINTEL = 0
	VkPerformanceParameterTypeStreamMarkerValidBitsIntel VkPerformanceParameterTypeINTEL = 1
)

type VkPerformanceValueTypeINTEL int64

var (
	VkPerformanceValueTypeUint32Intel VkPerformanceValueTypeINTEL = 0
	VkPerformanceValueTypeUint64Intel VkPerformanceValueTypeINTEL = 1
	VkPerformanceValueTypeFloatIntel  VkPerformanceValueTypeINTEL = 2
	VkPerformanceValueTypeBoolIntel   VkPerformanceValueTypeINTEL = 3
	VkPerformanceValueTypeStringIntel VkPerformanceValueTypeINTEL = 4
)

type VkShaderFloatControlsIndependence int64

var (
	VkShaderFloatControlsIndependence32BitOnly VkShaderFloatControlsIndependence = 0
	VkShaderFloatControlsIndependenceAll       VkShaderFloatControlsIndependence = 1
	VkShaderFloatControlsIndependenceNone      VkShaderFloatControlsIndependence = 2
)

type VkPipelineExecutableStatisticFormatKHR int64

var (
	VkPipelineExecutableStatisticFormatBool32Khr  VkPipelineExecutableStatisticFormatKHR = 0
	VkPipelineExecutableStatisticFormatInt64Khr   VkPipelineExecutableStatisticFormatKHR = 1
	VkPipelineExecutableStatisticFormatUint64Khr  VkPipelineExecutableStatisticFormatKHR = 2
	VkPipelineExecutableStatisticFormatFloat64Khr VkPipelineExecutableStatisticFormatKHR = 3
)

type VkLineRasterizationMode int64

var (
	VkLineRasterizationModeDefault           VkLineRasterizationMode = 0
	VkLineRasterizationModeRectangular       VkLineRasterizationMode = 1
	VkLineRasterizationModeBresenham         VkLineRasterizationMode = 2
	VkLineRasterizationModeRectangularSmooth VkLineRasterizationMode = 3
)

type VkShaderModuleCreateFlagBits int64

var ()

type VkPipelineCompilerControlFlagBitsAMD int64

var ()

type VkFaultLevel int64

var (
	VkFaultLevelUnassigned  VkFaultLevel = 0
	VkFaultLevelCritical    VkFaultLevel = 1
	VkFaultLevelRecoverable VkFaultLevel = 2
	VkFaultLevelWarning     VkFaultLevel = 3
)

type VkFaultType int64

var (
	VkFaultTypeInvalid           VkFaultType = 0
	VkFaultTypeUnassigned        VkFaultType = 1
	VkFaultTypeImplementation    VkFaultType = 2
	VkFaultTypeSystem            VkFaultType = 3
	VkFaultTypePhysicalDevice    VkFaultType = 4
	VkFaultTypeCommandBufferFull VkFaultType = 5
	VkFaultTypeInvalidApiUsage   VkFaultType = 6
)

type VkFaultQueryBehavior int64

var (
	VkFaultQueryBehaviorGetAndClearAllFaults VkFaultQueryBehavior = 0
)

type VkToolPurposeFlagBits int64

var (
	VkToolPurposeValidationBit         VkToolPurposeFlagBits = (1 << 0)
	VkToolPurposeProfilingBit          VkToolPurposeFlagBits = (1 << 1)
	VkToolPurposeTracingBit            VkToolPurposeFlagBits = (1 << 2)
	VkToolPurposeAdditionalFeaturesBit VkToolPurposeFlagBits = (1 << 3)
	VkToolPurposeModifyingFeaturesBit  VkToolPurposeFlagBits = (1 << 4)
)

type VkPipelineMatchControl int64

var (
	VkPipelineMatchControlApplicationUuidExactMatch VkPipelineMatchControl = 0
)

type VkFragmentShadingRateCombinerOpKHR int64

var (
	VkFragmentShadingRateCombinerOpKeepKhr    VkFragmentShadingRateCombinerOpKHR = 0
	VkFragmentShadingRateCombinerOpReplaceKhr VkFragmentShadingRateCombinerOpKHR = 1
	VkFragmentShadingRateCombinerOpMinKhr     VkFragmentShadingRateCombinerOpKHR = 2
	VkFragmentShadingRateCombinerOpMaxKhr     VkFragmentShadingRateCombinerOpKHR = 3
	VkFragmentShadingRateCombinerOpMulKhr     VkFragmentShadingRateCombinerOpKHR = 4
)

type VkFragmentShadingRateNV int64

var (
	VkFragmentShadingRate1InvocationPerPixelNv     VkFragmentShadingRateNV = 0
	VkFragmentShadingRate1InvocationPer1x2PixelsNv VkFragmentShadingRateNV = 1
	VkFragmentShadingRate1InvocationPer2x1PixelsNv VkFragmentShadingRateNV = 4
	VkFragmentShadingRate1InvocationPer2x2PixelsNv VkFragmentShadingRateNV = 5
	VkFragmentShadingRate1InvocationPer2x4PixelsNv VkFragmentShadingRateNV = 6
	VkFragmentShadingRate1InvocationPer4x2PixelsNv VkFragmentShadingRateNV = 9
	VkFragmentShadingRate1InvocationPer4x4PixelsNv VkFragmentShadingRateNV = 10
	VkFragmentShadingRate2InvocationsPerPixelNv    VkFragmentShadingRateNV = 11
	VkFragmentShadingRate4InvocationsPerPixelNv    VkFragmentShadingRateNV = 12
	VkFragmentShadingRate8InvocationsPerPixelNv    VkFragmentShadingRateNV = 13
	VkFragmentShadingRate16InvocationsPerPixelNv   VkFragmentShadingRateNV = 14
	VkFragmentShadingRateNoInvocationsNv           VkFragmentShadingRateNV = 15
)

type VkFragmentShadingRateTypeNV int64

var (
	VkFragmentShadingRateTypeFragmentSizeNv VkFragmentShadingRateTypeNV = 0
	VkFragmentShadingRateTypeEnumsNv        VkFragmentShadingRateTypeNV = 1
)

type VkSubpassMergeStatusEXT int64

var (
	VkSubpassMergeStatusMergedExt                               VkSubpassMergeStatusEXT = 0
	VkSubpassMergeStatusDisallowedExt                           VkSubpassMergeStatusEXT = 1
	VkSubpassMergeStatusNotMergedSideEffectsExt                 VkSubpassMergeStatusEXT = 2
	VkSubpassMergeStatusNotMergedSamplesMismatchExt             VkSubpassMergeStatusEXT = 3
	VkSubpassMergeStatusNotMergedViewsMismatchExt               VkSubpassMergeStatusEXT = 4
	VkSubpassMergeStatusNotMergedAliasingExt                    VkSubpassMergeStatusEXT = 5
	VkSubpassMergeStatusNotMergedDependenciesExt                VkSubpassMergeStatusEXT = 6
	VkSubpassMergeStatusNotMergedIncompatibleInputAttachmentExt VkSubpassMergeStatusEXT = 7
	VkSubpassMergeStatusNotMergedTooManyAttachmentsExt          VkSubpassMergeStatusEXT = 8
	VkSubpassMergeStatusNotMergedInsufficientStorageExt         VkSubpassMergeStatusEXT = 9
	VkSubpassMergeStatusNotMergedDepthStencilCountExt           VkSubpassMergeStatusEXT = 10
	VkSubpassMergeStatusNotMergedResolveAttachmentReuseExt      VkSubpassMergeStatusEXT = 11
	VkSubpassMergeStatusNotMergedSingleSubpassExt               VkSubpassMergeStatusEXT = 12
	VkSubpassMergeStatusNotMergedUnspecifiedExt                 VkSubpassMergeStatusEXT = 13
)

type VkAccessFlagBits2 int64

var (
	VkAccess2None                           VkAccessFlagBits2 = (1 << 0)
	VkAccess2IndirectCommandReadBit         VkAccessFlagBits2 = (1 << 0)
	VkAccess2IndexReadBit                   VkAccessFlagBits2 = (1 << 1)
	VkAccess2VertexAttributeReadBit         VkAccessFlagBits2 = (1 << 2)
	VkAccess2UniformReadBit                 VkAccessFlagBits2 = (1 << 3)
	VkAccess2InputAttachmentReadBit         VkAccessFlagBits2 = (1 << 4)
	VkAccess2ShaderReadBit                  VkAccessFlagBits2 = (1 << 5)
	VkAccess2ShaderWriteBit                 VkAccessFlagBits2 = (1 << 6)
	VkAccess2ColorAttachmentReadBit         VkAccessFlagBits2 = (1 << 7)
	VkAccess2ColorAttachmentWriteBit        VkAccessFlagBits2 = (1 << 8)
	VkAccess2DepthStencilAttachmentReadBit  VkAccessFlagBits2 = (1 << 9)
	VkAccess2DepthStencilAttachmentWriteBit VkAccessFlagBits2 = (1 << 10)
	VkAccess2TransferReadBit                VkAccessFlagBits2 = (1 << 11)
	VkAccess2TransferWriteBit               VkAccessFlagBits2 = (1 << 12)
	VkAccess2HostReadBit                    VkAccessFlagBits2 = (1 << 13)
	VkAccess2HostWriteBit                   VkAccessFlagBits2 = (1 << 14)
	VkAccess2MemoryReadBit                  VkAccessFlagBits2 = (1 << 15)
	VkAccess2MemoryWriteBit                 VkAccessFlagBits2 = (1 << 16)
	VkAccess2ShaderSampledReadBit           VkAccessFlagBits2 = (1 << 32)
	VkAccess2ShaderStorageReadBit           VkAccessFlagBits2 = (1 << 33)
	VkAccess2ShaderStorageWriteBit          VkAccessFlagBits2 = (1 << 34)
)

type VkPipelineStageFlagBits2 int64

var (
	VkPipelineStage2None                            VkPipelineStageFlagBits2 = (1 << 0)
	VkPipelineStage2TopOfPipeBit                    VkPipelineStageFlagBits2 = (1 << 0)
	VkPipelineStage2DrawIndirectBit                 VkPipelineStageFlagBits2 = (1 << 1)
	VkPipelineStage2VertexInputBit                  VkPipelineStageFlagBits2 = (1 << 2)
	VkPipelineStage2VertexShaderBit                 VkPipelineStageFlagBits2 = (1 << 3)
	VkPipelineStage2TessellationControlShaderBit    VkPipelineStageFlagBits2 = (1 << 4)
	VkPipelineStage2TessellationEvaluationShaderBit VkPipelineStageFlagBits2 = (1 << 5)
	VkPipelineStage2GeometryShaderBit               VkPipelineStageFlagBits2 = (1 << 6)
	VkPipelineStage2FragmentShaderBit               VkPipelineStageFlagBits2 = (1 << 7)
	VkPipelineStage2EarlyFragmentTestsBit           VkPipelineStageFlagBits2 = (1 << 8)
	VkPipelineStage2LateFragmentTestsBit            VkPipelineStageFlagBits2 = (1 << 9)
	VkPipelineStage2ColorAttachmentOutputBit        VkPipelineStageFlagBits2 = (1 << 10)
	VkPipelineStage2ComputeShaderBit                VkPipelineStageFlagBits2 = (1 << 11)
	VkPipelineStage2AllTransferBit                  VkPipelineStageFlagBits2 = (1 << 12)
	VkPipelineStage2TransferBit                     VkPipelineStageFlagBits2 = (1 << 0)
	VkPipelineStage2BottomOfPipeBit                 VkPipelineStageFlagBits2 = (1 << 13)
	VkPipelineStage2HostBit                         VkPipelineStageFlagBits2 = (1 << 14)
	VkPipelineStage2AllGraphicsBit                  VkPipelineStageFlagBits2 = (1 << 15)
	VkPipelineStage2AllCommandsBit                  VkPipelineStageFlagBits2 = (1 << 16)
	VkPipelineStage2CopyBit                         VkPipelineStageFlagBits2 = (1 << 32)
	VkPipelineStage2ResolveBit                      VkPipelineStageFlagBits2 = (1 << 33)
	VkPipelineStage2BlitBit                         VkPipelineStageFlagBits2 = (1 << 34)
	VkPipelineStage2ClearBit                        VkPipelineStageFlagBits2 = (1 << 35)
	VkPipelineStage2IndexInputBit                   VkPipelineStageFlagBits2 = (1 << 36)
	VkPipelineStage2VertexAttributeInputBit         VkPipelineStageFlagBits2 = (1 << 37)
	VkPipelineStage2PreRasterizationShadersBit      VkPipelineStageFlagBits2 = (1 << 38)
)

type VkSubmitFlagBits int64

var (
	VkSubmitProtectedBit VkSubmitFlagBits = (1 << 0)
)

type VkEventCreateFlagBits int64

var ()

type VkPipelineLayoutCreateFlagBits int64

var ()

type VkSciSyncClientTypeNV int64

var (
	VkSciSyncClientTypeSignalerNv       VkSciSyncClientTypeNV = 0
	VkSciSyncClientTypeWaiterNv         VkSciSyncClientTypeNV = 1
	VkSciSyncClientTypeSignalerWaiterNv VkSciSyncClientTypeNV = 2
)

type VkSciSyncPrimitiveTypeNV int64

var (
	VkSciSyncPrimitiveTypeFenceNv     VkSciSyncPrimitiveTypeNV = 0
	VkSciSyncPrimitiveTypeSemaphoreNv VkSciSyncPrimitiveTypeNV = 1
)

type VkProvokingVertexModeEXT int64

var (
	VkProvokingVertexModeFirstVertexExt VkProvokingVertexModeEXT = 0
	VkProvokingVertexModeLastVertexExt  VkProvokingVertexModeEXT = 1
)

type VkPipelineCacheValidationVersion int64

var (
	VkPipelineCacheValidationVersionSafetyCriticalOne VkPipelineCacheValidationVersion = 1
)

type VkAccelerationStructureMotionInstanceTypeNV int64

var (
	VkAccelerationStructureMotionInstanceTypeStaticNv       VkAccelerationStructureMotionInstanceTypeNV = 0
	VkAccelerationStructureMotionInstanceTypeMatrixMotionNv VkAccelerationStructureMotionInstanceTypeNV = 1
	VkAccelerationStructureMotionInstanceTypeSrtMotionNv    VkAccelerationStructureMotionInstanceTypeNV = 2
)

type VkPipelineColorBlendStateCreateFlagBits int64

var ()

type VkPipelineDepthStencilStateCreateFlagBits int64

var ()

type VkGraphicsPipelineLibraryFlagBitsEXT int64

var (
	VkGraphicsPipelineLibraryVertexInputInterfaceBitExt    VkGraphicsPipelineLibraryFlagBitsEXT = (1 << 0)
	VkGraphicsPipelineLibraryPreRasterizationShadersBitExt VkGraphicsPipelineLibraryFlagBitsEXT = (1 << 1)
	VkGraphicsPipelineLibraryFragmentShaderBitExt          VkGraphicsPipelineLibraryFlagBitsEXT = (1 << 2)
	VkGraphicsPipelineLibraryFragmentOutputInterfaceBitExt VkGraphicsPipelineLibraryFlagBitsEXT = (1 << 3)
)

type VkRenderingAttachmentFlagBitsKHR int64

var ()

type VkResolveImageFlagBitsKHR int64

var ()

type VkDeviceAddressBindingFlagBitsEXT int64

var (
	VkDeviceAddressBindingInternalObjectBitExt VkDeviceAddressBindingFlagBitsEXT = (1 << 0)
)

type VkDeviceAddressBindingTypeEXT int64

var (
	VkDeviceAddressBindingTypeBindExt   VkDeviceAddressBindingTypeEXT = 0
	VkDeviceAddressBindingTypeUnbindExt VkDeviceAddressBindingTypeEXT = 1
)

type VkFrameBoundaryFlagBitsEXT int64

var (
	VkFrameBoundaryFrameEndBitExt VkFrameBoundaryFlagBitsEXT = (1 << 0)
)

type VkPresentScalingFlagBitsKHR int64

var (
	VkPresentScalingOneToOneBitKhr           VkPresentScalingFlagBitsKHR = (1 << 0)
	VkPresentScalingOneToOneBitExt           VkPresentScalingFlagBitsKHR = (1 << 0)
	VkPresentScalingAspectRatioStretchBitKhr VkPresentScalingFlagBitsKHR = (1 << 1)
	VkPresentScalingAspectRatioStretchBitExt VkPresentScalingFlagBitsKHR = (1 << 0)
	VkPresentScalingStretchBitKhr            VkPresentScalingFlagBitsKHR = (1 << 2)
	VkPresentScalingStretchBitExt            VkPresentScalingFlagBitsKHR = (1 << 0)
)

type VkPresentGravityFlagBitsKHR int64

var (
	VkPresentGravityMinBitKhr      VkPresentGravityFlagBitsKHR = (1 << 0)
	VkPresentGravityMinBitExt      VkPresentGravityFlagBitsKHR = (1 << 0)
	VkPresentGravityMaxBitKhr      VkPresentGravityFlagBitsKHR = (1 << 1)
	VkPresentGravityMaxBitExt      VkPresentGravityFlagBitsKHR = (1 << 0)
	VkPresentGravityCenteredBitKhr VkPresentGravityFlagBitsKHR = (1 << 2)
	VkPresentGravityCenteredBitExt VkPresentGravityFlagBitsKHR = (1 << 0)
)

type VkPhysicalDeviceSchedulingControlsFlagBitsARM int64

var (
	VkPhysicalDeviceSchedulingControlsShaderCoreCountArm VkPhysicalDeviceSchedulingControlsFlagBitsARM = (1 << 0)
)

type VkPresentStageFlagBitsEXT int64

var (
	VkPresentStageQueueOperationsEndBitExt     VkPresentStageFlagBitsEXT = (1 << 0)
	VkPresentStageRequestDequeuedBitExt        VkPresentStageFlagBitsEXT = (1 << 1)
	VkPresentStageImageFirstPixelOutBitExt     VkPresentStageFlagBitsEXT = (1 << 2)
	VkPresentStageImageFirstPixelVisibleBitExt VkPresentStageFlagBitsEXT = (1 << 3)
)

type VkPastPresentationTimingFlagBitsEXT int64

var (
	VkPastPresentationTimingAllowPartialResultsBitExt    VkPastPresentationTimingFlagBitsEXT = (1 << 0)
	VkPastPresentationTimingAllowOutOfOrderResultsBitExt VkPastPresentationTimingFlagBitsEXT = (1 << 1)
)

type VkPresentTimingInfoFlagBitsEXT int64

var (
	VkPresentTimingInfoPresentAtRelativeTimeBitExt        VkPresentTimingInfoFlagBitsEXT = (1 << 0)
	VkPresentTimingInfoPresentAtNearestRefreshCycleBitExt VkPresentTimingInfoFlagBitsEXT = (1 << 1)
)

type VkVideoCodecOperationFlagBitsKHR int64

var (
	VkVideoCodecOperationNoneKhr VkVideoCodecOperationFlagBitsKHR = (1 << 0)
)

type VkVideoChromaSubsamplingFlagBitsKHR int64

var (
	VkVideoChromaSubsamplingInvalidKhr       VkVideoChromaSubsamplingFlagBitsKHR = (1 << 0)
	VkVideoChromaSubsamplingMonochromeBitKhr VkVideoChromaSubsamplingFlagBitsKHR = (1 << 0)
	VkVideoChromaSubsampling420BitKhr        VkVideoChromaSubsamplingFlagBitsKHR = (1 << 1)
	VkVideoChromaSubsampling422BitKhr        VkVideoChromaSubsamplingFlagBitsKHR = (1 << 2)
	VkVideoChromaSubsampling444BitKhr        VkVideoChromaSubsamplingFlagBitsKHR = (1 << 3)
)

type VkVideoComponentBitDepthFlagBitsKHR int64

var (
	VkVideoComponentBitDepthInvalidKhr VkVideoComponentBitDepthFlagBitsKHR = (1 << 0)
	VkVideoComponentBitDepth8BitKhr    VkVideoComponentBitDepthFlagBitsKHR = (1 << 0)
	VkVideoComponentBitDepth10BitKhr   VkVideoComponentBitDepthFlagBitsKHR = (1 << 2)
	VkVideoComponentBitDepth12BitKhr   VkVideoComponentBitDepthFlagBitsKHR = (1 << 4)
)

type VkVideoCapabilityFlagBitsKHR int64

var (
	VkVideoCapabilityProtectedContentBitKhr        VkVideoCapabilityFlagBitsKHR = (1 << 0)
	VkVideoCapabilitySeparateReferenceImagesBitKhr VkVideoCapabilityFlagBitsKHR = (1 << 1)
)

type VkVideoSessionCreateFlagBitsKHR int64

var (
	VkVideoSessionCreateProtectedContentBitKhr VkVideoSessionCreateFlagBitsKHR = (1 << 0)
)

type VkVideoSessionParametersCreateFlagBitsKHR int64

var ()

type VkVideoDecodeH264PictureLayoutFlagBitsKHR int64

var (
	VkVideoDecodeH264PictureLayoutProgressiveKhr                   VkVideoDecodeH264PictureLayoutFlagBitsKHR = (1 << 0)
	VkVideoDecodeH264PictureLayoutInterlacedInterleavedLinesBitKhr VkVideoDecodeH264PictureLayoutFlagBitsKHR = (1 << 0)
	VkVideoDecodeH264PictureLayoutInterlacedSeparatePlanesBitKhr   VkVideoDecodeH264PictureLayoutFlagBitsKHR = (1 << 1)
)

type VkVideoCodingControlFlagBitsKHR int64

var (
	VkVideoCodingControlResetBitKhr VkVideoCodingControlFlagBitsKHR = (1 << 0)
)

type VkQueryResultStatusKHR int64

var (
	VkQueryResultStatusErrorKhr    VkQueryResultStatusKHR = -1
	VkQueryResultStatusNotReadyKhr VkQueryResultStatusKHR = 0
	VkQueryResultStatusCompleteKhr VkQueryResultStatusKHR = 1
)

type VkVideoDecodeUsageFlagBitsKHR int64

var (
	VkVideoDecodeUsageDefaultKhr        VkVideoDecodeUsageFlagBitsKHR = (1 << 0)
	VkVideoDecodeUsageTranscodingBitKhr VkVideoDecodeUsageFlagBitsKHR = (1 << 0)
	VkVideoDecodeUsageOfflineBitKhr     VkVideoDecodeUsageFlagBitsKHR = (1 << 1)
	VkVideoDecodeUsageStreamingBitKhr   VkVideoDecodeUsageFlagBitsKHR = (1 << 2)
)

type VkVideoDecodeCapabilityFlagBitsKHR int64

var (
	VkVideoDecodeCapabilityDpbAndOutputCoincideBitKhr VkVideoDecodeCapabilityFlagBitsKHR = (1 << 0)
	VkVideoDecodeCapabilityDpbAndOutputDistinctBitKhr VkVideoDecodeCapabilityFlagBitsKHR = (1 << 1)
)

type VkVideoEncodeFlagBitsKHR int64

var ()

type VkVideoEncodeUsageFlagBitsKHR int64

var (
	VkVideoEncodeUsageDefaultKhr         VkVideoEncodeUsageFlagBitsKHR = (1 << 0)
	VkVideoEncodeUsageTranscodingBitKhr  VkVideoEncodeUsageFlagBitsKHR = (1 << 0)
	VkVideoEncodeUsageStreamingBitKhr    VkVideoEncodeUsageFlagBitsKHR = (1 << 1)
	VkVideoEncodeUsageRecordingBitKhr    VkVideoEncodeUsageFlagBitsKHR = (1 << 2)
	VkVideoEncodeUsageConferencingBitKhr VkVideoEncodeUsageFlagBitsKHR = (1 << 3)
)

type VkVideoEncodeContentFlagBitsKHR int64

var (
	VkVideoEncodeContentDefaultKhr     VkVideoEncodeContentFlagBitsKHR = (1 << 0)
	VkVideoEncodeContentCameraBitKhr   VkVideoEncodeContentFlagBitsKHR = (1 << 0)
	VkVideoEncodeContentDesktopBitKhr  VkVideoEncodeContentFlagBitsKHR = (1 << 1)
	VkVideoEncodeContentRenderedBitKhr VkVideoEncodeContentFlagBitsKHR = (1 << 2)
)

type VkVideoEncodeTuningModeKHR int64

var (
	VkVideoEncodeTuningModeDefaultKhr         VkVideoEncodeTuningModeKHR = 0
	VkVideoEncodeTuningModeHighQualityKhr     VkVideoEncodeTuningModeKHR = 1
	VkVideoEncodeTuningModeLowLatencyKhr      VkVideoEncodeTuningModeKHR = 2
	VkVideoEncodeTuningModeUltraLowLatencyKhr VkVideoEncodeTuningModeKHR = 3
	VkVideoEncodeTuningModeLosslessKhr        VkVideoEncodeTuningModeKHR = 4
)

type VkVideoEncodeCapabilityFlagBitsKHR int64

var (
	VkVideoEncodeCapabilityPrecedingExternallyEncodedBytesBitKhr           VkVideoEncodeCapabilityFlagBitsKHR = (1 << 0)
	VkVideoEncodeCapabilityInsufficientBitstreamBufferRangeDetectionBitKhr VkVideoEncodeCapabilityFlagBitsKHR = (1 << 1)
)

type VkVideoEncodeFeedbackFlagBitsKHR int64

var (
	VkVideoEncodeFeedbackBitstreamBufferOffsetBitKhr VkVideoEncodeFeedbackFlagBitsKHR = (1 << 0)
	VkVideoEncodeFeedbackBitstreamBytesWrittenBitKhr VkVideoEncodeFeedbackFlagBitsKHR = (1 << 1)
	VkVideoEncodeFeedbackBitstreamHasOverridesBitKhr VkVideoEncodeFeedbackFlagBitsKHR = (1 << 2)
)

type VkVideoEncodeRateControlModeFlagBitsKHR int64

var (
	VkVideoEncodeRateControlModeDefaultKhr     VkVideoEncodeRateControlModeFlagBitsKHR = (1 << 0)
	VkVideoEncodeRateControlModeDisabledBitKhr VkVideoEncodeRateControlModeFlagBitsKHR = (1 << 0)
	VkVideoEncodeRateControlModeCbrBitKhr      VkVideoEncodeRateControlModeFlagBitsKHR = (1 << 1)
	VkVideoEncodeRateControlModeVbrBitKhr      VkVideoEncodeRateControlModeFlagBitsKHR = (1 << 2)
)

type VkVideoEncodeIntraRefreshModeFlagBitsKHR int64

var (
	VkVideoEncodeIntraRefreshModeNoneKhr                   VkVideoEncodeIntraRefreshModeFlagBitsKHR = (1 << 0)
	VkVideoEncodeIntraRefreshModePerPicturePartitionBitKhr VkVideoEncodeIntraRefreshModeFlagBitsKHR = (1 << 0)
	VkVideoEncodeIntraRefreshModeBlockBasedBitKhr          VkVideoEncodeIntraRefreshModeFlagBitsKHR = (1 << 1)
	VkVideoEncodeIntraRefreshModeBlockRowBasedBitKhr       VkVideoEncodeIntraRefreshModeFlagBitsKHR = (1 << 2)
	VkVideoEncodeIntraRefreshModeBlockColumnBasedBitKhr    VkVideoEncodeIntraRefreshModeFlagBitsKHR = (1 << 3)
)

type VkVideoEncodeH264CapabilityFlagBitsKHR int64

var (
	VkVideoEncodeH264CapabilityHrdComplianceBitKhr                  VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 0)
	VkVideoEncodeH264CapabilityPredictionWeightTableGeneratedBitKhr VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 1)
	VkVideoEncodeH264CapabilityRowUnalignedSliceBitKhr              VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 2)
	VkVideoEncodeH264CapabilityDifferentSliceTypeBitKhr             VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 3)
	VkVideoEncodeH264CapabilityBFrameInL0ListBitKhr                 VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 4)
	VkVideoEncodeH264CapabilityBFrameInL1ListBitKhr                 VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 5)
	VkVideoEncodeH264CapabilityPerPictureTypeMinMaxQpBitKhr         VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 6)
	VkVideoEncodeH264CapabilityPerSliceConstantQpBitKhr             VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 7)
	VkVideoEncodeH264CapabilityGeneratePrefixNaluBitKhr             VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 8)
)

type VkVideoEncodeH264StdFlagBitsKHR int64

var (
	VkVideoEncodeH264StdSeparateColorPlaneFlagSetBitKhr          VkVideoEncodeH264StdFlagBitsKHR = (1 << 0)
	VkVideoEncodeH264StdQpprimeYZeroTransformBypassFlagSetBitKhr VkVideoEncodeH264StdFlagBitsKHR = (1 << 1)
	VkVideoEncodeH264StdScalingMatrixPresentFlagSetBitKhr        VkVideoEncodeH264StdFlagBitsKHR = (1 << 2)
	VkVideoEncodeH264StdChromaQpIndexOffsetBitKhr                VkVideoEncodeH264StdFlagBitsKHR = (1 << 3)
	VkVideoEncodeH264StdSecondChromaQpIndexOffsetBitKhr          VkVideoEncodeH264StdFlagBitsKHR = (1 << 4)
	VkVideoEncodeH264StdPicInitQpMinus26BitKhr                   VkVideoEncodeH264StdFlagBitsKHR = (1 << 5)
	VkVideoEncodeH264StdWeightedPredFlagSetBitKhr                VkVideoEncodeH264StdFlagBitsKHR = (1 << 6)
	VkVideoEncodeH264StdWeightedBipredIdcExplicitBitKhr          VkVideoEncodeH264StdFlagBitsKHR = (1 << 7)
	VkVideoEncodeH264StdWeightedBipredIdcImplicitBitKhr          VkVideoEncodeH264StdFlagBitsKHR = (1 << 8)
	VkVideoEncodeH264StdTransform8x8ModeFlagSetBitKhr            VkVideoEncodeH264StdFlagBitsKHR = (1 << 9)
	VkVideoEncodeH264StdDirectSpatialMvPredFlagUnsetBitKhr       VkVideoEncodeH264StdFlagBitsKHR = (1 << 10)
	VkVideoEncodeH264StdEntropyCodingModeFlagUnsetBitKhr         VkVideoEncodeH264StdFlagBitsKHR = (1 << 11)
	VkVideoEncodeH264StdEntropyCodingModeFlagSetBitKhr           VkVideoEncodeH264StdFlagBitsKHR = (1 << 12)
	VkVideoEncodeH264StdDirect8x8InferenceFlagUnsetBitKhr        VkVideoEncodeH264StdFlagBitsKHR = (1 << 13)
	VkVideoEncodeH264StdConstrainedIntraPredFlagSetBitKhr        VkVideoEncodeH264StdFlagBitsKHR = (1 << 14)
	VkVideoEncodeH264StdDeblockingFilterDisabledBitKhr           VkVideoEncodeH264StdFlagBitsKHR = (1 << 15)
	VkVideoEncodeH264StdDeblockingFilterEnabledBitKhr            VkVideoEncodeH264StdFlagBitsKHR = (1 << 16)
	VkVideoEncodeH264StdDeblockingFilterPartialBitKhr            VkVideoEncodeH264StdFlagBitsKHR = (1 << 17)
	VkVideoEncodeH264StdSliceQpDeltaBitKhr                       VkVideoEncodeH264StdFlagBitsKHR = (1 << 19)
	VkVideoEncodeH264StdDifferentSliceQpDeltaBitKhr              VkVideoEncodeH264StdFlagBitsKHR = (1 << 20)
)

type VkVideoEncodeH264RateControlFlagBitsKHR int64

var (
	VkVideoEncodeH264RateControlAttemptHrdComplianceBitKhr       VkVideoEncodeH264RateControlFlagBitsKHR = (1 << 0)
	VkVideoEncodeH264RateControlRegularGopBitKhr                 VkVideoEncodeH264RateControlFlagBitsKHR = (1 << 1)
	VkVideoEncodeH264RateControlReferencePatternFlatBitKhr       VkVideoEncodeH264RateControlFlagBitsKHR = (1 << 2)
	VkVideoEncodeH264RateControlReferencePatternDyadicBitKhr     VkVideoEncodeH264RateControlFlagBitsKHR = (1 << 3)
	VkVideoEncodeH264RateControlTemporalLayerPatternDyadicBitKhr VkVideoEncodeH264RateControlFlagBitsKHR = (1 << 4)
)

type VkHostImageCopyFlagBits int64

var (
	VkHostImageCopyMemcpyBit VkHostImageCopyFlagBits = (1 << 0)
	VkHostImageCopyMemcpy    VkHostImageCopyFlagBits = (1 << 0)
)

type VkPartitionedAccelerationStructureOpTypeNV int64

var (
	VkPartitionedAccelerationStructureOpTypeWriteInstanceNv             VkPartitionedAccelerationStructureOpTypeNV = 0
	VkPartitionedAccelerationStructureOpTypeUpdateInstanceNv            VkPartitionedAccelerationStructureOpTypeNV = 1
	VkPartitionedAccelerationStructureOpTypeWritePartitionTranslationNv VkPartitionedAccelerationStructureOpTypeNV = 2
)

type VkPartitionedAccelerationStructureInstanceFlagBitsNV int64

var (
	VkPartitionedAccelerationStructureInstanceFlagTriangleFacingCullDisableBitNv VkPartitionedAccelerationStructureInstanceFlagBitsNV = (1 << 0)
	VkPartitionedAccelerationStructureInstanceFlagTriangleFlipFacingBitNv        VkPartitionedAccelerationStructureInstanceFlagBitsNV = (1 << 1)
	VkPartitionedAccelerationStructureInstanceFlagForceOpaqueBitNv               VkPartitionedAccelerationStructureInstanceFlagBitsNV = (1 << 2)
	VkPartitionedAccelerationStructureInstanceFlagForceNoOpaqueBitNv             VkPartitionedAccelerationStructureInstanceFlagBitsNV = (1 << 3)
	VkPartitionedAccelerationStructureInstanceFlagEnableExplicitBoundingBoxNv    VkPartitionedAccelerationStructureInstanceFlagBitsNV = (1 << 4)
)

type VkImageFormatConstraintsFlagBitsFUCHSIA int64

var ()

type VkImageConstraintsInfoFlagBitsFUCHSIA int64

var (
	VkImageConstraintsInfoCpuReadRarelyFuchsia     VkImageConstraintsInfoFlagBitsFUCHSIA = (1 << 0)
	VkImageConstraintsInfoCpuReadOftenFuchsia      VkImageConstraintsInfoFlagBitsFUCHSIA = (1 << 1)
	VkImageConstraintsInfoCpuWriteRarelyFuchsia    VkImageConstraintsInfoFlagBitsFUCHSIA = (1 << 2)
	VkImageConstraintsInfoCpuWriteOftenFuchsia     VkImageConstraintsInfoFlagBitsFUCHSIA = (1 << 3)
	VkImageConstraintsInfoProtectedOptionalFuchsia VkImageConstraintsInfoFlagBitsFUCHSIA = (1 << 4)
)

type VkFormatFeatureFlagBits2 int64

var (
	VkFormatFeature2SampledImageBit                                                     VkFormatFeatureFlagBits2 = (1 << 0)
	VkFormatFeature2StorageImageBit                                                     VkFormatFeatureFlagBits2 = (1 << 1)
	VkFormatFeature2StorageImageAtomicBit                                               VkFormatFeatureFlagBits2 = (1 << 2)
	VkFormatFeature2UniformTexelBufferBit                                               VkFormatFeatureFlagBits2 = (1 << 3)
	VkFormatFeature2StorageTexelBufferBit                                               VkFormatFeatureFlagBits2 = (1 << 4)
	VkFormatFeature2StorageTexelBufferAtomicBit                                         VkFormatFeatureFlagBits2 = (1 << 5)
	VkFormatFeature2VertexBufferBit                                                     VkFormatFeatureFlagBits2 = (1 << 6)
	VkFormatFeature2ColorAttachmentBit                                                  VkFormatFeatureFlagBits2 = (1 << 7)
	VkFormatFeature2ColorAttachmentBlendBit                                             VkFormatFeatureFlagBits2 = (1 << 8)
	VkFormatFeature2DepthStencilAttachmentBit                                           VkFormatFeatureFlagBits2 = (1 << 9)
	VkFormatFeature2BlitSrcBit                                                          VkFormatFeatureFlagBits2 = (1 << 10)
	VkFormatFeature2BlitDstBit                                                          VkFormatFeatureFlagBits2 = (1 << 11)
	VkFormatFeature2SampledImageFilterLinearBit                                         VkFormatFeatureFlagBits2 = (1 << 12)
	VkFormatFeature2TransferSrcBit                                                      VkFormatFeatureFlagBits2 = (1 << 14)
	VkFormatFeature2TransferDstBit                                                      VkFormatFeatureFlagBits2 = (1 << 15)
	VkFormatFeature2SampledImageFilterMinmaxBit                                         VkFormatFeatureFlagBits2 = (1 << 16)
	VkFormatFeature2MidpointChromaSamplesBit                                            VkFormatFeatureFlagBits2 = (1 << 17)
	VkFormatFeature2SampledImageYcbcrConversionLinearFilterBit                          VkFormatFeatureFlagBits2 = (1 << 18)
	VkFormatFeature2SampledImageYcbcrConversionSeparateReconstructionFilterBit          VkFormatFeatureFlagBits2 = (1 << 19)
	VkFormatFeature2SampledImageYcbcrConversionChromaReconstructionExplicitBit          VkFormatFeatureFlagBits2 = (1 << 20)
	VkFormatFeature2SampledImageYcbcrConversionChromaReconstructionExplicitForceableBit VkFormatFeatureFlagBits2 = (1 << 21)
	VkFormatFeature2DisjointBit                                                         VkFormatFeatureFlagBits2 = (1 << 22)
	VkFormatFeature2CositedChromaSamplesBit                                             VkFormatFeatureFlagBits2 = (1 << 23)
	VkFormatFeature2StorageReadWithoutFormatBit                                         VkFormatFeatureFlagBits2 = (1 << 31)
	VkFormatFeature2StorageWriteWithoutFormatBit                                        VkFormatFeatureFlagBits2 = (1 << 32)
	VkFormatFeature2SampledImageDepthComparisonBit                                      VkFormatFeatureFlagBits2 = (1 << 33)
)

type VkRenderingFlagBits int64

var (
	VkRenderingContentsSecondaryCommandBuffersBit VkRenderingFlagBits = (1 << 0)
	VkRenderingSuspendingBit                      VkRenderingFlagBits = (1 << 1)
	VkRenderingResumingBit                        VkRenderingFlagBits = (1 << 2)
)

type VkVideoEncodeH265CapabilityFlagBitsKHR int64

var (
	VkVideoEncodeH265CapabilityHrdComplianceBitKhr                  VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 0)
	VkVideoEncodeH265CapabilityPredictionWeightTableGeneratedBitKhr VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 1)
	VkVideoEncodeH265CapabilityRowUnalignedSliceSegmentBitKhr       VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 2)
	VkVideoEncodeH265CapabilityDifferentSliceSegmentTypeBitKhr      VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 3)
	VkVideoEncodeH265CapabilityBFrameInL0ListBitKhr                 VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 4)
	VkVideoEncodeH265CapabilityBFrameInL1ListBitKhr                 VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 5)
	VkVideoEncodeH265CapabilityPerPictureTypeMinMaxQpBitKhr         VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 6)
	VkVideoEncodeH265CapabilityPerSliceSegmentConstantQpBitKhr      VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 7)
	VkVideoEncodeH265CapabilityMultipleTilesPerSliceSegmentBitKhr   VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 8)
	VkVideoEncodeH265CapabilityMultipleSliceSegmentsPerTileBitKhr   VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 9)
)

type VkVideoEncodeH265StdFlagBitsKHR int64

var (
	VkVideoEncodeH265StdSeparateColorPlaneFlagSetBitKhr              VkVideoEncodeH265StdFlagBitsKHR = (1 << 0)
	VkVideoEncodeH265StdSampleAdaptiveOffsetEnabledFlagSetBitKhr     VkVideoEncodeH265StdFlagBitsKHR = (1 << 1)
	VkVideoEncodeH265StdScalingListDataPresentFlagSetBitKhr          VkVideoEncodeH265StdFlagBitsKHR = (1 << 2)
	VkVideoEncodeH265StdPcmEnabledFlagSetBitKhr                      VkVideoEncodeH265StdFlagBitsKHR = (1 << 3)
	VkVideoEncodeH265StdSpsTemporalMvpEnabledFlagSetBitKhr           VkVideoEncodeH265StdFlagBitsKHR = (1 << 4)
	VkVideoEncodeH265StdInitQpMinus26BitKhr                          VkVideoEncodeH265StdFlagBitsKHR = (1 << 5)
	VkVideoEncodeH265StdWeightedPredFlagSetBitKhr                    VkVideoEncodeH265StdFlagBitsKHR = (1 << 6)
	VkVideoEncodeH265StdWeightedBipredFlagSetBitKhr                  VkVideoEncodeH265StdFlagBitsKHR = (1 << 7)
	VkVideoEncodeH265StdLog2ParallelMergeLevelMinus2BitKhr           VkVideoEncodeH265StdFlagBitsKHR = (1 << 8)
	VkVideoEncodeH265StdSignDataHidingEnabledFlagSetBitKhr           VkVideoEncodeH265StdFlagBitsKHR = (1 << 9)
	VkVideoEncodeH265StdTransformSkipEnabledFlagSetBitKhr            VkVideoEncodeH265StdFlagBitsKHR = (1 << 10)
	VkVideoEncodeH265StdTransformSkipEnabledFlagUnsetBitKhr          VkVideoEncodeH265StdFlagBitsKHR = (1 << 11)
	VkVideoEncodeH265StdPpsSliceChromaQpOffsetsPresentFlagSetBitKhr  VkVideoEncodeH265StdFlagBitsKHR = (1 << 12)
	VkVideoEncodeH265StdTransquantBypassEnabledFlagSetBitKhr         VkVideoEncodeH265StdFlagBitsKHR = (1 << 13)
	VkVideoEncodeH265StdConstrainedIntraPredFlagSetBitKhr            VkVideoEncodeH265StdFlagBitsKHR = (1 << 14)
	VkVideoEncodeH265StdEntropyCodingSyncEnabledFlagSetBitKhr        VkVideoEncodeH265StdFlagBitsKHR = (1 << 15)
	VkVideoEncodeH265StdDeblockingFilterOverrideEnabledFlagSetBitKhr VkVideoEncodeH265StdFlagBitsKHR = (1 << 16)
	VkVideoEncodeH265StdDependentSliceSegmentsEnabledFlagSetBitKhr   VkVideoEncodeH265StdFlagBitsKHR = (1 << 17)
	VkVideoEncodeH265StdDependentSliceSegmentFlagSetBitKhr           VkVideoEncodeH265StdFlagBitsKHR = (1 << 18)
	VkVideoEncodeH265StdSliceQpDeltaBitKhr                           VkVideoEncodeH265StdFlagBitsKHR = (1 << 19)
	VkVideoEncodeH265StdDifferentSliceQpDeltaBitKhr                  VkVideoEncodeH265StdFlagBitsKHR = (1 << 20)
)

type VkVideoEncodeH265RateControlFlagBitsKHR int64

var (
	VkVideoEncodeH265RateControlAttemptHrdComplianceBitKhr          VkVideoEncodeH265RateControlFlagBitsKHR = (1 << 0)
	VkVideoEncodeH265RateControlRegularGopBitKhr                    VkVideoEncodeH265RateControlFlagBitsKHR = (1 << 1)
	VkVideoEncodeH265RateControlReferencePatternFlatBitKhr          VkVideoEncodeH265RateControlFlagBitsKHR = (1 << 2)
	VkVideoEncodeH265RateControlReferencePatternDyadicBitKhr        VkVideoEncodeH265RateControlFlagBitsKHR = (1 << 3)
	VkVideoEncodeH265RateControlTemporalSubLayerPatternDyadicBitKhr VkVideoEncodeH265RateControlFlagBitsKHR = (1 << 4)
)

type VkVideoEncodeH265CtbSizeFlagBitsKHR int64

var (
	VkVideoEncodeH265CtbSize16BitKhr VkVideoEncodeH265CtbSizeFlagBitsKHR = (1 << 0)
	VkVideoEncodeH265CtbSize32BitKhr VkVideoEncodeH265CtbSizeFlagBitsKHR = (1 << 1)
	VkVideoEncodeH265CtbSize64BitKhr VkVideoEncodeH265CtbSizeFlagBitsKHR = (1 << 2)
)

type VkVideoEncodeH265TransformBlockSizeFlagBitsKHR int64

var (
	VkVideoEncodeH265TransformBlockSize4BitKhr  VkVideoEncodeH265TransformBlockSizeFlagBitsKHR = (1 << 0)
	VkVideoEncodeH265TransformBlockSize8BitKhr  VkVideoEncodeH265TransformBlockSizeFlagBitsKHR = (1 << 1)
	VkVideoEncodeH265TransformBlockSize16BitKhr VkVideoEncodeH265TransformBlockSizeFlagBitsKHR = (1 << 2)
	VkVideoEncodeH265TransformBlockSize32BitKhr VkVideoEncodeH265TransformBlockSizeFlagBitsKHR = (1 << 3)
)

type VkVideoEncodeAV1CapabilityFlagBitsKHR int64

var (
	VkVideoEncodeAv1CapabilityPerRateControlGroupMinMaxQIndexBitKhr VkVideoEncodeAV1CapabilityFlagBitsKHR = (1 << 0)
	VkVideoEncodeAv1CapabilityGenerateObuExtensionHeaderBitKhr      VkVideoEncodeAV1CapabilityFlagBitsKHR = (1 << 1)
	VkVideoEncodeAv1CapabilityPrimaryReferenceCdfOnlyBitKhr         VkVideoEncodeAV1CapabilityFlagBitsKHR = (1 << 2)
	VkVideoEncodeAv1CapabilityFrameSizeOverrideBitKhr               VkVideoEncodeAV1CapabilityFlagBitsKHR = (1 << 3)
	VkVideoEncodeAv1CapabilityMotionVectorScalingBitKhr             VkVideoEncodeAV1CapabilityFlagBitsKHR = (1 << 4)
)

type VkVideoEncodeAV1StdFlagBitsKHR int64

var (
	VkVideoEncodeAv1StdUniformTileSpacingFlagSetBitKhr VkVideoEncodeAV1StdFlagBitsKHR = (1 << 0)
	VkVideoEncodeAv1StdSkipModePresentUnsetBitKhr      VkVideoEncodeAV1StdFlagBitsKHR = (1 << 1)
	VkVideoEncodeAv1StdPrimaryRefFrameBitKhr           VkVideoEncodeAV1StdFlagBitsKHR = (1 << 2)
	VkVideoEncodeAv1StdDeltaQBitKhr                    VkVideoEncodeAV1StdFlagBitsKHR = (1 << 3)
)

type VkVideoEncodeAV1RateControlFlagBitsKHR int64

var (
	VkVideoEncodeAv1RateControlRegularGopBitKhr                 VkVideoEncodeAV1RateControlFlagBitsKHR = (1 << 0)
	VkVideoEncodeAv1RateControlTemporalLayerPatternDyadicBitKhr VkVideoEncodeAV1RateControlFlagBitsKHR = (1 << 1)
	VkVideoEncodeAv1RateControlReferencePatternFlatBitKhr       VkVideoEncodeAV1RateControlFlagBitsKHR = (1 << 2)
	VkVideoEncodeAv1RateControlReferencePatternDyadicBitKhr     VkVideoEncodeAV1RateControlFlagBitsKHR = (1 << 3)
)

type VkVideoEncodeAV1SuperblockSizeFlagBitsKHR int64

var (
	VkVideoEncodeAv1SuperblockSize64BitKhr  VkVideoEncodeAV1SuperblockSizeFlagBitsKHR = (1 << 0)
	VkVideoEncodeAv1SuperblockSize128BitKhr VkVideoEncodeAV1SuperblockSizeFlagBitsKHR = (1 << 1)
)

type VkVideoEncodeAV1PredictionModeKHR int64

var (
	VkVideoEncodeAv1PredictionModeIntraOnlyKhr              VkVideoEncodeAV1PredictionModeKHR = 0
	VkVideoEncodeAv1PredictionModeSingleReferenceKhr        VkVideoEncodeAV1PredictionModeKHR = 1
	VkVideoEncodeAv1PredictionModeUnidirectionalCompoundKhr VkVideoEncodeAV1PredictionModeKHR = 2
	VkVideoEncodeAv1PredictionModeBidirectionalCompoundKhr  VkVideoEncodeAV1PredictionModeKHR = 3
)

type VkVideoEncodeAV1RateControlGroupKHR int64

var (
	VkVideoEncodeAv1RateControlGroupIntraKhr        VkVideoEncodeAV1RateControlGroupKHR = 0
	VkVideoEncodeAv1RateControlGroupPredictiveKhr   VkVideoEncodeAV1RateControlGroupKHR = 1
	VkVideoEncodeAv1RateControlGroupBipredictiveKhr VkVideoEncodeAV1RateControlGroupKHR = 2
)

type VkExportMetalObjectTypeFlagBitsEXT int64

var (
	VkExportMetalObjectTypeMetalDeviceBitExt       VkExportMetalObjectTypeFlagBitsEXT = (1 << 0)
	VkExportMetalObjectTypeMetalCommandQueueBitExt VkExportMetalObjectTypeFlagBitsEXT = (1 << 1)
	VkExportMetalObjectTypeMetalBufferBitExt       VkExportMetalObjectTypeFlagBitsEXT = (1 << 2)
	VkExportMetalObjectTypeMetalTextureBitExt      VkExportMetalObjectTypeFlagBitsEXT = (1 << 3)
	VkExportMetalObjectTypeMetalIosurfaceBitExt    VkExportMetalObjectTypeFlagBitsEXT = (1 << 4)
	VkExportMetalObjectTypeMetalSharedEventBitExt  VkExportMetalObjectTypeFlagBitsEXT = (1 << 5)
)

type VkInstanceCreateFlagBits int64

var ()

type VkImageCompressionFlagBitsEXT int64

var (
	VkImageCompressionDefaultExt           VkImageCompressionFlagBitsEXT = (1 << 0)
	VkImageCompressionFixedRateDefaultExt  VkImageCompressionFlagBitsEXT = (1 << 0)
	VkImageCompressionFixedRateExplicitExt VkImageCompressionFlagBitsEXT = (1 << 1)
	VkImageCompressionDisabledExt          VkImageCompressionFlagBitsEXT = (1 << 2)
)

type VkImageCompressionFixedRateFlagBitsEXT int64

var (
	VkImageCompressionFixedRateNoneExt     VkImageCompressionFixedRateFlagBitsEXT = (1 << 0)
	VkImageCompressionFixedRate1bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 0)
	VkImageCompressionFixedRate2bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 1)
	VkImageCompressionFixedRate3bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 2)
	VkImageCompressionFixedRate4bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 3)
	VkImageCompressionFixedRate5bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 4)
	VkImageCompressionFixedRate6bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 5)
	VkImageCompressionFixedRate7bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 6)
	VkImageCompressionFixedRate8bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 7)
	VkImageCompressionFixedRate9bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 8)
	VkImageCompressionFixedRate10bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 9)
	VkImageCompressionFixedRate11bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 10)
	VkImageCompressionFixedRate12bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 11)
	VkImageCompressionFixedRate13bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 12)
	VkImageCompressionFixedRate14bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 13)
	VkImageCompressionFixedRate15bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 14)
	VkImageCompressionFixedRate16bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 15)
	VkImageCompressionFixedRate17bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 16)
	VkImageCompressionFixedRate18bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 17)
	VkImageCompressionFixedRate19bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 18)
	VkImageCompressionFixedRate20bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 19)
	VkImageCompressionFixedRate21bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 20)
	VkImageCompressionFixedRate22bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 21)
	VkImageCompressionFixedRate23bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 22)
	VkImageCompressionFixedRate24bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 23)
)

type VkPipelineRobustnessBufferBehavior int64

var (
	VkPipelineRobustnessBufferBehaviorDeviceDefault       VkPipelineRobustnessBufferBehavior = 0
	VkPipelineRobustnessBufferBehaviorDisabled            VkPipelineRobustnessBufferBehavior = 1
	VkPipelineRobustnessBufferBehaviorRobustBufferAccess  VkPipelineRobustnessBufferBehavior = 2
	VkPipelineRobustnessBufferBehaviorRobustBufferAccess2 VkPipelineRobustnessBufferBehavior = 3
)

type VkPipelineRobustnessImageBehavior int64

var (
	VkPipelineRobustnessImageBehaviorDeviceDefault      VkPipelineRobustnessImageBehavior = 0
	VkPipelineRobustnessImageBehaviorDisabled           VkPipelineRobustnessImageBehavior = 1
	VkPipelineRobustnessImageBehaviorRobustImageAccess  VkPipelineRobustnessImageBehavior = 2
	VkPipelineRobustnessImageBehaviorRobustImageAccess2 VkPipelineRobustnessImageBehavior = 3
)

type VkOpticalFlowGridSizeFlagBitsNV int64

var (
	VkOpticalFlowGridSizeUnknownNv VkOpticalFlowGridSizeFlagBitsNV = (1 << 0)
	VkOpticalFlowGridSize1x1BitNv  VkOpticalFlowGridSizeFlagBitsNV = (1 << 0)
	VkOpticalFlowGridSize2x2BitNv  VkOpticalFlowGridSizeFlagBitsNV = (1 << 1)
	VkOpticalFlowGridSize4x4BitNv  VkOpticalFlowGridSizeFlagBitsNV = (1 << 2)
	VkOpticalFlowGridSize8x8BitNv  VkOpticalFlowGridSizeFlagBitsNV = (1 << 3)
)

type VkOpticalFlowUsageFlagBitsNV int64

var (
	VkOpticalFlowUsageUnknownNv       VkOpticalFlowUsageFlagBitsNV = (1 << 0)
	VkOpticalFlowUsageInputBitNv      VkOpticalFlowUsageFlagBitsNV = (1 << 0)
	VkOpticalFlowUsageOutputBitNv     VkOpticalFlowUsageFlagBitsNV = (1 << 1)
	VkOpticalFlowUsageHintBitNv       VkOpticalFlowUsageFlagBitsNV = (1 << 2)
	VkOpticalFlowUsageCostBitNv       VkOpticalFlowUsageFlagBitsNV = (1 << 3)
	VkOpticalFlowUsageGlobalFlowBitNv VkOpticalFlowUsageFlagBitsNV = (1 << 4)
)

type VkOpticalFlowPerformanceLevelNV int64

var (
	VkOpticalFlowPerformanceLevelUnknownNv VkOpticalFlowPerformanceLevelNV = 0
	VkOpticalFlowPerformanceLevelSlowNv    VkOpticalFlowPerformanceLevelNV = 1
	VkOpticalFlowPerformanceLevelMediumNv  VkOpticalFlowPerformanceLevelNV = 2
	VkOpticalFlowPerformanceLevelFastNv    VkOpticalFlowPerformanceLevelNV = 3
)

type VkOpticalFlowSessionBindingPointNV int64

var (
	VkOpticalFlowSessionBindingPointUnknownNv            VkOpticalFlowSessionBindingPointNV = 0
	VkOpticalFlowSessionBindingPointInputNv              VkOpticalFlowSessionBindingPointNV = 1
	VkOpticalFlowSessionBindingPointReferenceNv          VkOpticalFlowSessionBindingPointNV = 2
	VkOpticalFlowSessionBindingPointHintNv               VkOpticalFlowSessionBindingPointNV = 3
	VkOpticalFlowSessionBindingPointFlowVectorNv         VkOpticalFlowSessionBindingPointNV = 4
	VkOpticalFlowSessionBindingPointBackwardFlowVectorNv VkOpticalFlowSessionBindingPointNV = 5
	VkOpticalFlowSessionBindingPointCostNv               VkOpticalFlowSessionBindingPointNV = 6
	VkOpticalFlowSessionBindingPointBackwardCostNv       VkOpticalFlowSessionBindingPointNV = 7
	VkOpticalFlowSessionBindingPointGlobalFlowNv         VkOpticalFlowSessionBindingPointNV = 8
)

type VkOpticalFlowSessionCreateFlagBitsNV int64

var (
	VkOpticalFlowSessionCreateEnableHintBitNv       VkOpticalFlowSessionCreateFlagBitsNV = (1 << 0)
	VkOpticalFlowSessionCreateEnableCostBitNv       VkOpticalFlowSessionCreateFlagBitsNV = (1 << 1)
	VkOpticalFlowSessionCreateEnableGlobalFlowBitNv VkOpticalFlowSessionCreateFlagBitsNV = (1 << 2)
	VkOpticalFlowSessionCreateAllowRegionsBitNv     VkOpticalFlowSessionCreateFlagBitsNV = (1 << 3)
	VkOpticalFlowSessionCreateBothDirectionsBitNv   VkOpticalFlowSessionCreateFlagBitsNV = (1 << 4)
)

type VkOpticalFlowExecuteFlagBitsNV int64

var (
	VkOpticalFlowExecuteDisableTemporalHintsBitNv VkOpticalFlowExecuteFlagBitsNV = (1 << 0)
)

type VkMicromapTypeEXT int64

var (
	VkMicromapTypeOpacityMicromapExt VkMicromapTypeEXT = 0
)

type VkBuildMicromapFlagBitsEXT int64

var (
	VkBuildMicromapPreferFastTraceBitExt VkBuildMicromapFlagBitsEXT = (1 << 0)
	VkBuildMicromapPreferFastBuildBitExt VkBuildMicromapFlagBitsEXT = (1 << 1)
	VkBuildMicromapAllowCompactionBitExt VkBuildMicromapFlagBitsEXT = (1 << 2)
)

type VkMicromapCreateFlagBitsEXT int64

var (
	VkMicromapCreateDeviceAddressCaptureReplayBitExt VkMicromapCreateFlagBitsEXT = (1 << 0)
)

type VkCopyMicromapModeEXT int64

var (
	VkCopyMicromapModeCloneExt       VkCopyMicromapModeEXT = 0
	VkCopyMicromapModeSerializeExt   VkCopyMicromapModeEXT = 1
	VkCopyMicromapModeDeserializeExt VkCopyMicromapModeEXT = 2
	VkCopyMicromapModeCompactExt     VkCopyMicromapModeEXT = 3
)

type VkBuildMicromapModeEXT int64

var (
	VkBuildMicromapModeBuildExt VkBuildMicromapModeEXT = 0
)

type VkOpacityMicromapFormatEXT int64

var (
	VkOpacityMicromapFormat2StateExt VkOpacityMicromapFormatEXT = 1
	VkOpacityMicromapFormat4StateExt VkOpacityMicromapFormatEXT = 2
)

type VkOpacityMicromapSpecialIndexEXT int64

var (
	VkOpacityMicromapSpecialIndexFullyTransparentExt        VkOpacityMicromapSpecialIndexEXT = -1
	VkOpacityMicromapSpecialIndexFullyOpaqueExt             VkOpacityMicromapSpecialIndexEXT = -2
	VkOpacityMicromapSpecialIndexFullyUnknownTransparentExt VkOpacityMicromapSpecialIndexEXT = -3
	VkOpacityMicromapSpecialIndexFullyUnknownOpaqueExt      VkOpacityMicromapSpecialIndexEXT = -4
)

type VkDepthBiasRepresentationEXT int64

var (
	VkDepthBiasRepresentationLeastRepresentableValueFormatExt     VkDepthBiasRepresentationEXT = 0
	VkDepthBiasRepresentationLeastRepresentableValueForceUnormExt VkDepthBiasRepresentationEXT = 1
	VkDepthBiasRepresentationFloatExt                             VkDepthBiasRepresentationEXT = 2
)

type VkDeviceFaultAddressTypeEXT int64

var (
	VkDeviceFaultAddressTypeNoneExt                      VkDeviceFaultAddressTypeEXT = 0
	VkDeviceFaultAddressTypeReadInvalidExt               VkDeviceFaultAddressTypeEXT = 1
	VkDeviceFaultAddressTypeWriteInvalidExt              VkDeviceFaultAddressTypeEXT = 2
	VkDeviceFaultAddressTypeExecuteInvalidExt            VkDeviceFaultAddressTypeEXT = 3
	VkDeviceFaultAddressTypeInstructionPointerUnknownExt VkDeviceFaultAddressTypeEXT = 4
	VkDeviceFaultAddressTypeInstructionPointerInvalidExt VkDeviceFaultAddressTypeEXT = 5
	VkDeviceFaultAddressTypeInstructionPointerFaultExt   VkDeviceFaultAddressTypeEXT = 6
)

type VkDeviceFaultVendorBinaryHeaderVersionEXT int64

var (
	VkDeviceFaultVendorBinaryHeaderVersionOneExt VkDeviceFaultVendorBinaryHeaderVersionEXT = 1
)

type VkIndirectCommandsLayoutUsageFlagBitsEXT int64

var (
	VkIndirectCommandsLayoutUsageExplicitPreprocessBitExt VkIndirectCommandsLayoutUsageFlagBitsEXT = (1 << 0)
	VkIndirectCommandsLayoutUsageUnorderedSequencesBitExt VkIndirectCommandsLayoutUsageFlagBitsEXT = (1 << 1)
)

type VkIndirectExecutionSetInfoTypeEXT int64

var (
	VkIndirectExecutionSetInfoTypePipelinesExt     VkIndirectExecutionSetInfoTypeEXT = 0
	VkIndirectExecutionSetInfoTypeShaderObjectsExt VkIndirectExecutionSetInfoTypeEXT = 1
)

type VkIndirectCommandsInputModeFlagBitsEXT int64

var (
	VkIndirectCommandsInputModeVulkanIndexBufferExt VkIndirectCommandsInputModeFlagBitsEXT = (1 << 0)
	VkIndirectCommandsInputModeDxgiIndexBufferExt   VkIndirectCommandsInputModeFlagBitsEXT = (1 << 1)
)

type VkIndirectCommandsTokenTypeEXT int64

var (
	VkIndirectCommandsTokenTypeExecutionSetExt     VkIndirectCommandsTokenTypeEXT = 0
	VkIndirectCommandsTokenTypePushConstantExt     VkIndirectCommandsTokenTypeEXT = 1
	VkIndirectCommandsTokenTypeSequenceIndexExt    VkIndirectCommandsTokenTypeEXT = 2
	VkIndirectCommandsTokenTypeIndexBufferExt      VkIndirectCommandsTokenTypeEXT = 3
	VkIndirectCommandsTokenTypeVertexBufferExt     VkIndirectCommandsTokenTypeEXT = 4
	VkIndirectCommandsTokenTypeDrawIndexedExt      VkIndirectCommandsTokenTypeEXT = 5
	VkIndirectCommandsTokenTypeDrawExt             VkIndirectCommandsTokenTypeEXT = 6
	VkIndirectCommandsTokenTypeDrawIndexedCountExt VkIndirectCommandsTokenTypeEXT = 7
	VkIndirectCommandsTokenTypeDrawCountExt        VkIndirectCommandsTokenTypeEXT = 8
	VkIndirectCommandsTokenTypeDispatchExt         VkIndirectCommandsTokenTypeEXT = 9
)

type VkDisplacementMicromapFormatNV int64

var (
	VkDisplacementMicromapFormat64Triangles64BytesNv    VkDisplacementMicromapFormatNV = 1
	VkDisplacementMicromapFormat256Triangles128BytesNv  VkDisplacementMicromapFormatNV = 2
	VkDisplacementMicromapFormat1024Triangles128BytesNv VkDisplacementMicromapFormatNV = 3
)

type VkShaderCreateFlagBitsEXT int64

var (
	VkShaderCreateLinkStageBitExt VkShaderCreateFlagBitsEXT = (1 << 0)
)

type VkShaderCodeTypeEXT int64

var (
	VkShaderCodeTypeBinaryExt VkShaderCodeTypeEXT = 0
	VkShaderCodeTypeSpirvExt  VkShaderCodeTypeEXT = 1
)

type VkScopeKHR int64

var (
	VkScopeDeviceKhr      VkScopeKHR = 1
	VkScopeWorkgroupKhr   VkScopeKHR = 2
	VkScopeSubgroupKhr    VkScopeKHR = 3
	VkScopeQueueFamilyKhr VkScopeKHR = 5
)

type VkComponentTypeKHR int64

var (
	VkComponentTypeFloat16Khr VkComponentTypeKHR = 0
	VkComponentTypeFloat32Khr VkComponentTypeKHR = 1
	VkComponentTypeFloat64Khr VkComponentTypeKHR = 2
	VkComponentTypeSint8Khr   VkComponentTypeKHR = 3
	VkComponentTypeSint16Khr  VkComponentTypeKHR = 4
	VkComponentTypeSint32Khr  VkComponentTypeKHR = 5
	VkComponentTypeSint64Khr  VkComponentTypeKHR = 6
	VkComponentTypeUint8Khr   VkComponentTypeKHR = 7
	VkComponentTypeUint16Khr  VkComponentTypeKHR = 8
	VkComponentTypeUint32Khr  VkComponentTypeKHR = 9
	VkComponentTypeUint64Khr  VkComponentTypeKHR = 10
)

type VkCubicFilterWeightsQCOM int64

var (
	VkCubicFilterWeightsCatmullRomQcom          VkCubicFilterWeightsQCOM = 0
	VkCubicFilterWeightsZeroTangentCardinalQcom VkCubicFilterWeightsQCOM = 1
	VkCubicFilterWeightsBSplineQcom             VkCubicFilterWeightsQCOM = 2
	VkCubicFilterWeightsMitchellNetravaliQcom   VkCubicFilterWeightsQCOM = 3
)

type VkBlockMatchWindowCompareModeQCOM int64

var (
	VkBlockMatchWindowCompareModeMinQcom VkBlockMatchWindowCompareModeQCOM = 0
	VkBlockMatchWindowCompareModeMaxQcom VkBlockMatchWindowCompareModeQCOM = 1
)

type VkPhysicalDeviceLayeredApiKHR int64

var (
	VkPhysicalDeviceLayeredApiVulkanKhr   VkPhysicalDeviceLayeredApiKHR = 0
	VkPhysicalDeviceLayeredApiD3d12Khr    VkPhysicalDeviceLayeredApiKHR = 1
	VkPhysicalDeviceLayeredApiMetalKhr    VkPhysicalDeviceLayeredApiKHR = 2
	VkPhysicalDeviceLayeredApiOpenglKhr   VkPhysicalDeviceLayeredApiKHR = 3
	VkPhysicalDeviceLayeredApiOpenglesKhr VkPhysicalDeviceLayeredApiKHR = 4
)

type VkLayeredDriverUnderlyingApiMSFT int64

var (
	VkLayeredDriverUnderlyingApiNoneMsft  VkLayeredDriverUnderlyingApiMSFT = 0
	VkLayeredDriverUnderlyingApiD3d12Msft VkLayeredDriverUnderlyingApiMSFT = 1
)

type VkLatencyMarkerNV int64

var (
	VkLatencyMarkerSimulationStartNv            VkLatencyMarkerNV = 0
	VkLatencyMarkerSimulationEndNv              VkLatencyMarkerNV = 1
	VkLatencyMarkerRendersubmitStartNv          VkLatencyMarkerNV = 2
	VkLatencyMarkerRendersubmitEndNv            VkLatencyMarkerNV = 3
	VkLatencyMarkerPresentStartNv               VkLatencyMarkerNV = 4
	VkLatencyMarkerPresentEndNv                 VkLatencyMarkerNV = 5
	VkLatencyMarkerInputSampleNv                VkLatencyMarkerNV = 6
	VkLatencyMarkerTriggerFlashNv               VkLatencyMarkerNV = 7
	VkLatencyMarkerOutOfBandRendersubmitStartNv VkLatencyMarkerNV = 8
	VkLatencyMarkerOutOfBandRendersubmitEndNv   VkLatencyMarkerNV = 9
	VkLatencyMarkerOutOfBandPresentStartNv      VkLatencyMarkerNV = 10
	VkLatencyMarkerOutOfBandPresentEndNv        VkLatencyMarkerNV = 11
)

type VkOutOfBandQueueTypeNV int64

var (
	VkOutOfBandQueueTypeRenderNv  VkOutOfBandQueueTypeNV = 0
	VkOutOfBandQueueTypePresentNv VkOutOfBandQueueTypeNV = 1
)

type VkMemoryUnmapFlagBits int64

var ()

type VkCompressedTriangleFormatAMDX int64

var (
	VkCompressedTriangleFormatDgf1Amdx VkCompressedTriangleFormatAMDX = 0
)

type VkWaylandSurfaceCreateFlagBitsKHR int64

var ()

type VkDepthClampModeEXT int64

var (
	VkDepthClampModeViewportRangeExt    VkDepthClampModeEXT = 0
	VkDepthClampModeUserDefinedRangeExt VkDepthClampModeEXT = 1
)

type VkAccessFlagBits3KHR int64

var (
	VkAccess3NoneKhr VkAccessFlagBits3KHR = (1 << 0)
)

type VkTileShadingRenderPassFlagBitsQCOM int64

var (
	VkTileShadingRenderPassEnableBitQcom           VkTileShadingRenderPassFlagBitsQCOM = (1 << 0)
	VkTileShadingRenderPassPerTileExecutionBitQcom VkTileShadingRenderPassFlagBitsQCOM = (1 << 1)
)

type VkCooperativeVectorMatrixLayoutNV int64

var (
	VkCooperativeVectorMatrixLayoutRowMajorNv           VkCooperativeVectorMatrixLayoutNV = 0
	VkCooperativeVectorMatrixLayoutColumnMajorNv        VkCooperativeVectorMatrixLayoutNV = 1
	VkCooperativeVectorMatrixLayoutInferencingOptimalNv VkCooperativeVectorMatrixLayoutNV = 2
	VkCooperativeVectorMatrixLayoutTrainingOptimalNv    VkCooperativeVectorMatrixLayoutNV = 3
)

type VkAddressCopyFlagBitsKHR int64

var (
	VkAddressCopyDeviceLocalBitKhr VkAddressCopyFlagBitsKHR = (1 << 0)
	VkAddressCopySparseBitKhr      VkAddressCopyFlagBitsKHR = (1 << 1)
	VkAddressCopyProtectedBitKhr   VkAddressCopyFlagBitsKHR = (1 << 2)
)

type VkTensorCreateFlagBitsARM int64

var (
	VkTensorCreateMutableFormatBitArm VkTensorCreateFlagBitsARM = (1 << 0)
	VkTensorCreateProtectedBitArm     VkTensorCreateFlagBitsARM = (1 << 1)
)

type VkTensorUsageFlagBitsARM int64

var (
	VkTensorUsageShaderBitArm        VkTensorUsageFlagBitsARM = (1 << 1)
	VkTensorUsageTransferSrcBitArm   VkTensorUsageFlagBitsARM = (1 << 2)
	VkTensorUsageTransferDstBitArm   VkTensorUsageFlagBitsARM = (1 << 3)
	VkTensorUsageImageAliasingBitArm VkTensorUsageFlagBitsARM = (1 << 4)
)

type VkTensorTilingARM int64

var (
	VkTensorTilingOptimalArm VkTensorTilingARM = 0
	VkTensorTilingLinearArm  VkTensorTilingARM = 1
)

type VkTensorViewCreateFlagBitsARM int64

var ()

type VkDefaultVertexAttributeValueKHR int64

var (
	VkDefaultVertexAttributeValueZeroZeroZeroZeroKhr VkDefaultVertexAttributeValueKHR = 0
	VkDefaultVertexAttributeValueZeroZeroZeroOneKhr  VkDefaultVertexAttributeValueKHR = 1
)

type VkDataGraphPipelineSessionCreateFlagBitsARM int64

var (
	VkDataGraphPipelineSessionCreateProtectedBitArm VkDataGraphPipelineSessionCreateFlagBitsARM = (1 << 0)
)

type VkDataGraphPipelineSessionBindPointARM int64

var (
	VkDataGraphPipelineSessionBindPointTransientArm VkDataGraphPipelineSessionBindPointARM = 0
)

type VkDataGraphPipelineSessionBindPointTypeARM int64

var (
	VkDataGraphPipelineSessionBindPointTypeMemoryArm VkDataGraphPipelineSessionBindPointTypeARM = 0
)

type VkDataGraphPipelinePropertyARM int64

var (
	VkDataGraphPipelinePropertyCreationLogArm VkDataGraphPipelinePropertyARM = 0
	VkDataGraphPipelinePropertyIdentifierArm  VkDataGraphPipelinePropertyARM = 1
)

type VkDataGraphPipelineDispatchFlagBitsARM int64

var ()

type VkPhysicalDeviceDataGraphProcessingEngineTypeARM int64

var (
	VkPhysicalDeviceDataGraphProcessingEngineTypeDefaultArm VkPhysicalDeviceDataGraphProcessingEngineTypeARM = 0
)

type VkPhysicalDeviceDataGraphOperationTypeARM int64

var (
	VkPhysicalDeviceDataGraphOperationTypeSpirvExtendedInstructionSetArm VkPhysicalDeviceDataGraphOperationTypeARM = 0
)

type VkDataGraphModelCacheTypeQCOM int64

var (
	VkDataGraphModelCacheTypeGenericBinaryQcom VkDataGraphModelCacheTypeQCOM = 0
)

type VkVideoEncodeRgbModelConversionFlagBitsVALVE int64

var (
	VkVideoEncodeRgbModelConversionRgbIdentityBitValve   VkVideoEncodeRgbModelConversionFlagBitsVALVE = (1 << 0)
	VkVideoEncodeRgbModelConversionYcbcrIdentityBitValve VkVideoEncodeRgbModelConversionFlagBitsVALVE = (1 << 1)
	VkVideoEncodeRgbModelConversionYcbcr709BitValve      VkVideoEncodeRgbModelConversionFlagBitsVALVE = (1 << 2)
	VkVideoEncodeRgbModelConversionYcbcr601BitValve      VkVideoEncodeRgbModelConversionFlagBitsVALVE = (1 << 3)
	VkVideoEncodeRgbModelConversionYcbcr2020BitValve     VkVideoEncodeRgbModelConversionFlagBitsVALVE = (1 << 4)
)

type VkVideoEncodeRgbRangeCompressionFlagBitsVALVE int64

var (
	VkVideoEncodeRgbRangeCompressionFullRangeBitValve   VkVideoEncodeRgbRangeCompressionFlagBitsVALVE = (1 << 0)
	VkVideoEncodeRgbRangeCompressionNarrowRangeBitValve VkVideoEncodeRgbRangeCompressionFlagBitsVALVE = (1 << 1)
)

type VkVideoEncodeRgbChromaOffsetFlagBitsVALVE int64

var (
	VkVideoEncodeRgbChromaOffsetCositedEvenBitValve VkVideoEncodeRgbChromaOffsetFlagBitsVALVE = (1 << 0)
	VkVideoEncodeRgbChromaOffsetMidpointBitValve    VkVideoEncodeRgbChromaOffsetFlagBitsVALVE = (1 << 1)
)

type VkSwapchainImageUsageFlagBitsOHOS int64

var (
	VkSwapchainImageUsageSharedBitOhos VkSwapchainImageUsageFlagBitsOHOS = (1 << 0)
)

type VkDescriptorMappingSourceEXT int64

var (
	VkDescriptorMappingSourceHeapWithConstantOffsetExt     VkDescriptorMappingSourceEXT = 0
	VkDescriptorMappingSourceHeapWithPushIndexExt          VkDescriptorMappingSourceEXT = 1
	VkDescriptorMappingSourceHeapWithIndirectIndexExt      VkDescriptorMappingSourceEXT = 2
	VkDescriptorMappingSourceHeapWithIndirectIndexArrayExt VkDescriptorMappingSourceEXT = 3
	VkDescriptorMappingSourceResourceHeapDataExt           VkDescriptorMappingSourceEXT = 4
	VkDescriptorMappingSourcePushDataExt                   VkDescriptorMappingSourceEXT = 5
	VkDescriptorMappingSourcePushAddressExt                VkDescriptorMappingSourceEXT = 6
	VkDescriptorMappingSourceIndirectAddressExt            VkDescriptorMappingSourceEXT = 7
)

type VkSpirvResourceTypeFlagBitsEXT int64

var (
	VkSpirvResourceTypeAllExt                       VkSpirvResourceTypeFlagBitsEXT = (1 << 0)
	VkSpirvResourceTypeSamplerBitExt                VkSpirvResourceTypeFlagBitsEXT = (1 << 0)
	VkSpirvResourceTypeSampledImageBitExt           VkSpirvResourceTypeFlagBitsEXT = (1 << 1)
	VkSpirvResourceTypeReadOnlyImageBitExt          VkSpirvResourceTypeFlagBitsEXT = (1 << 2)
	VkSpirvResourceTypeReadWriteImageBitExt         VkSpirvResourceTypeFlagBitsEXT = (1 << 3)
	VkSpirvResourceTypeCombinedSampledImageBitExt   VkSpirvResourceTypeFlagBitsEXT = (1 << 4)
	VkSpirvResourceTypeUniformBufferBitExt          VkSpirvResourceTypeFlagBitsEXT = (1 << 5)
	VkSpirvResourceTypeReadOnlyStorageBufferBitExt  VkSpirvResourceTypeFlagBitsEXT = (1 << 6)
	VkSpirvResourceTypeReadWriteStorageBufferBitExt VkSpirvResourceTypeFlagBitsEXT = (1 << 7)
)
