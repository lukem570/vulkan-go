#ifndef VOLK_WRAPPERS_H
#define VOLK_WRAPPERS_H

#if defined(__APPLE__)
#define VK_USE_PLATFORM_MACOS_MVK
#endif

#if defined(__linux__) || defined(__unix__)
#include <wayland-client.h>
#define VK_USE_PLATFORM_WAYLAND_KHR
#include <xcb/xcb.h>
#define VK_USE_PLATFORM_XCB_KHR
#include <X11/Xlib.h>
#define VK_USE_PLATFORM_XLIB_KHR
#endif

#if defined(_WIN32)
#include <windows.h>
#define VK_USE_PLATFORM_WIN32_KHR
#endif

#include "volk.h"
#include <stdlib.h>

#include <vulkan/vulkan.h>

static inline void vk_set_VkAccelerationStructureInstanceKHR_instanceCustomIndex(VkAccelerationStructureInstanceKHR* s, uint32_t v) { s->instanceCustomIndex = v; }
static inline uint32_t vk_get_VkAccelerationStructureInstanceKHR_instanceCustomIndex(const VkAccelerationStructureInstanceKHR* s) { return (uint32_t)s->instanceCustomIndex; }
static inline void vk_set_VkAccelerationStructureInstanceKHR_mask(VkAccelerationStructureInstanceKHR* s, uint32_t v) { s->mask = v; }
static inline uint32_t vk_get_VkAccelerationStructureInstanceKHR_mask(const VkAccelerationStructureInstanceKHR* s) { return (uint32_t)s->mask; }
static inline void vk_set_VkAccelerationStructureInstanceKHR_instanceShaderBindingTableRecordOffset(VkAccelerationStructureInstanceKHR* s, uint32_t v) { s->instanceShaderBindingTableRecordOffset = v; }
static inline uint32_t vk_get_VkAccelerationStructureInstanceKHR_instanceShaderBindingTableRecordOffset(const VkAccelerationStructureInstanceKHR* s) { return (uint32_t)s->instanceShaderBindingTableRecordOffset; }
static inline void vk_set_VkAccelerationStructureInstanceKHR_flags(VkAccelerationStructureInstanceKHR* s, VkGeometryInstanceFlagsKHR v) { s->flags = v; }
static inline VkGeometryInstanceFlagsKHR vk_get_VkAccelerationStructureInstanceKHR_flags(const VkAccelerationStructureInstanceKHR* s) { return (VkGeometryInstanceFlagsKHR)s->flags; }

VkResult vulkan_platform_initialize(void);

VkResult fn_vkAcquireNextImage2KHR(VkDevice device, const VkAcquireNextImageInfoKHR* pAcquireInfo, uint32_t* pImageIndex);
VkResult fn_vkAcquireNextImageKHR(VkDevice device, VkSwapchainKHR swapchain, uint64_t timeout, VkSemaphore semaphore, VkFence fence, uint32_t* pImageIndex);
VkResult fn_vkAllocateCommandBuffers(VkDevice device, const VkCommandBufferAllocateInfo* pAllocateInfo, VkCommandBuffer* pCommandBuffers);
VkResult fn_vkAllocateDescriptorSets(VkDevice device, const VkDescriptorSetAllocateInfo* pAllocateInfo, VkDescriptorSet* pDescriptorSets);
VkResult fn_vkAllocateMemory(VkDevice device, const VkMemoryAllocateInfo* pAllocateInfo, const VkAllocationCallbacks* pAllocator, VkDeviceMemory* pMemory);
VkResult fn_vkBeginCommandBuffer(VkCommandBuffer commandBuffer, const VkCommandBufferBeginInfo* pBeginInfo);
VkResult fn_vkBindBufferMemory(VkDevice device, VkBuffer buffer, VkDeviceMemory memory, VkDeviceSize memoryOffset);
VkResult fn_vkBindBufferMemory2(VkDevice device, uint32_t bindInfoCount, const VkBindBufferMemoryInfo* pBindInfos);
VkResult fn_vkBindImageMemory(VkDevice device, VkImage image, VkDeviceMemory memory, VkDeviceSize memoryOffset);
VkResult fn_vkBindImageMemory2(VkDevice device, uint32_t bindInfoCount, const VkBindImageMemoryInfo* pBindInfos);
VkResult fn_vkBuildAccelerationStructuresKHR(VkDevice device, VkDeferredOperationKHR deferredOperation, uint32_t infoCount, const VkAccelerationStructureBuildGeometryInfoKHR* pInfos, const VkAccelerationStructureBuildRangeInfoKHR* const* ppBuildRangeInfos);
void fn_vkCmdBeginDebugUtilsLabelEXT(VkCommandBuffer commandBuffer, const VkDebugUtilsLabelEXT* pLabelInfo);
void fn_vkCmdBeginQuery(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query, VkQueryControlFlags flags);
void fn_vkCmdBeginRenderPass(VkCommandBuffer commandBuffer, const VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassContents contents);
void fn_vkCmdBeginRenderPass2(VkCommandBuffer commandBuffer, const VkRenderPassBeginInfo* pRenderPassBegin, const VkSubpassBeginInfo* pSubpassBeginInfo);
void fn_vkCmdBeginRendering(VkCommandBuffer commandBuffer, const VkRenderingInfo* pRenderingInfo);
void fn_vkCmdBindDescriptorSets(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t firstSet, uint32_t descriptorSetCount, const VkDescriptorSet* pDescriptorSets, uint32_t dynamicOffsetCount, const uint32_t* pDynamicOffsets);
void fn_vkCmdBindDescriptorSets2(VkCommandBuffer commandBuffer, const VkBindDescriptorSetsInfo* pBindDescriptorSetsInfo);
void fn_vkCmdBindIndexBuffer(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkIndexType indexType);
void fn_vkCmdBindIndexBuffer2(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkDeviceSize size, VkIndexType indexType);
void fn_vkCmdBindPipeline(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipeline pipeline);
void fn_vkCmdBindShadersEXT(VkCommandBuffer commandBuffer, uint32_t stageCount, const VkShaderStageFlagBits* pStages, const VkShaderEXT* pShaders);
void fn_vkCmdBindVertexBuffers(VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, const VkBuffer* pBuffers, const VkDeviceSize* pOffsets);
void fn_vkCmdBindVertexBuffers2(VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, const VkBuffer* pBuffers, const VkDeviceSize* pOffsets, const VkDeviceSize* pSizes, const VkDeviceSize* pStrides);
void fn_vkCmdBlitImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkImageBlit* pRegions, VkFilter filter);
void fn_vkCmdBlitImage2(VkCommandBuffer commandBuffer, const VkBlitImageInfo2* pBlitImageInfo);
void fn_vkCmdBuildAccelerationStructuresIndirectKHR(VkCommandBuffer commandBuffer, uint32_t infoCount, const VkAccelerationStructureBuildGeometryInfoKHR* pInfos, const VkDeviceAddress* pIndirectDeviceAddresses, const uint32_t* pIndirectStrides, const void* ppMaxPrimitiveCounts);
void fn_vkCmdBuildAccelerationStructuresKHR(VkCommandBuffer commandBuffer, uint32_t infoCount, const VkAccelerationStructureBuildGeometryInfoKHR* pInfos, const VkAccelerationStructureBuildRangeInfoKHR* const* ppBuildRangeInfos);
void fn_vkCmdClearAttachments(VkCommandBuffer commandBuffer, uint32_t attachmentCount, const VkClearAttachment* pAttachments, uint32_t rectCount, const VkClearRect* pRects);
void fn_vkCmdClearColorImage(VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, const VkClearColorValue* pColor, uint32_t rangeCount, const VkImageSubresourceRange* pRanges);
void fn_vkCmdClearDepthStencilImage(VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, const VkClearDepthStencilValue* pDepthStencil, uint32_t rangeCount, const VkImageSubresourceRange* pRanges);
void fn_vkCmdCopyAccelerationStructureKHR(VkCommandBuffer commandBuffer, const VkCopyAccelerationStructureInfoKHR* pInfo);
void fn_vkCmdCopyAccelerationStructureToMemoryKHR(VkCommandBuffer commandBuffer, const VkCopyAccelerationStructureToMemoryInfoKHR* pInfo);
void fn_vkCmdCopyBuffer(VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkBuffer dstBuffer, uint32_t regionCount, const VkBufferCopy* pRegions);
void fn_vkCmdCopyBuffer2(VkCommandBuffer commandBuffer, const VkCopyBufferInfo2* pCopyBufferInfo);
void fn_vkCmdCopyBufferToImage(VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkBufferImageCopy* pRegions);
void fn_vkCmdCopyBufferToImage2(VkCommandBuffer commandBuffer, const VkCopyBufferToImageInfo2* pCopyBufferToImageInfo);
void fn_vkCmdCopyImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkImageCopy* pRegions);
void fn_vkCmdCopyImage2(VkCommandBuffer commandBuffer, const VkCopyImageInfo2* pCopyImageInfo);
void fn_vkCmdCopyImageToBuffer(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkBuffer dstBuffer, uint32_t regionCount, const VkBufferImageCopy* pRegions);
void fn_vkCmdCopyImageToBuffer2(VkCommandBuffer commandBuffer, const VkCopyImageToBufferInfo2* pCopyImageToBufferInfo);
void fn_vkCmdCopyMemoryToAccelerationStructureKHR(VkCommandBuffer commandBuffer, const VkCopyMemoryToAccelerationStructureInfoKHR* pInfo);
void fn_vkCmdCopyQueryPoolResults(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize stride, VkQueryResultFlags flags);
void fn_vkCmdDispatch(VkCommandBuffer commandBuffer, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ);
void fn_vkCmdDispatchBase(VkCommandBuffer commandBuffer, uint32_t baseGroupX, uint32_t baseGroupY, uint32_t baseGroupZ, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ);
void fn_vkCmdDispatchIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset);
void fn_vkCmdDraw(VkCommandBuffer commandBuffer, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance);
void fn_vkCmdDrawIndexed(VkCommandBuffer commandBuffer, uint32_t indexCount, uint32_t instanceCount, uint32_t firstIndex, int32_t vertexOffset, uint32_t firstInstance);
void fn_vkCmdDrawIndexedIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride);
void fn_vkCmdDrawIndexedIndirectCount(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride);
void fn_vkCmdDrawIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride);
void fn_vkCmdDrawIndirectCount(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride);
void fn_vkCmdDrawMultiEXT(VkCommandBuffer commandBuffer, uint32_t drawCount, const VkMultiDrawInfoEXT* pVertexInfo, uint32_t instanceCount, uint32_t firstInstance, uint32_t stride);
void fn_vkCmdDrawMultiIndexedEXT(VkCommandBuffer commandBuffer, uint32_t drawCount, const VkMultiDrawIndexedInfoEXT* pIndexInfo, uint32_t instanceCount, uint32_t firstInstance, uint32_t stride, const int32_t* pVertexOffset);
void fn_vkCmdEndDebugUtilsLabelEXT(VkCommandBuffer commandBuffer);
void fn_vkCmdEndQuery(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query);
void fn_vkCmdEndRenderPass(VkCommandBuffer commandBuffer);
void fn_vkCmdEndRenderPass2(VkCommandBuffer commandBuffer, const VkSubpassEndInfo* pSubpassEndInfo);
void fn_vkCmdEndRendering(VkCommandBuffer commandBuffer);
void fn_vkCmdExecuteCommands(VkCommandBuffer commandBuffer, uint32_t commandBufferCount, const VkCommandBuffer* pCommandBuffers);
void fn_vkCmdExecuteGeneratedCommandsEXT(VkCommandBuffer commandBuffer, VkBool32 isPreprocessed, const VkGeneratedCommandsInfoEXT* pGeneratedCommandsInfo);
void fn_vkCmdFillBuffer(VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize size, uint32_t data);
void fn_vkCmdInsertDebugUtilsLabelEXT(VkCommandBuffer commandBuffer, const VkDebugUtilsLabelEXT* pLabelInfo);
void fn_vkCmdNextSubpass(VkCommandBuffer commandBuffer, VkSubpassContents contents);
void fn_vkCmdNextSubpass2(VkCommandBuffer commandBuffer, const VkSubpassBeginInfo* pSubpassBeginInfo, const VkSubpassEndInfo* pSubpassEndInfo);
void fn_vkCmdPipelineBarrier(VkCommandBuffer commandBuffer, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, VkDependencyFlags dependencyFlags, uint32_t memoryBarrierCount, const VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, const VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, const VkImageMemoryBarrier* pImageMemoryBarriers);
void fn_vkCmdPipelineBarrier2(VkCommandBuffer commandBuffer, const VkDependencyInfo* pDependencyInfo);
void fn_vkCmdPreprocessGeneratedCommandsEXT(VkCommandBuffer commandBuffer, const VkGeneratedCommandsInfoEXT* pGeneratedCommandsInfo, VkCommandBuffer stateCommandBuffer);
void fn_vkCmdPushConstants(VkCommandBuffer commandBuffer, VkPipelineLayout layout, VkShaderStageFlags stageFlags, uint32_t offset, uint32_t size, const void* pValues);
void fn_vkCmdPushConstants2(VkCommandBuffer commandBuffer, const VkPushConstantsInfo* pPushConstantsInfo);
void fn_vkCmdPushDescriptorSet(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t set, uint32_t descriptorWriteCount, const VkWriteDescriptorSet* pDescriptorWrites);
void fn_vkCmdPushDescriptorSet2(VkCommandBuffer commandBuffer, const VkPushDescriptorSetInfo* pPushDescriptorSetInfo);
void fn_vkCmdPushDescriptorSetWithTemplate(VkCommandBuffer commandBuffer, VkDescriptorUpdateTemplate descriptorUpdateTemplate, VkPipelineLayout layout, uint32_t set, const void* pData);
void fn_vkCmdPushDescriptorSetWithTemplate2(VkCommandBuffer commandBuffer, const VkPushDescriptorSetWithTemplateInfo* pPushDescriptorSetWithTemplateInfo);
void fn_vkCmdResetEvent(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask);
void fn_vkCmdResetEvent2(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags2 stageMask);
void fn_vkCmdResetQueryPool(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount);
void fn_vkCmdResolveImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkImageResolve* pRegions);
void fn_vkCmdResolveImage2(VkCommandBuffer commandBuffer, const VkResolveImageInfo2* pResolveImageInfo);
void fn_vkCmdSetAlphaToCoverageEnableEXT(VkCommandBuffer commandBuffer, VkBool32 alphaToCoverageEnable);
void fn_vkCmdSetAlphaToOneEnableEXT(VkCommandBuffer commandBuffer, VkBool32 alphaToOneEnable);
void fn_vkCmdSetBlendConstants(VkCommandBuffer commandBuffer, const float* blendConstants);
void fn_vkCmdSetColorBlendAdvancedEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, const VkColorBlendAdvancedEXT* pColorBlendAdvanced);
void fn_vkCmdSetColorBlendEnableEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, const VkBool32* pColorBlendEnables);
void fn_vkCmdSetColorBlendEquationEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, const VkColorBlendEquationEXT* pColorBlendEquations);
void fn_vkCmdSetColorWriteMaskEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, const VkColorComponentFlags* pColorWriteMasks);
void fn_vkCmdSetConservativeRasterizationModeEXT(VkCommandBuffer commandBuffer, VkConservativeRasterizationModeEXT conservativeRasterizationMode);
void fn_vkCmdSetCoverageModulationModeNV(VkCommandBuffer commandBuffer, VkCoverageModulationModeNV coverageModulationMode);
void fn_vkCmdSetCoverageModulationTableEnableNV(VkCommandBuffer commandBuffer, VkBool32 coverageModulationTableEnable);
void fn_vkCmdSetCoverageModulationTableNV(VkCommandBuffer commandBuffer, uint32_t coverageModulationTableCount, const float* pCoverageModulationTable);
void fn_vkCmdSetCoverageReductionModeNV(VkCommandBuffer commandBuffer, VkCoverageReductionModeNV coverageReductionMode);
void fn_vkCmdSetCoverageToColorEnableNV(VkCommandBuffer commandBuffer, VkBool32 coverageToColorEnable);
void fn_vkCmdSetCoverageToColorLocationNV(VkCommandBuffer commandBuffer, uint32_t coverageToColorLocation);
void fn_vkCmdSetCullMode(VkCommandBuffer commandBuffer, VkCullModeFlags cullMode);
void fn_vkCmdSetDepthBias(VkCommandBuffer commandBuffer, float depthBiasConstantFactor, float depthBiasClamp, float depthBiasSlopeFactor);
void fn_vkCmdSetDepthBiasEnable(VkCommandBuffer commandBuffer, VkBool32 depthBiasEnable);
void fn_vkCmdSetDepthBounds(VkCommandBuffer commandBuffer, float minDepthBounds, float maxDepthBounds);
void fn_vkCmdSetDepthBoundsTestEnable(VkCommandBuffer commandBuffer, VkBool32 depthBoundsTestEnable);
void fn_vkCmdSetDepthClampEnableEXT(VkCommandBuffer commandBuffer, VkBool32 depthClampEnable);
void fn_vkCmdSetDepthClampRangeEXT(VkCommandBuffer commandBuffer, VkDepthClampModeEXT depthClampMode, const VkDepthClampRangeEXT* pDepthClampRange);
void fn_vkCmdSetDepthClipEnableEXT(VkCommandBuffer commandBuffer, VkBool32 depthClipEnable);
void fn_vkCmdSetDepthClipNegativeOneToOneEXT(VkCommandBuffer commandBuffer, VkBool32 negativeOneToOne);
void fn_vkCmdSetDepthCompareOp(VkCommandBuffer commandBuffer, VkCompareOp depthCompareOp);
void fn_vkCmdSetDepthTestEnable(VkCommandBuffer commandBuffer, VkBool32 depthTestEnable);
void fn_vkCmdSetDepthWriteEnable(VkCommandBuffer commandBuffer, VkBool32 depthWriteEnable);
void fn_vkCmdSetDeviceMask(VkCommandBuffer commandBuffer, uint32_t deviceMask);
void fn_vkCmdSetEvent(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask);
void fn_vkCmdSetEvent2(VkCommandBuffer commandBuffer, VkEvent event, const VkDependencyInfo* pDependencyInfo);
void fn_vkCmdSetExtraPrimitiveOverestimationSizeEXT(VkCommandBuffer commandBuffer, float extraPrimitiveOverestimationSize);
void fn_vkCmdSetFrontFace(VkCommandBuffer commandBuffer, VkFrontFace frontFace);
void fn_vkCmdSetLineRasterizationModeEXT(VkCommandBuffer commandBuffer, VkLineRasterizationModeEXT lineRasterizationMode);
void fn_vkCmdSetLineStipple(VkCommandBuffer commandBuffer, uint32_t lineStippleFactor, uint16_t lineStipplePattern);
void fn_vkCmdSetLineStippleEnableEXT(VkCommandBuffer commandBuffer, VkBool32 stippledLineEnable);
void fn_vkCmdSetLineWidth(VkCommandBuffer commandBuffer, float lineWidth);
void fn_vkCmdSetLogicOpEXT(VkCommandBuffer commandBuffer, VkLogicOp logicOp);
void fn_vkCmdSetLogicOpEnableEXT(VkCommandBuffer commandBuffer, VkBool32 logicOpEnable);
void fn_vkCmdSetPatchControlPointsEXT(VkCommandBuffer commandBuffer, uint32_t patchControlPoints);
void fn_vkCmdSetPolygonModeEXT(VkCommandBuffer commandBuffer, VkPolygonMode polygonMode);
void fn_vkCmdSetPrimitiveRestartEnable(VkCommandBuffer commandBuffer, VkBool32 primitiveRestartEnable);
void fn_vkCmdSetPrimitiveTopology(VkCommandBuffer commandBuffer, VkPrimitiveTopology primitiveTopology);
void fn_vkCmdSetProvokingVertexModeEXT(VkCommandBuffer commandBuffer, VkProvokingVertexModeEXT provokingVertexMode);
void fn_vkCmdSetRasterizationSamplesEXT(VkCommandBuffer commandBuffer, VkSampleCountFlagBits rasterizationSamples);
void fn_vkCmdSetRasterizationStreamEXT(VkCommandBuffer commandBuffer, uint32_t rasterizationStream);
void fn_vkCmdSetRasterizerDiscardEnable(VkCommandBuffer commandBuffer, VkBool32 rasterizerDiscardEnable);
void fn_vkCmdSetRayTracingPipelineStackSizeKHR(VkCommandBuffer commandBuffer, uint32_t pipelineStackSize);
void fn_vkCmdSetRenderingAttachmentLocations(VkCommandBuffer commandBuffer, const VkRenderingAttachmentLocationInfo* pLocationInfo);
void fn_vkCmdSetRenderingInputAttachmentIndices(VkCommandBuffer commandBuffer, const VkRenderingInputAttachmentIndexInfo* pInputAttachmentIndexInfo);
void fn_vkCmdSetRepresentativeFragmentTestEnableNV(VkCommandBuffer commandBuffer, VkBool32 representativeFragmentTestEnable);
void fn_vkCmdSetSampleLocationsEnableEXT(VkCommandBuffer commandBuffer, VkBool32 sampleLocationsEnable);
void fn_vkCmdSetSampleMaskEXT(VkCommandBuffer commandBuffer, VkSampleCountFlagBits samples, const VkSampleMask* pSampleMask);
void fn_vkCmdSetScissor(VkCommandBuffer commandBuffer, uint32_t firstScissor, uint32_t scissorCount, const VkRect2D* pScissors);
void fn_vkCmdSetScissorWithCount(VkCommandBuffer commandBuffer, uint32_t scissorCount, const VkRect2D* pScissors);
void fn_vkCmdSetShadingRateImageEnableNV(VkCommandBuffer commandBuffer, VkBool32 shadingRateImageEnable);
void fn_vkCmdSetStencilCompareMask(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t compareMask);
void fn_vkCmdSetStencilOp(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, VkStencilOp failOp, VkStencilOp passOp, VkStencilOp depthFailOp, VkCompareOp compareOp);
void fn_vkCmdSetStencilReference(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t reference);
void fn_vkCmdSetStencilTestEnable(VkCommandBuffer commandBuffer, VkBool32 stencilTestEnable);
void fn_vkCmdSetStencilWriteMask(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t writeMask);
void fn_vkCmdSetTessellationDomainOriginEXT(VkCommandBuffer commandBuffer, VkTessellationDomainOrigin domainOrigin);
void fn_vkCmdSetVertexInputEXT(VkCommandBuffer commandBuffer, uint32_t vertexBindingDescriptionCount, const VkVertexInputBindingDescription2EXT* pVertexBindingDescriptions, uint32_t vertexAttributeDescriptionCount, const VkVertexInputAttributeDescription2EXT* pVertexAttributeDescriptions);
void fn_vkCmdSetViewport(VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, const VkViewport* pViewports);
void fn_vkCmdSetViewportSwizzleNV(VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, const VkViewportSwizzleNV* pViewportSwizzles);
void fn_vkCmdSetViewportWScalingEnableNV(VkCommandBuffer commandBuffer, VkBool32 viewportWScalingEnable);
void fn_vkCmdSetViewportWithCount(VkCommandBuffer commandBuffer, uint32_t viewportCount, const VkViewport* pViewports);
void fn_vkCmdTraceRaysIndirect2KHR(VkCommandBuffer commandBuffer, VkDeviceAddress indirectDeviceAddress);
void fn_vkCmdTraceRaysIndirectKHR(VkCommandBuffer commandBuffer, const VkStridedDeviceAddressRegionKHR* pRaygenShaderBindingTable, const VkStridedDeviceAddressRegionKHR* pMissShaderBindingTable, const VkStridedDeviceAddressRegionKHR* pHitShaderBindingTable, const VkStridedDeviceAddressRegionKHR* pCallableShaderBindingTable, VkDeviceAddress indirectDeviceAddress);
void fn_vkCmdTraceRaysKHR(VkCommandBuffer commandBuffer, const VkStridedDeviceAddressRegionKHR* pRaygenShaderBindingTable, const VkStridedDeviceAddressRegionKHR* pMissShaderBindingTable, const VkStridedDeviceAddressRegionKHR* pHitShaderBindingTable, const VkStridedDeviceAddressRegionKHR* pCallableShaderBindingTable, uint32_t width, uint32_t height, uint32_t depth);
void fn_vkCmdUpdateBuffer(VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize dataSize, const void* pData);
void fn_vkCmdWaitEvents(VkCommandBuffer commandBuffer, uint32_t eventCount, const VkEvent* pEvents, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, uint32_t memoryBarrierCount, const VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, const VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, const VkImageMemoryBarrier* pImageMemoryBarriers);
void fn_vkCmdWaitEvents2(VkCommandBuffer commandBuffer, uint32_t eventCount, const VkEvent* pEvents, const VkDependencyInfo* pDependencyInfos);
void fn_vkCmdWriteAccelerationStructuresPropertiesKHR(VkCommandBuffer commandBuffer, uint32_t accelerationStructureCount, const VkAccelerationStructureKHR* pAccelerationStructures, VkQueryType queryType, VkQueryPool queryPool, uint32_t firstQuery);
void fn_vkCmdWriteTimestamp(VkCommandBuffer commandBuffer, VkPipelineStageFlagBits pipelineStage, VkQueryPool queryPool, uint32_t query);
void fn_vkCmdWriteTimestamp2(VkCommandBuffer commandBuffer, VkPipelineStageFlags2 stage, VkQueryPool queryPool, uint32_t query);
VkResult fn_vkCopyAccelerationStructureKHR(VkDevice device, VkDeferredOperationKHR deferredOperation, const VkCopyAccelerationStructureInfoKHR* pInfo);
VkResult fn_vkCopyAccelerationStructureToMemoryKHR(VkDevice device, VkDeferredOperationKHR deferredOperation, const VkCopyAccelerationStructureToMemoryInfoKHR* pInfo);
VkResult fn_vkCopyImageToImage(VkDevice device, const VkCopyImageToImageInfo* pCopyImageToImageInfo);
VkResult fn_vkCopyImageToMemory(VkDevice device, const VkCopyImageToMemoryInfo* pCopyImageToMemoryInfo);
VkResult fn_vkCopyMemoryToAccelerationStructureKHR(VkDevice device, VkDeferredOperationKHR deferredOperation, const VkCopyMemoryToAccelerationStructureInfoKHR* pInfo);
VkResult fn_vkCopyMemoryToImage(VkDevice device, const VkCopyMemoryToImageInfo* pCopyMemoryToImageInfo);
VkResult fn_vkCreateAccelerationStructureKHR(VkDevice device, const VkAccelerationStructureCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkAccelerationStructureKHR* pAccelerationStructure);
VkResult fn_vkCreateBuffer(VkDevice device, const VkBufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkBuffer* pBuffer);
VkResult fn_vkCreateBufferView(VkDevice device, const VkBufferViewCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkBufferView* pView);
VkResult fn_vkCreateCommandPool(VkDevice device, const VkCommandPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkCommandPool* pCommandPool);
VkResult fn_vkCreateComputePipelines(VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkComputePipelineCreateInfo* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines);
VkResult fn_vkCreateDebugUtilsMessengerEXT(VkInstance instance, const VkDebugUtilsMessengerCreateInfoEXT* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDebugUtilsMessengerEXT* pMessenger);
VkResult fn_vkCreateDeferredOperationKHR(VkDevice device, const VkAllocationCallbacks* pAllocator, VkDeferredOperationKHR* pDeferredOperation);
VkResult fn_vkCreateDescriptorPool(VkDevice device, const VkDescriptorPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDescriptorPool* pDescriptorPool);
VkResult fn_vkCreateDescriptorSetLayout(VkDevice device, const VkDescriptorSetLayoutCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDescriptorSetLayout* pSetLayout);
VkResult fn_vkCreateDescriptorUpdateTemplate(VkDevice device, const VkDescriptorUpdateTemplateCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDescriptorUpdateTemplate* pDescriptorUpdateTemplate);
VkResult fn_vkCreateDevice(VkPhysicalDevice physicalDevice, const VkDeviceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDevice* pDevice);
VkResult fn_vkCreateEvent(VkDevice device, const VkEventCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkEvent* pEvent);
VkResult fn_vkCreateFence(VkDevice device, const VkFenceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFence* pFence);
VkResult fn_vkCreateFramebuffer(VkDevice device, const VkFramebufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFramebuffer* pFramebuffer);
VkResult fn_vkCreateGraphicsPipelines(VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkGraphicsPipelineCreateInfo* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines);
VkResult fn_vkCreateImage(VkDevice device, const VkImageCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkImage* pImage);
VkResult fn_vkCreateImageView(VkDevice device, const VkImageViewCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkImageView* pView);
VkResult fn_vkCreateIndirectCommandsLayoutEXT(VkDevice device, const VkIndirectCommandsLayoutCreateInfoEXT* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkIndirectCommandsLayoutEXT* pIndirectCommandsLayout);
VkResult fn_vkCreateIndirectExecutionSetEXT(VkDevice device, const VkIndirectExecutionSetCreateInfoEXT* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkIndirectExecutionSetEXT* pIndirectExecutionSet);
VkResult fn_vkCreateInstance(const VkInstanceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkInstance* pInstance);
VkResult fn_vkCreatePipelineCache(VkDevice device, const VkPipelineCacheCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkPipelineCache* pPipelineCache);
VkResult fn_vkCreatePipelineLayout(VkDevice device, const VkPipelineLayoutCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkPipelineLayout* pPipelineLayout);
VkResult fn_vkCreatePrivateDataSlot(VkDevice device, const VkPrivateDataSlotCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkPrivateDataSlot* pPrivateDataSlot);
VkResult fn_vkCreateQueryPool(VkDevice device, const VkQueryPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkQueryPool* pQueryPool);
VkResult fn_vkCreateRayTracingPipelinesKHR(VkDevice device, VkDeferredOperationKHR deferredOperation, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkRayTracingPipelineCreateInfoKHR* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines);
VkResult fn_vkCreateRenderPass(VkDevice device, const VkRenderPassCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass);
VkResult fn_vkCreateRenderPass2(VkDevice device, const VkRenderPassCreateInfo2* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass);
VkResult fn_vkCreateSampler(VkDevice device, const VkSamplerCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSampler* pSampler);
VkResult fn_vkCreateSamplerYcbcrConversion(VkDevice device, const VkSamplerYcbcrConversionCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSamplerYcbcrConversion* pYcbcrConversion);
VkResult fn_vkCreateSemaphore(VkDevice device, const VkSemaphoreCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSemaphore* pSemaphore);
VkResult fn_vkCreateShaderModule(VkDevice device, const VkShaderModuleCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkShaderModule* pShaderModule);
VkResult fn_vkCreateShadersEXT(VkDevice device, uint32_t createInfoCount, const VkShaderCreateInfoEXT* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkShaderEXT* pShaders);
VkResult fn_vkCreateSwapchainKHR(VkDevice device, const VkSwapchainCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSwapchainKHR* pSwapchain);
VkResult fn_vkDeferredOperationJoinKHR(VkDevice device, VkDeferredOperationKHR operation);
void fn_vkDestroyAccelerationStructureKHR(VkDevice device, VkAccelerationStructureKHR accelerationStructure, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyBuffer(VkDevice device, VkBuffer buffer, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyBufferView(VkDevice device, VkBufferView bufferView, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyCommandPool(VkDevice device, VkCommandPool commandPool, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyDebugUtilsMessengerEXT(VkInstance instance, VkDebugUtilsMessengerEXT messenger, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyDeferredOperationKHR(VkDevice device, VkDeferredOperationKHR operation, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyDescriptorPool(VkDevice device, VkDescriptorPool descriptorPool, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyDescriptorSetLayout(VkDevice device, VkDescriptorSetLayout descriptorSetLayout, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyDescriptorUpdateTemplate(VkDevice device, VkDescriptorUpdateTemplate descriptorUpdateTemplate, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyDevice(VkDevice device, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyEvent(VkDevice device, VkEvent event, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyFence(VkDevice device, VkFence fence, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyFramebuffer(VkDevice device, VkFramebuffer framebuffer, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyImage(VkDevice device, VkImage image, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyImageView(VkDevice device, VkImageView imageView, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyIndirectCommandsLayoutEXT(VkDevice device, VkIndirectCommandsLayoutEXT indirectCommandsLayout, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyIndirectExecutionSetEXT(VkDevice device, VkIndirectExecutionSetEXT indirectExecutionSet, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyInstance(VkInstance instance, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyPipeline(VkDevice device, VkPipeline pipeline, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyPipelineCache(VkDevice device, VkPipelineCache pipelineCache, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyPipelineLayout(VkDevice device, VkPipelineLayout pipelineLayout, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyPrivateDataSlot(VkDevice device, VkPrivateDataSlot privateDataSlot, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyQueryPool(VkDevice device, VkQueryPool queryPool, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyRenderPass(VkDevice device, VkRenderPass renderPass, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroySampler(VkDevice device, VkSampler sampler, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroySamplerYcbcrConversion(VkDevice device, VkSamplerYcbcrConversion ycbcrConversion, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroySemaphore(VkDevice device, VkSemaphore semaphore, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyShaderEXT(VkDevice device, VkShaderEXT shader, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroyShaderModule(VkDevice device, VkShaderModule shaderModule, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroySurfaceKHR(VkInstance instance, VkSurfaceKHR surface, const VkAllocationCallbacks* pAllocator);
void fn_vkDestroySwapchainKHR(VkDevice device, VkSwapchainKHR swapchain, const VkAllocationCallbacks* pAllocator);
VkResult fn_vkDeviceWaitIdle(VkDevice device);
VkResult fn_vkEndCommandBuffer(VkCommandBuffer commandBuffer);
VkResult fn_vkEnumerateDeviceExtensionProperties(VkPhysicalDevice physicalDevice, const char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties);
VkResult fn_vkEnumerateDeviceLayerProperties(VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkLayerProperties* pProperties);
VkResult fn_vkEnumerateInstanceExtensionProperties(const char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties);
VkResult fn_vkEnumerateInstanceLayerProperties(uint32_t* pPropertyCount, VkLayerProperties* pProperties);
VkResult fn_vkEnumerateInstanceVersion(uint32_t* pApiVersion);
VkResult fn_vkEnumeratePhysicalDeviceGroups(VkInstance instance, uint32_t* pPhysicalDeviceGroupCount, VkPhysicalDeviceGroupProperties* pPhysicalDeviceGroupProperties);
VkResult fn_vkEnumeratePhysicalDevices(VkInstance instance, uint32_t* pPhysicalDeviceCount, VkPhysicalDevice* pPhysicalDevices);
VkResult fn_vkFlushMappedMemoryRanges(VkDevice device, uint32_t memoryRangeCount, const VkMappedMemoryRange* pMemoryRanges);
void fn_vkFreeCommandBuffers(VkDevice device, VkCommandPool commandPool, uint32_t commandBufferCount, const VkCommandBuffer* pCommandBuffers);
VkResult fn_vkFreeDescriptorSets(VkDevice device, VkDescriptorPool descriptorPool, uint32_t descriptorSetCount, const VkDescriptorSet* pDescriptorSets);
void fn_vkFreeMemory(VkDevice device, VkDeviceMemory memory, const VkAllocationCallbacks* pAllocator);
void fn_vkGetAccelerationStructureBuildSizesKHR(VkDevice device, VkAccelerationStructureBuildTypeKHR buildType, const VkAccelerationStructureBuildGeometryInfoKHR* pBuildInfo, const uint32_t* pMaxPrimitiveCounts, VkAccelerationStructureBuildSizesInfoKHR* pSizeInfo);
VkDeviceAddress fn_vkGetAccelerationStructureDeviceAddressKHR(VkDevice device, const VkAccelerationStructureDeviceAddressInfoKHR* pInfo);
VkDeviceAddress fn_vkGetBufferDeviceAddress(VkDevice device, const VkBufferDeviceAddressInfo* pInfo);
void fn_vkGetBufferMemoryRequirements(VkDevice device, VkBuffer buffer, VkMemoryRequirements* pMemoryRequirements);
void fn_vkGetBufferMemoryRequirements2(VkDevice device, const VkBufferMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements);
uint64_t fn_vkGetBufferOpaqueCaptureAddress(VkDevice device, const VkBufferDeviceAddressInfo* pInfo);
uint32_t fn_vkGetDeferredOperationMaxConcurrencyKHR(VkDevice device, VkDeferredOperationKHR operation);
VkResult fn_vkGetDeferredOperationResultKHR(VkDevice device, VkDeferredOperationKHR operation);
void fn_vkGetDescriptorSetLayoutSupport(VkDevice device, const VkDescriptorSetLayoutCreateInfo* pCreateInfo, VkDescriptorSetLayoutSupport* pSupport);
void fn_vkGetDeviceAccelerationStructureCompatibilityKHR(VkDevice device, const VkAccelerationStructureVersionInfoKHR* pVersionInfo, VkAccelerationStructureCompatibilityKHR* pCompatibility);
void fn_vkGetDeviceBufferMemoryRequirements(VkDevice device, const VkDeviceBufferMemoryRequirements* pInfo, VkMemoryRequirements2* pMemoryRequirements);
void fn_vkGetDeviceGroupPeerMemoryFeatures(VkDevice device, uint32_t heapIndex, uint32_t localDeviceIndex, uint32_t remoteDeviceIndex, VkPeerMemoryFeatureFlags* pPeerMemoryFeatures);
VkResult fn_vkGetDeviceGroupPresentCapabilitiesKHR(VkDevice device, VkDeviceGroupPresentCapabilitiesKHR* pDeviceGroupPresentCapabilities);
VkResult fn_vkGetDeviceGroupSurfacePresentModesKHR(VkDevice device, VkSurfaceKHR surface, VkDeviceGroupPresentModeFlagsKHR* pModes);
void fn_vkGetDeviceImageMemoryRequirements(VkDevice device, const VkDeviceImageMemoryRequirements* pInfo, VkMemoryRequirements2* pMemoryRequirements);
void fn_vkGetDeviceImageSparseMemoryRequirements(VkDevice device, const VkDeviceImageMemoryRequirements* pInfo, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2* pSparseMemoryRequirements);
void fn_vkGetDeviceImageSubresourceLayout(VkDevice device, const VkDeviceImageSubresourceInfo* pInfo, VkSubresourceLayout2* pLayout);
void fn_vkGetDeviceMemoryCommitment(VkDevice device, VkDeviceMemory memory, VkDeviceSize* pCommittedMemoryInBytes);
uint64_t fn_vkGetDeviceMemoryOpaqueCaptureAddress(VkDevice device, const VkDeviceMemoryOpaqueCaptureAddressInfo* pInfo);
PFN_vkVoidFunction fn_vkGetDeviceProcAddr(VkDevice device, const char* pName);
void fn_vkGetDeviceQueue(VkDevice device, uint32_t queueFamilyIndex, uint32_t queueIndex, VkQueue* pQueue);
void fn_vkGetDeviceQueue2(VkDevice device, const VkDeviceQueueInfo2* pQueueInfo, VkQueue* pQueue);
VkResult fn_vkGetEventStatus(VkDevice device, VkEvent event);
VkResult fn_vkGetFenceStatus(VkDevice device, VkFence fence);
void fn_vkGetGeneratedCommandsMemoryRequirementsEXT(VkDevice device, const VkGeneratedCommandsMemoryRequirementsInfoEXT* pInfo, VkMemoryRequirements2* pMemoryRequirements);
void fn_vkGetImageMemoryRequirements(VkDevice device, VkImage image, VkMemoryRequirements* pMemoryRequirements);
void fn_vkGetImageMemoryRequirements2(VkDevice device, const VkImageMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements);
void fn_vkGetImageSparseMemoryRequirements(VkDevice device, VkImage image, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements* pSparseMemoryRequirements);
void fn_vkGetImageSparseMemoryRequirements2(VkDevice device, const VkImageSparseMemoryRequirementsInfo2* pInfo, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2* pSparseMemoryRequirements);
void fn_vkGetImageSubresourceLayout(VkDevice device, VkImage image, const VkImageSubresource* pSubresource, VkSubresourceLayout* pLayout);
void fn_vkGetImageSubresourceLayout2(VkDevice device, VkImage image, const VkImageSubresource2* pSubresource, VkSubresourceLayout2* pLayout);
PFN_vkVoidFunction fn_vkGetInstanceProcAddr(VkInstance instance, const char* pName);
void fn_vkGetPhysicalDeviceExternalBufferProperties(VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalBufferInfo* pExternalBufferInfo, VkExternalBufferProperties* pExternalBufferProperties);
void fn_vkGetPhysicalDeviceExternalFenceProperties(VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalFenceInfo* pExternalFenceInfo, VkExternalFenceProperties* pExternalFenceProperties);
void fn_vkGetPhysicalDeviceExternalSemaphoreProperties(VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalSemaphoreInfo* pExternalSemaphoreInfo, VkExternalSemaphoreProperties* pExternalSemaphoreProperties);
void fn_vkGetPhysicalDeviceFeatures(VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures* pFeatures);
void fn_vkGetPhysicalDeviceFeatures2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures2* pFeatures);
void fn_vkGetPhysicalDeviceFormatProperties(VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties* pFormatProperties);
void fn_vkGetPhysicalDeviceFormatProperties2(VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties2* pFormatProperties);
VkResult fn_vkGetPhysicalDeviceImageFormatProperties(VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkImageTiling tiling, VkImageUsageFlags usage, VkImageCreateFlags flags, VkImageFormatProperties* pImageFormatProperties);
VkResult fn_vkGetPhysicalDeviceImageFormatProperties2(VkPhysicalDevice physicalDevice, const VkPhysicalDeviceImageFormatInfo2* pImageFormatInfo, VkImageFormatProperties2* pImageFormatProperties);
void fn_vkGetPhysicalDeviceMemoryProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties* pMemoryProperties);
void fn_vkGetPhysicalDeviceMemoryProperties2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties2* pMemoryProperties);
VkResult fn_vkGetPhysicalDevicePresentRectanglesKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pRectCount, VkRect2D* pRects);
void fn_vkGetPhysicalDeviceProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties* pProperties);
void fn_vkGetPhysicalDeviceProperties2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties2* pProperties);
void fn_vkGetPhysicalDeviceQueueFamilyProperties(VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties* pQueueFamilyProperties);
void fn_vkGetPhysicalDeviceQueueFamilyProperties2(VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties2* pQueueFamilyProperties);
void fn_vkGetPhysicalDeviceSparseImageFormatProperties(VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkSampleCountFlagBits samples, VkImageUsageFlags usage, VkImageTiling tiling, uint32_t* pPropertyCount, VkSparseImageFormatProperties* pProperties);
void fn_vkGetPhysicalDeviceSparseImageFormatProperties2(VkPhysicalDevice physicalDevice, const VkPhysicalDeviceSparseImageFormatInfo2* pFormatInfo, uint32_t* pPropertyCount, VkSparseImageFormatProperties2* pProperties);
VkResult fn_vkGetPhysicalDeviceSurfaceCapabilitiesKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, VkSurfaceCapabilitiesKHR* pSurfaceCapabilities);
VkResult fn_vkGetPhysicalDeviceSurfaceFormatsKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pSurfaceFormatCount, VkSurfaceFormatKHR* pSurfaceFormats);
VkResult fn_vkGetPhysicalDeviceSurfacePresentModesKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pPresentModeCount, VkPresentModeKHR* pPresentModes);
VkResult fn_vkGetPhysicalDeviceSurfaceSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, VkSurfaceKHR surface, VkBool32* pSupported);
VkResult fn_vkGetPhysicalDeviceToolProperties(VkPhysicalDevice physicalDevice, uint32_t* pToolCount, VkPhysicalDeviceToolProperties* pToolProperties);
VkResult fn_vkGetPipelineCacheData(VkDevice device, VkPipelineCache pipelineCache, size_t* pDataSize, void* pData);
void fn_vkGetPrivateData(VkDevice device, VkObjectType objectType, uint64_t objectHandle, VkPrivateDataSlot privateDataSlot, uint64_t* pData);
VkResult fn_vkGetQueryPoolResults(VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, size_t dataSize, void* pData, VkDeviceSize stride, VkQueryResultFlags flags);
VkResult fn_vkGetRayTracingCaptureReplayShaderGroupHandlesKHR(VkDevice device, VkPipeline pipeline, uint32_t firstGroup, uint32_t groupCount, size_t dataSize, void* pData);
VkResult fn_vkGetRayTracingShaderGroupHandlesKHR(VkDevice device, VkPipeline pipeline, uint32_t firstGroup, uint32_t groupCount, size_t dataSize, void* pData);
VkDeviceSize fn_vkGetRayTracingShaderGroupStackSizeKHR(VkDevice device, VkPipeline pipeline, uint32_t group, VkShaderGroupShaderKHR groupShader);
void fn_vkGetRenderAreaGranularity(VkDevice device, VkRenderPass renderPass, VkExtent2D* pGranularity);
void fn_vkGetRenderingAreaGranularity(VkDevice device, const VkRenderingAreaInfo* pRenderingAreaInfo, VkExtent2D* pGranularity);
VkResult fn_vkGetSemaphoreCounterValue(VkDevice device, VkSemaphore semaphore, uint64_t* pValue);
VkResult fn_vkGetShaderBinaryDataEXT(VkDevice device, VkShaderEXT shader, size_t* pDataSize, void* pData);
VkResult fn_vkGetSwapchainImagesKHR(VkDevice device, VkSwapchainKHR swapchain, uint32_t* pSwapchainImageCount, VkImage* pSwapchainImages);
VkResult fn_vkInvalidateMappedMemoryRanges(VkDevice device, uint32_t memoryRangeCount, const VkMappedMemoryRange* pMemoryRanges);
VkResult fn_vkMapMemory(VkDevice device, VkDeviceMemory memory, VkDeviceSize offset, VkDeviceSize size, VkMemoryMapFlags flags, void** ppData);
VkResult fn_vkMapMemory2(VkDevice device, const VkMemoryMapInfo* pMemoryMapInfo, void** ppData);
VkResult fn_vkMergePipelineCaches(VkDevice device, VkPipelineCache dstCache, uint32_t srcCacheCount, const VkPipelineCache* pSrcCaches);
void fn_vkQueueBeginDebugUtilsLabelEXT(VkQueue queue, const VkDebugUtilsLabelEXT* pLabelInfo);
VkResult fn_vkQueueBindSparse(VkQueue queue, uint32_t bindInfoCount, const VkBindSparseInfo* pBindInfo, VkFence fence);
void fn_vkQueueEndDebugUtilsLabelEXT(VkQueue queue);
void fn_vkQueueInsertDebugUtilsLabelEXT(VkQueue queue, const VkDebugUtilsLabelEXT* pLabelInfo);
VkResult fn_vkQueuePresentKHR(VkQueue queue, const VkPresentInfoKHR* pPresentInfo);
VkResult fn_vkQueueSubmit(VkQueue queue, uint32_t submitCount, const VkSubmitInfo* pSubmits, VkFence fence);
VkResult fn_vkQueueSubmit2(VkQueue queue, uint32_t submitCount, const VkSubmitInfo2* pSubmits, VkFence fence);
VkResult fn_vkQueueWaitIdle(VkQueue queue);
VkResult fn_vkResetCommandBuffer(VkCommandBuffer commandBuffer, VkCommandBufferResetFlags flags);
VkResult fn_vkResetCommandPool(VkDevice device, VkCommandPool commandPool, VkCommandPoolResetFlags flags);
VkResult fn_vkResetDescriptorPool(VkDevice device, VkDescriptorPool descriptorPool, VkDescriptorPoolResetFlags flags);
VkResult fn_vkResetEvent(VkDevice device, VkEvent event);
VkResult fn_vkResetFences(VkDevice device, uint32_t fenceCount, const VkFence* pFences);
void fn_vkResetQueryPool(VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount);
VkResult fn_vkSetDebugUtilsObjectNameEXT(VkDevice device, const VkDebugUtilsObjectNameInfoEXT* pNameInfo);
VkResult fn_vkSetDebugUtilsObjectTagEXT(VkDevice device, const VkDebugUtilsObjectTagInfoEXT* pTagInfo);
VkResult fn_vkSetEvent(VkDevice device, VkEvent event);
VkResult fn_vkSetPrivateData(VkDevice device, VkObjectType objectType, uint64_t objectHandle, VkPrivateDataSlot privateDataSlot, uint64_t data);
VkResult fn_vkSignalSemaphore(VkDevice device, const VkSemaphoreSignalInfo* pSignalInfo);
void fn_vkSubmitDebugUtilsMessageEXT(VkInstance instance, VkDebugUtilsMessageSeverityFlagBitsEXT messageSeverity, VkDebugUtilsMessageTypeFlagsEXT messageTypes, const VkDebugUtilsMessengerCallbackDataEXT* pCallbackData);
VkResult fn_vkTransitionImageLayout(VkDevice device, uint32_t transitionCount, const VkHostImageLayoutTransitionInfo* pTransitions);
void fn_vkTrimCommandPool(VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlags flags);
void fn_vkUnmapMemory(VkDevice device, VkDeviceMemory memory);
VkResult fn_vkUnmapMemory2(VkDevice device, const VkMemoryUnmapInfo* pMemoryUnmapInfo);
void fn_vkUpdateDescriptorSetWithTemplate(VkDevice device, VkDescriptorSet descriptorSet, VkDescriptorUpdateTemplate descriptorUpdateTemplate, const void* pData);
void fn_vkUpdateDescriptorSets(VkDevice device, uint32_t descriptorWriteCount, const VkWriteDescriptorSet* pDescriptorWrites, uint32_t descriptorCopyCount, const VkCopyDescriptorSet* pDescriptorCopies);
void fn_vkUpdateIndirectExecutionSetPipelineEXT(VkDevice device, VkIndirectExecutionSetEXT indirectExecutionSet, uint32_t executionSetWriteCount, const VkWriteIndirectExecutionSetPipelineEXT* pExecutionSetWrites);
void fn_vkUpdateIndirectExecutionSetShaderEXT(VkDevice device, VkIndirectExecutionSetEXT indirectExecutionSet, uint32_t executionSetWriteCount, const VkWriteIndirectExecutionSetShaderEXT* pExecutionSetWrites);
VkResult fn_vkWaitForFences(VkDevice device, uint32_t fenceCount, const VkFence* pFences, VkBool32 waitAll, uint64_t timeout);
VkResult fn_vkWaitSemaphores(VkDevice device, const VkSemaphoreWaitInfo* pWaitInfo, uint64_t timeout);
VkResult fn_vkWriteAccelerationStructuresPropertiesKHR(VkDevice device, uint32_t accelerationStructureCount, const VkAccelerationStructureKHR* pAccelerationStructures, VkQueryType queryType, size_t dataSize, void* pData, size_t stride);

#ifdef VK_USE_PLATFORM_MACOS_MVK
VkResult fn_vkCreateMacOSSurfaceMVK(VkInstance instance, const VkMacOSSurfaceCreateInfoMVK* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
#endif // VK_USE_PLATFORM_MACOS_MVK

#ifdef VK_USE_PLATFORM_WAYLAND_KHR
VkResult fn_vkCreateWaylandSurfaceKHR(VkInstance instance, const VkWaylandSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
VkBool32 fn_vkGetPhysicalDeviceWaylandPresentationSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, struct wl_display* display);
#endif // VK_USE_PLATFORM_WAYLAND_KHR

#ifdef VK_USE_PLATFORM_WIN32_KHR
VkResult fn_vkCreateWin32SurfaceKHR(VkInstance instance, const VkWin32SurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
VkBool32 fn_vkGetPhysicalDeviceWin32PresentationSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex);
#endif // VK_USE_PLATFORM_WIN32_KHR

#ifdef VK_USE_PLATFORM_XCB_KHR
VkResult fn_vkCreateXcbSurfaceKHR(VkInstance instance, const VkXcbSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
VkBool32 fn_vkGetPhysicalDeviceXcbPresentationSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, xcb_connection_t* connection, xcb_visualid_t visual_id);
#endif // VK_USE_PLATFORM_XCB_KHR

#ifdef VK_USE_PLATFORM_XLIB_KHR
VkResult fn_vkCreateXlibSurfaceKHR(VkInstance instance, const VkXlibSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
VkBool32 fn_vkGetPhysicalDeviceXlibPresentationSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, Display* dpy, VisualID visualID);
#endif // VK_USE_PLATFORM_XLIB_KHR

/* Callback trampolines */
extern void* go_bridge_AllocationCallbacks_PfnAllocation(void*, size_t, size_t, VkSystemAllocationScope);
void* trampoline_AllocationCallbacks_PfnAllocation(void*, size_t, size_t, VkSystemAllocationScope);
extern void* go_bridge_AllocationCallbacks_PfnReallocation(void*, void*, size_t, size_t, VkSystemAllocationScope);
void* trampoline_AllocationCallbacks_PfnReallocation(void*, void*, size_t, size_t, VkSystemAllocationScope);
extern void go_bridge_AllocationCallbacks_PfnFree(void*, void*);
void trampoline_AllocationCallbacks_PfnFree(void*, void*);
extern void go_bridge_AllocationCallbacks_PfnInternalAllocation(void*, size_t, VkInternalAllocationType, VkSystemAllocationScope);
void trampoline_AllocationCallbacks_PfnInternalAllocation(void*, size_t, VkInternalAllocationType, VkSystemAllocationScope);
extern void go_bridge_AllocationCallbacks_PfnInternalFree(void*, size_t, VkInternalAllocationType, VkSystemAllocationScope);
void trampoline_AllocationCallbacks_PfnInternalFree(void*, size_t, VkInternalAllocationType, VkSystemAllocationScope);
extern VkBool32 go_bridge_DebugUtilsMessengerCreateInfoEXT_PfnUserCallback(VkDebugUtilsMessageSeverityFlagBitsEXT, VkDebugUtilsMessageTypeFlagsEXT, VkDebugUtilsMessengerCallbackDataEXT*, void*);
VkBool32 trampoline_DebugUtilsMessengerCreateInfoEXT_PfnUserCallback(VkDebugUtilsMessageSeverityFlagBitsEXT, VkDebugUtilsMessageTypeFlagsEXT, VkDebugUtilsMessengerCallbackDataEXT*, void*);

#endif
