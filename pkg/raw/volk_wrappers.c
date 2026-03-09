#define VOLK_IMPLEMENTATION
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
#include <vulkan/vulkan.h>
#include "volk_wrappers.h"

VkResult fn_vkAcquireNextImage2KHR(VkDevice device, const VkAcquireNextImageInfoKHR* pAcquireInfo, uint32_t* pImageIndex) {
	return vkAcquireNextImage2KHR(device, pAcquireInfo, pImageIndex);
}

VkResult fn_vkAcquireNextImageKHR(VkDevice device, VkSwapchainKHR swapchain, uint64_t timeout, VkSemaphore semaphore, VkFence fence, uint32_t* pImageIndex) {
	return vkAcquireNextImageKHR(device, swapchain, timeout, semaphore, fence, pImageIndex);
}

VkResult fn_vkAllocateCommandBuffers(VkDevice device, const VkCommandBufferAllocateInfo* pAllocateInfo, VkCommandBuffer* pCommandBuffers) {
	return vkAllocateCommandBuffers(device, pAllocateInfo, pCommandBuffers);
}

VkResult fn_vkAllocateDescriptorSets(VkDevice device, const VkDescriptorSetAllocateInfo* pAllocateInfo, VkDescriptorSet* pDescriptorSets) {
	return vkAllocateDescriptorSets(device, pAllocateInfo, pDescriptorSets);
}

VkResult fn_vkAllocateMemory(VkDevice device, const VkMemoryAllocateInfo* pAllocateInfo, const VkAllocationCallbacks* pAllocator, VkDeviceMemory* pMemory) {
	return vkAllocateMemory(device, pAllocateInfo, pAllocator, pMemory);
}

VkResult fn_vkBeginCommandBuffer(VkCommandBuffer commandBuffer, const VkCommandBufferBeginInfo* pBeginInfo) {
	return vkBeginCommandBuffer(commandBuffer, pBeginInfo);
}

VkResult fn_vkBindBufferMemory(VkDevice device, VkBuffer buffer, VkDeviceMemory memory, VkDeviceSize memoryOffset) {
	return vkBindBufferMemory(device, buffer, memory, memoryOffset);
}

VkResult fn_vkBindBufferMemory2(VkDevice device, uint32_t bindInfoCount, const VkBindBufferMemoryInfo* pBindInfos) {
	return vkBindBufferMemory2(device, bindInfoCount, pBindInfos);
}

VkResult fn_vkBindImageMemory(VkDevice device, VkImage image, VkDeviceMemory memory, VkDeviceSize memoryOffset) {
	return vkBindImageMemory(device, image, memory, memoryOffset);
}

VkResult fn_vkBindImageMemory2(VkDevice device, uint32_t bindInfoCount, const VkBindImageMemoryInfo* pBindInfos) {
	return vkBindImageMemory2(device, bindInfoCount, pBindInfos);
}

void fn_vkCmdBeginDebugUtilsLabelEXT(VkCommandBuffer commandBuffer, const VkDebugUtilsLabelEXT* pLabelInfo) {
	vkCmdBeginDebugUtilsLabelEXT(commandBuffer, pLabelInfo);
}

void fn_vkCmdBeginQuery(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query, VkQueryControlFlags flags) {
	vkCmdBeginQuery(commandBuffer, queryPool, query, flags);
}

void fn_vkCmdBeginRenderPass(VkCommandBuffer commandBuffer, const VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassContents contents) {
	vkCmdBeginRenderPass(commandBuffer, pRenderPassBegin, contents);
}

void fn_vkCmdBeginRenderPass2(VkCommandBuffer commandBuffer, const VkRenderPassBeginInfo* pRenderPassBegin, const VkSubpassBeginInfo* pSubpassBeginInfo) {
	vkCmdBeginRenderPass2(commandBuffer, pRenderPassBegin, pSubpassBeginInfo);
}

void fn_vkCmdBeginRendering(VkCommandBuffer commandBuffer, const VkRenderingInfo* pRenderingInfo) {
	vkCmdBeginRendering(commandBuffer, pRenderingInfo);
}

void fn_vkCmdBindDescriptorSets(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t firstSet, uint32_t descriptorSetCount, const VkDescriptorSet* pDescriptorSets, uint32_t dynamicOffsetCount, const uint32_t* pDynamicOffsets) {
	vkCmdBindDescriptorSets(commandBuffer, pipelineBindPoint, layout, firstSet, descriptorSetCount, pDescriptorSets, dynamicOffsetCount, pDynamicOffsets);
}

void fn_vkCmdBindDescriptorSets2(VkCommandBuffer commandBuffer, const VkBindDescriptorSetsInfo* pBindDescriptorSetsInfo) {
	vkCmdBindDescriptorSets2(commandBuffer, pBindDescriptorSetsInfo);
}

void fn_vkCmdBindIndexBuffer(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkIndexType indexType) {
	vkCmdBindIndexBuffer(commandBuffer, buffer, offset, indexType);
}

void fn_vkCmdBindIndexBuffer2(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkDeviceSize size, VkIndexType indexType) {
	vkCmdBindIndexBuffer2(commandBuffer, buffer, offset, size, indexType);
}

void fn_vkCmdBindPipeline(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipeline pipeline) {
	vkCmdBindPipeline(commandBuffer, pipelineBindPoint, pipeline);
}

void fn_vkCmdBindShadersEXT(VkCommandBuffer commandBuffer, uint32_t stageCount, const VkShaderStageFlagBits* pStages, const VkShaderEXT* pShaders) {
	vkCmdBindShadersEXT(commandBuffer, stageCount, pStages, pShaders);
}

void fn_vkCmdBindVertexBuffers(VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, const VkBuffer* pBuffers, const VkDeviceSize* pOffsets) {
	vkCmdBindVertexBuffers(commandBuffer, firstBinding, bindingCount, pBuffers, pOffsets);
}

void fn_vkCmdBindVertexBuffers2(VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, const VkBuffer* pBuffers, const VkDeviceSize* pOffsets, const VkDeviceSize* pSizes, const VkDeviceSize* pStrides) {
	vkCmdBindVertexBuffers2(commandBuffer, firstBinding, bindingCount, pBuffers, pOffsets, pSizes, pStrides);
}

void fn_vkCmdBlitImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkImageBlit* pRegions, VkFilter filter) {
	vkCmdBlitImage(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions, filter);
}

void fn_vkCmdBlitImage2(VkCommandBuffer commandBuffer, const VkBlitImageInfo2* pBlitImageInfo) {
	vkCmdBlitImage2(commandBuffer, pBlitImageInfo);
}

void fn_vkCmdClearAttachments(VkCommandBuffer commandBuffer, uint32_t attachmentCount, const VkClearAttachment* pAttachments, uint32_t rectCount, const VkClearRect* pRects) {
	vkCmdClearAttachments(commandBuffer, attachmentCount, pAttachments, rectCount, pRects);
}

void fn_vkCmdClearColorImage(VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, const VkClearColorValue* pColor, uint32_t rangeCount, const VkImageSubresourceRange* pRanges) {
	vkCmdClearColorImage(commandBuffer, image, imageLayout, pColor, rangeCount, pRanges);
}

void fn_vkCmdClearDepthStencilImage(VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, const VkClearDepthStencilValue* pDepthStencil, uint32_t rangeCount, const VkImageSubresourceRange* pRanges) {
	vkCmdClearDepthStencilImage(commandBuffer, image, imageLayout, pDepthStencil, rangeCount, pRanges);
}

void fn_vkCmdCopyBuffer(VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkBuffer dstBuffer, uint32_t regionCount, const VkBufferCopy* pRegions) {
	vkCmdCopyBuffer(commandBuffer, srcBuffer, dstBuffer, regionCount, pRegions);
}

void fn_vkCmdCopyBuffer2(VkCommandBuffer commandBuffer, const VkCopyBufferInfo2* pCopyBufferInfo) {
	vkCmdCopyBuffer2(commandBuffer, pCopyBufferInfo);
}

void fn_vkCmdCopyBufferToImage(VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkBufferImageCopy* pRegions) {
	vkCmdCopyBufferToImage(commandBuffer, srcBuffer, dstImage, dstImageLayout, regionCount, pRegions);
}

void fn_vkCmdCopyBufferToImage2(VkCommandBuffer commandBuffer, const VkCopyBufferToImageInfo2* pCopyBufferToImageInfo) {
	vkCmdCopyBufferToImage2(commandBuffer, pCopyBufferToImageInfo);
}

void fn_vkCmdCopyImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkImageCopy* pRegions) {
	vkCmdCopyImage(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions);
}

void fn_vkCmdCopyImage2(VkCommandBuffer commandBuffer, const VkCopyImageInfo2* pCopyImageInfo) {
	vkCmdCopyImage2(commandBuffer, pCopyImageInfo);
}

void fn_vkCmdCopyImageToBuffer(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkBuffer dstBuffer, uint32_t regionCount, const VkBufferImageCopy* pRegions) {
	vkCmdCopyImageToBuffer(commandBuffer, srcImage, srcImageLayout, dstBuffer, regionCount, pRegions);
}

void fn_vkCmdCopyImageToBuffer2(VkCommandBuffer commandBuffer, const VkCopyImageToBufferInfo2* pCopyImageToBufferInfo) {
	vkCmdCopyImageToBuffer2(commandBuffer, pCopyImageToBufferInfo);
}

void fn_vkCmdCopyQueryPoolResults(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize stride, VkQueryResultFlags flags) {
	vkCmdCopyQueryPoolResults(commandBuffer, queryPool, firstQuery, queryCount, dstBuffer, dstOffset, stride, flags);
}

void fn_vkCmdDispatch(VkCommandBuffer commandBuffer, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	vkCmdDispatch(commandBuffer, groupCountX, groupCountY, groupCountZ);
}

void fn_vkCmdDispatchBase(VkCommandBuffer commandBuffer, uint32_t baseGroupX, uint32_t baseGroupY, uint32_t baseGroupZ, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	vkCmdDispatchBase(commandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ);
}

void fn_vkCmdDispatchIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset) {
	vkCmdDispatchIndirect(commandBuffer, buffer, offset);
}

void fn_vkCmdDraw(VkCommandBuffer commandBuffer, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance) {
	vkCmdDraw(commandBuffer, vertexCount, instanceCount, firstVertex, firstInstance);
}

void fn_vkCmdDrawIndexed(VkCommandBuffer commandBuffer, uint32_t indexCount, uint32_t instanceCount, uint32_t firstIndex, int32_t vertexOffset, uint32_t firstInstance) {
	vkCmdDrawIndexed(commandBuffer, indexCount, instanceCount, firstIndex, vertexOffset, firstInstance);
}

void fn_vkCmdDrawIndexedIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
	vkCmdDrawIndexedIndirect(commandBuffer, buffer, offset, drawCount, stride);
}

void fn_vkCmdDrawIndexedIndirectCount(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	vkCmdDrawIndexedIndirectCount(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}

void fn_vkCmdDrawIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
	vkCmdDrawIndirect(commandBuffer, buffer, offset, drawCount, stride);
}

void fn_vkCmdDrawIndirectCount(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	vkCmdDrawIndirectCount(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}

void fn_vkCmdEndDebugUtilsLabelEXT(VkCommandBuffer commandBuffer) {
	vkCmdEndDebugUtilsLabelEXT(commandBuffer);
}

void fn_vkCmdEndQuery(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query) {
	vkCmdEndQuery(commandBuffer, queryPool, query);
}

void fn_vkCmdEndRenderPass(VkCommandBuffer commandBuffer) {
	vkCmdEndRenderPass(commandBuffer);
}

void fn_vkCmdEndRenderPass2(VkCommandBuffer commandBuffer, const VkSubpassEndInfo* pSubpassEndInfo) {
	vkCmdEndRenderPass2(commandBuffer, pSubpassEndInfo);
}

void fn_vkCmdEndRendering(VkCommandBuffer commandBuffer) {
	vkCmdEndRendering(commandBuffer);
}

void fn_vkCmdExecuteCommands(VkCommandBuffer commandBuffer, uint32_t commandBufferCount, const VkCommandBuffer* pCommandBuffers) {
	vkCmdExecuteCommands(commandBuffer, commandBufferCount, pCommandBuffers);
}

void fn_vkCmdFillBuffer(VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize size, uint32_t data) {
	vkCmdFillBuffer(commandBuffer, dstBuffer, dstOffset, size, data);
}

void fn_vkCmdInsertDebugUtilsLabelEXT(VkCommandBuffer commandBuffer, const VkDebugUtilsLabelEXT* pLabelInfo) {
	vkCmdInsertDebugUtilsLabelEXT(commandBuffer, pLabelInfo);
}

void fn_vkCmdNextSubpass(VkCommandBuffer commandBuffer, VkSubpassContents contents) {
	vkCmdNextSubpass(commandBuffer, contents);
}

void fn_vkCmdNextSubpass2(VkCommandBuffer commandBuffer, const VkSubpassBeginInfo* pSubpassBeginInfo, const VkSubpassEndInfo* pSubpassEndInfo) {
	vkCmdNextSubpass2(commandBuffer, pSubpassBeginInfo, pSubpassEndInfo);
}

void fn_vkCmdPipelineBarrier(VkCommandBuffer commandBuffer, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, VkDependencyFlags dependencyFlags, uint32_t memoryBarrierCount, const VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, const VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, const VkImageMemoryBarrier* pImageMemoryBarriers) {
	vkCmdPipelineBarrier(commandBuffer, srcStageMask, dstStageMask, dependencyFlags, memoryBarrierCount, pMemoryBarriers, bufferMemoryBarrierCount, pBufferMemoryBarriers, imageMemoryBarrierCount, pImageMemoryBarriers);
}

void fn_vkCmdPipelineBarrier2(VkCommandBuffer commandBuffer, const VkDependencyInfo* pDependencyInfo) {
	vkCmdPipelineBarrier2(commandBuffer, pDependencyInfo);
}

void fn_vkCmdPushConstants(VkCommandBuffer commandBuffer, VkPipelineLayout layout, VkShaderStageFlags stageFlags, uint32_t offset, uint32_t size, const void* pValues) {
	vkCmdPushConstants(commandBuffer, layout, stageFlags, offset, size, pValues);
}

void fn_vkCmdPushConstants2(VkCommandBuffer commandBuffer, const VkPushConstantsInfo* pPushConstantsInfo) {
	vkCmdPushConstants2(commandBuffer, pPushConstantsInfo);
}

void fn_vkCmdPushDescriptorSet(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t set, uint32_t descriptorWriteCount, const VkWriteDescriptorSet* pDescriptorWrites) {
	vkCmdPushDescriptorSet(commandBuffer, pipelineBindPoint, layout, set, descriptorWriteCount, pDescriptorWrites);
}

void fn_vkCmdPushDescriptorSet2(VkCommandBuffer commandBuffer, const VkPushDescriptorSetInfo* pPushDescriptorSetInfo) {
	vkCmdPushDescriptorSet2(commandBuffer, pPushDescriptorSetInfo);
}

void fn_vkCmdPushDescriptorSetWithTemplate(VkCommandBuffer commandBuffer, VkDescriptorUpdateTemplate descriptorUpdateTemplate, VkPipelineLayout layout, uint32_t set, const void* pData) {
	vkCmdPushDescriptorSetWithTemplate(commandBuffer, descriptorUpdateTemplate, layout, set, pData);
}

void fn_vkCmdPushDescriptorSetWithTemplate2(VkCommandBuffer commandBuffer, const VkPushDescriptorSetWithTemplateInfo* pPushDescriptorSetWithTemplateInfo) {
	vkCmdPushDescriptorSetWithTemplate2(commandBuffer, pPushDescriptorSetWithTemplateInfo);
}

void fn_vkCmdResetEvent(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask) {
	vkCmdResetEvent(commandBuffer, event, stageMask);
}

void fn_vkCmdResetEvent2(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags2 stageMask) {
	vkCmdResetEvent2(commandBuffer, event, stageMask);
}

void fn_vkCmdResetQueryPool(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount) {
	vkCmdResetQueryPool(commandBuffer, queryPool, firstQuery, queryCount);
}

void fn_vkCmdResolveImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkImageResolve* pRegions) {
	vkCmdResolveImage(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions);
}

void fn_vkCmdResolveImage2(VkCommandBuffer commandBuffer, const VkResolveImageInfo2* pResolveImageInfo) {
	vkCmdResolveImage2(commandBuffer, pResolveImageInfo);
}

void fn_vkCmdSetAlphaToCoverageEnableEXT(VkCommandBuffer commandBuffer, VkBool32 alphaToCoverageEnable) {
	vkCmdSetAlphaToCoverageEnableEXT(commandBuffer, alphaToCoverageEnable);
}

void fn_vkCmdSetAlphaToOneEnableEXT(VkCommandBuffer commandBuffer, VkBool32 alphaToOneEnable) {
	vkCmdSetAlphaToOneEnableEXT(commandBuffer, alphaToOneEnable);
}

void fn_vkCmdSetBlendConstants(VkCommandBuffer commandBuffer, const float* blendConstants) {
	vkCmdSetBlendConstants(commandBuffer, blendConstants);
}

void fn_vkCmdSetColorBlendAdvancedEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, const VkColorBlendAdvancedEXT* pColorBlendAdvanced) {
	vkCmdSetColorBlendAdvancedEXT(commandBuffer, firstAttachment, attachmentCount, pColorBlendAdvanced);
}

void fn_vkCmdSetColorBlendEnableEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, const VkBool32* pColorBlendEnables) {
	vkCmdSetColorBlendEnableEXT(commandBuffer, firstAttachment, attachmentCount, pColorBlendEnables);
}

void fn_vkCmdSetColorBlendEquationEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, const VkColorBlendEquationEXT* pColorBlendEquations) {
	vkCmdSetColorBlendEquationEXT(commandBuffer, firstAttachment, attachmentCount, pColorBlendEquations);
}

void fn_vkCmdSetColorWriteMaskEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, const VkColorComponentFlags* pColorWriteMasks) {
	vkCmdSetColorWriteMaskEXT(commandBuffer, firstAttachment, attachmentCount, pColorWriteMasks);
}

void fn_vkCmdSetConservativeRasterizationModeEXT(VkCommandBuffer commandBuffer, VkConservativeRasterizationModeEXT conservativeRasterizationMode) {
	vkCmdSetConservativeRasterizationModeEXT(commandBuffer, conservativeRasterizationMode);
}

void fn_vkCmdSetCoverageModulationModeNV(VkCommandBuffer commandBuffer, VkCoverageModulationModeNV coverageModulationMode) {
	vkCmdSetCoverageModulationModeNV(commandBuffer, coverageModulationMode);
}

void fn_vkCmdSetCoverageModulationTableEnableNV(VkCommandBuffer commandBuffer, VkBool32 coverageModulationTableEnable) {
	vkCmdSetCoverageModulationTableEnableNV(commandBuffer, coverageModulationTableEnable);
}

void fn_vkCmdSetCoverageModulationTableNV(VkCommandBuffer commandBuffer, uint32_t coverageModulationTableCount, const float* pCoverageModulationTable) {
	vkCmdSetCoverageModulationTableNV(commandBuffer, coverageModulationTableCount, pCoverageModulationTable);
}

void fn_vkCmdSetCoverageReductionModeNV(VkCommandBuffer commandBuffer, VkCoverageReductionModeNV coverageReductionMode) {
	vkCmdSetCoverageReductionModeNV(commandBuffer, coverageReductionMode);
}

void fn_vkCmdSetCoverageToColorEnableNV(VkCommandBuffer commandBuffer, VkBool32 coverageToColorEnable) {
	vkCmdSetCoverageToColorEnableNV(commandBuffer, coverageToColorEnable);
}

void fn_vkCmdSetCoverageToColorLocationNV(VkCommandBuffer commandBuffer, uint32_t coverageToColorLocation) {
	vkCmdSetCoverageToColorLocationNV(commandBuffer, coverageToColorLocation);
}

void fn_vkCmdSetCullMode(VkCommandBuffer commandBuffer, VkCullModeFlags cullMode) {
	vkCmdSetCullMode(commandBuffer, cullMode);
}

void fn_vkCmdSetDepthBias(VkCommandBuffer commandBuffer, float depthBiasConstantFactor, float depthBiasClamp, float depthBiasSlopeFactor) {
	vkCmdSetDepthBias(commandBuffer, depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor);
}

void fn_vkCmdSetDepthBiasEnable(VkCommandBuffer commandBuffer, VkBool32 depthBiasEnable) {
	vkCmdSetDepthBiasEnable(commandBuffer, depthBiasEnable);
}

void fn_vkCmdSetDepthBounds(VkCommandBuffer commandBuffer, float minDepthBounds, float maxDepthBounds) {
	vkCmdSetDepthBounds(commandBuffer, minDepthBounds, maxDepthBounds);
}

void fn_vkCmdSetDepthBoundsTestEnable(VkCommandBuffer commandBuffer, VkBool32 depthBoundsTestEnable) {
	vkCmdSetDepthBoundsTestEnable(commandBuffer, depthBoundsTestEnable);
}

void fn_vkCmdSetDepthClampEnableEXT(VkCommandBuffer commandBuffer, VkBool32 depthClampEnable) {
	vkCmdSetDepthClampEnableEXT(commandBuffer, depthClampEnable);
}

void fn_vkCmdSetDepthClampRangeEXT(VkCommandBuffer commandBuffer, VkDepthClampModeEXT depthClampMode, const VkDepthClampRangeEXT* pDepthClampRange) {
	vkCmdSetDepthClampRangeEXT(commandBuffer, depthClampMode, pDepthClampRange);
}

void fn_vkCmdSetDepthClipEnableEXT(VkCommandBuffer commandBuffer, VkBool32 depthClipEnable) {
	vkCmdSetDepthClipEnableEXT(commandBuffer, depthClipEnable);
}

void fn_vkCmdSetDepthClipNegativeOneToOneEXT(VkCommandBuffer commandBuffer, VkBool32 negativeOneToOne) {
	vkCmdSetDepthClipNegativeOneToOneEXT(commandBuffer, negativeOneToOne);
}

void fn_vkCmdSetDepthCompareOp(VkCommandBuffer commandBuffer, VkCompareOp depthCompareOp) {
	vkCmdSetDepthCompareOp(commandBuffer, depthCompareOp);
}

void fn_vkCmdSetDepthTestEnable(VkCommandBuffer commandBuffer, VkBool32 depthTestEnable) {
	vkCmdSetDepthTestEnable(commandBuffer, depthTestEnable);
}

void fn_vkCmdSetDepthWriteEnable(VkCommandBuffer commandBuffer, VkBool32 depthWriteEnable) {
	vkCmdSetDepthWriteEnable(commandBuffer, depthWriteEnable);
}

void fn_vkCmdSetDeviceMask(VkCommandBuffer commandBuffer, uint32_t deviceMask) {
	vkCmdSetDeviceMask(commandBuffer, deviceMask);
}

void fn_vkCmdSetEvent(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask) {
	vkCmdSetEvent(commandBuffer, event, stageMask);
}

void fn_vkCmdSetEvent2(VkCommandBuffer commandBuffer, VkEvent event, const VkDependencyInfo* pDependencyInfo) {
	vkCmdSetEvent2(commandBuffer, event, pDependencyInfo);
}

void fn_vkCmdSetExtraPrimitiveOverestimationSizeEXT(VkCommandBuffer commandBuffer, float extraPrimitiveOverestimationSize) {
	vkCmdSetExtraPrimitiveOverestimationSizeEXT(commandBuffer, extraPrimitiveOverestimationSize);
}

void fn_vkCmdSetFrontFace(VkCommandBuffer commandBuffer, VkFrontFace frontFace) {
	vkCmdSetFrontFace(commandBuffer, frontFace);
}

void fn_vkCmdSetLineRasterizationModeEXT(VkCommandBuffer commandBuffer, VkLineRasterizationModeEXT lineRasterizationMode) {
	vkCmdSetLineRasterizationModeEXT(commandBuffer, lineRasterizationMode);
}

void fn_vkCmdSetLineStipple(VkCommandBuffer commandBuffer, uint32_t lineStippleFactor, uint16_t lineStipplePattern) {
	vkCmdSetLineStipple(commandBuffer, lineStippleFactor, lineStipplePattern);
}

void fn_vkCmdSetLineStippleEnableEXT(VkCommandBuffer commandBuffer, VkBool32 stippledLineEnable) {
	vkCmdSetLineStippleEnableEXT(commandBuffer, stippledLineEnable);
}

void fn_vkCmdSetLineWidth(VkCommandBuffer commandBuffer, float lineWidth) {
	vkCmdSetLineWidth(commandBuffer, lineWidth);
}

void fn_vkCmdSetLogicOpEXT(VkCommandBuffer commandBuffer, VkLogicOp logicOp) {
	vkCmdSetLogicOpEXT(commandBuffer, logicOp);
}

void fn_vkCmdSetLogicOpEnableEXT(VkCommandBuffer commandBuffer, VkBool32 logicOpEnable) {
	vkCmdSetLogicOpEnableEXT(commandBuffer, logicOpEnable);
}

void fn_vkCmdSetPatchControlPointsEXT(VkCommandBuffer commandBuffer, uint32_t patchControlPoints) {
	vkCmdSetPatchControlPointsEXT(commandBuffer, patchControlPoints);
}

void fn_vkCmdSetPolygonModeEXT(VkCommandBuffer commandBuffer, VkPolygonMode polygonMode) {
	vkCmdSetPolygonModeEXT(commandBuffer, polygonMode);
}

void fn_vkCmdSetPrimitiveRestartEnable(VkCommandBuffer commandBuffer, VkBool32 primitiveRestartEnable) {
	vkCmdSetPrimitiveRestartEnable(commandBuffer, primitiveRestartEnable);
}

void fn_vkCmdSetPrimitiveTopology(VkCommandBuffer commandBuffer, VkPrimitiveTopology primitiveTopology) {
	vkCmdSetPrimitiveTopology(commandBuffer, primitiveTopology);
}

void fn_vkCmdSetProvokingVertexModeEXT(VkCommandBuffer commandBuffer, VkProvokingVertexModeEXT provokingVertexMode) {
	vkCmdSetProvokingVertexModeEXT(commandBuffer, provokingVertexMode);
}

void fn_vkCmdSetRasterizationSamplesEXT(VkCommandBuffer commandBuffer, VkSampleCountFlagBits rasterizationSamples) {
	vkCmdSetRasterizationSamplesEXT(commandBuffer, rasterizationSamples);
}

void fn_vkCmdSetRasterizationStreamEXT(VkCommandBuffer commandBuffer, uint32_t rasterizationStream) {
	vkCmdSetRasterizationStreamEXT(commandBuffer, rasterizationStream);
}

void fn_vkCmdSetRasterizerDiscardEnable(VkCommandBuffer commandBuffer, VkBool32 rasterizerDiscardEnable) {
	vkCmdSetRasterizerDiscardEnable(commandBuffer, rasterizerDiscardEnable);
}

void fn_vkCmdSetRenderingAttachmentLocations(VkCommandBuffer commandBuffer, const VkRenderingAttachmentLocationInfo* pLocationInfo) {
	vkCmdSetRenderingAttachmentLocations(commandBuffer, pLocationInfo);
}

void fn_vkCmdSetRenderingInputAttachmentIndices(VkCommandBuffer commandBuffer, const VkRenderingInputAttachmentIndexInfo* pInputAttachmentIndexInfo) {
	vkCmdSetRenderingInputAttachmentIndices(commandBuffer, pInputAttachmentIndexInfo);
}

void fn_vkCmdSetRepresentativeFragmentTestEnableNV(VkCommandBuffer commandBuffer, VkBool32 representativeFragmentTestEnable) {
	vkCmdSetRepresentativeFragmentTestEnableNV(commandBuffer, representativeFragmentTestEnable);
}

void fn_vkCmdSetSampleLocationsEnableEXT(VkCommandBuffer commandBuffer, VkBool32 sampleLocationsEnable) {
	vkCmdSetSampleLocationsEnableEXT(commandBuffer, sampleLocationsEnable);
}

void fn_vkCmdSetSampleMaskEXT(VkCommandBuffer commandBuffer, VkSampleCountFlagBits samples, const VkSampleMask* pSampleMask) {
	vkCmdSetSampleMaskEXT(commandBuffer, samples, pSampleMask);
}

void fn_vkCmdSetScissor(VkCommandBuffer commandBuffer, uint32_t firstScissor, uint32_t scissorCount, const VkRect2D* pScissors) {
	vkCmdSetScissor(commandBuffer, firstScissor, scissorCount, pScissors);
}

void fn_vkCmdSetScissorWithCount(VkCommandBuffer commandBuffer, uint32_t scissorCount, const VkRect2D* pScissors) {
	vkCmdSetScissorWithCount(commandBuffer, scissorCount, pScissors);
}

void fn_vkCmdSetShadingRateImageEnableNV(VkCommandBuffer commandBuffer, VkBool32 shadingRateImageEnable) {
	vkCmdSetShadingRateImageEnableNV(commandBuffer, shadingRateImageEnable);
}

void fn_vkCmdSetStencilCompareMask(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t compareMask) {
	vkCmdSetStencilCompareMask(commandBuffer, faceMask, compareMask);
}

void fn_vkCmdSetStencilOp(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, VkStencilOp failOp, VkStencilOp passOp, VkStencilOp depthFailOp, VkCompareOp compareOp) {
	vkCmdSetStencilOp(commandBuffer, faceMask, failOp, passOp, depthFailOp, compareOp);
}

void fn_vkCmdSetStencilReference(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t reference) {
	vkCmdSetStencilReference(commandBuffer, faceMask, reference);
}

void fn_vkCmdSetStencilTestEnable(VkCommandBuffer commandBuffer, VkBool32 stencilTestEnable) {
	vkCmdSetStencilTestEnable(commandBuffer, stencilTestEnable);
}

void fn_vkCmdSetStencilWriteMask(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t writeMask) {
	vkCmdSetStencilWriteMask(commandBuffer, faceMask, writeMask);
}

void fn_vkCmdSetTessellationDomainOriginEXT(VkCommandBuffer commandBuffer, VkTessellationDomainOrigin domainOrigin) {
	vkCmdSetTessellationDomainOriginEXT(commandBuffer, domainOrigin);
}

void fn_vkCmdSetVertexInputEXT(VkCommandBuffer commandBuffer, uint32_t vertexBindingDescriptionCount, const VkVertexInputBindingDescription2EXT* pVertexBindingDescriptions, uint32_t vertexAttributeDescriptionCount, const VkVertexInputAttributeDescription2EXT* pVertexAttributeDescriptions) {
	vkCmdSetVertexInputEXT(commandBuffer, vertexBindingDescriptionCount, pVertexBindingDescriptions, vertexAttributeDescriptionCount, pVertexAttributeDescriptions);
}

void fn_vkCmdSetViewport(VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, const VkViewport* pViewports) {
	vkCmdSetViewport(commandBuffer, firstViewport, viewportCount, pViewports);
}

void fn_vkCmdSetViewportSwizzleNV(VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, const VkViewportSwizzleNV* pViewportSwizzles) {
	vkCmdSetViewportSwizzleNV(commandBuffer, firstViewport, viewportCount, pViewportSwizzles);
}

void fn_vkCmdSetViewportWScalingEnableNV(VkCommandBuffer commandBuffer, VkBool32 viewportWScalingEnable) {
	vkCmdSetViewportWScalingEnableNV(commandBuffer, viewportWScalingEnable);
}

void fn_vkCmdSetViewportWithCount(VkCommandBuffer commandBuffer, uint32_t viewportCount, const VkViewport* pViewports) {
	vkCmdSetViewportWithCount(commandBuffer, viewportCount, pViewports);
}

void fn_vkCmdUpdateBuffer(VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize dataSize, const void* pData) {
	vkCmdUpdateBuffer(commandBuffer, dstBuffer, dstOffset, dataSize, pData);
}

void fn_vkCmdWaitEvents(VkCommandBuffer commandBuffer, uint32_t eventCount, const VkEvent* pEvents, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, uint32_t memoryBarrierCount, const VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, const VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, const VkImageMemoryBarrier* pImageMemoryBarriers) {
	vkCmdWaitEvents(commandBuffer, eventCount, pEvents, srcStageMask, dstStageMask, memoryBarrierCount, pMemoryBarriers, bufferMemoryBarrierCount, pBufferMemoryBarriers, imageMemoryBarrierCount, pImageMemoryBarriers);
}

void fn_vkCmdWaitEvents2(VkCommandBuffer commandBuffer, uint32_t eventCount, const VkEvent* pEvents, const VkDependencyInfo* pDependencyInfos) {
	vkCmdWaitEvents2(commandBuffer, eventCount, pEvents, pDependencyInfos);
}

void fn_vkCmdWriteTimestamp(VkCommandBuffer commandBuffer, VkPipelineStageFlagBits pipelineStage, VkQueryPool queryPool, uint32_t query) {
	vkCmdWriteTimestamp(commandBuffer, pipelineStage, queryPool, query);
}

void fn_vkCmdWriteTimestamp2(VkCommandBuffer commandBuffer, VkPipelineStageFlags2 stage, VkQueryPool queryPool, uint32_t query) {
	vkCmdWriteTimestamp2(commandBuffer, stage, queryPool, query);
}

VkResult fn_vkCopyImageToImage(VkDevice device, const VkCopyImageToImageInfo* pCopyImageToImageInfo) {
	return vkCopyImageToImage(device, pCopyImageToImageInfo);
}

VkResult fn_vkCopyImageToMemory(VkDevice device, const VkCopyImageToMemoryInfo* pCopyImageToMemoryInfo) {
	return vkCopyImageToMemory(device, pCopyImageToMemoryInfo);
}

VkResult fn_vkCopyMemoryToImage(VkDevice device, const VkCopyMemoryToImageInfo* pCopyMemoryToImageInfo) {
	return vkCopyMemoryToImage(device, pCopyMemoryToImageInfo);
}

VkResult fn_vkCreateBuffer(VkDevice device, const VkBufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkBuffer* pBuffer) {
	return vkCreateBuffer(device, pCreateInfo, pAllocator, pBuffer);
}

VkResult fn_vkCreateBufferView(VkDevice device, const VkBufferViewCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkBufferView* pView) {
	return vkCreateBufferView(device, pCreateInfo, pAllocator, pView);
}

VkResult fn_vkCreateCommandPool(VkDevice device, const VkCommandPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkCommandPool* pCommandPool) {
	return vkCreateCommandPool(device, pCreateInfo, pAllocator, pCommandPool);
}

VkResult fn_vkCreateComputePipelines(VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkComputePipelineCreateInfo* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return vkCreateComputePipelines(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}

VkResult fn_vkCreateDebugUtilsMessengerEXT(VkInstance instance, const VkDebugUtilsMessengerCreateInfoEXT* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDebugUtilsMessengerEXT* pMessenger) {
	return vkCreateDebugUtilsMessengerEXT(instance, pCreateInfo, pAllocator, pMessenger);
}

VkResult fn_vkCreateDescriptorPool(VkDevice device, const VkDescriptorPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDescriptorPool* pDescriptorPool) {
	return vkCreateDescriptorPool(device, pCreateInfo, pAllocator, pDescriptorPool);
}

VkResult fn_vkCreateDescriptorSetLayout(VkDevice device, const VkDescriptorSetLayoutCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDescriptorSetLayout* pSetLayout) {
	return vkCreateDescriptorSetLayout(device, pCreateInfo, pAllocator, pSetLayout);
}

VkResult fn_vkCreateDescriptorUpdateTemplate(VkDevice device, const VkDescriptorUpdateTemplateCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDescriptorUpdateTemplate* pDescriptorUpdateTemplate) {
	return vkCreateDescriptorUpdateTemplate(device, pCreateInfo, pAllocator, pDescriptorUpdateTemplate);
}

VkResult fn_vkCreateDevice(VkPhysicalDevice physicalDevice, const VkDeviceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDevice* pDevice) {
	return vkCreateDevice(physicalDevice, pCreateInfo, pAllocator, pDevice);
}

VkResult fn_vkCreateEvent(VkDevice device, const VkEventCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkEvent* pEvent) {
	return vkCreateEvent(device, pCreateInfo, pAllocator, pEvent);
}

VkResult fn_vkCreateFence(VkDevice device, const VkFenceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFence* pFence) {
	return vkCreateFence(device, pCreateInfo, pAllocator, pFence);
}

VkResult fn_vkCreateFramebuffer(VkDevice device, const VkFramebufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFramebuffer* pFramebuffer) {
	return vkCreateFramebuffer(device, pCreateInfo, pAllocator, pFramebuffer);
}

VkResult fn_vkCreateGraphicsPipelines(VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkGraphicsPipelineCreateInfo* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return vkCreateGraphicsPipelines(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}

VkResult fn_vkCreateImage(VkDevice device, const VkImageCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkImage* pImage) {
	return vkCreateImage(device, pCreateInfo, pAllocator, pImage);
}

VkResult fn_vkCreateImageView(VkDevice device, const VkImageViewCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkImageView* pView) {
	return vkCreateImageView(device, pCreateInfo, pAllocator, pView);
}

VkResult fn_vkCreateInstance(const VkInstanceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkInstance* pInstance) {
	return vkCreateInstance(pCreateInfo, pAllocator, pInstance);
}

VkResult fn_vkCreatePipelineCache(VkDevice device, const VkPipelineCacheCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkPipelineCache* pPipelineCache) {
	return vkCreatePipelineCache(device, pCreateInfo, pAllocator, pPipelineCache);
}

VkResult fn_vkCreatePipelineLayout(VkDevice device, const VkPipelineLayoutCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkPipelineLayout* pPipelineLayout) {
	return vkCreatePipelineLayout(device, pCreateInfo, pAllocator, pPipelineLayout);
}

VkResult fn_vkCreatePrivateDataSlot(VkDevice device, const VkPrivateDataSlotCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkPrivateDataSlot* pPrivateDataSlot) {
	return vkCreatePrivateDataSlot(device, pCreateInfo, pAllocator, pPrivateDataSlot);
}

VkResult fn_vkCreateQueryPool(VkDevice device, const VkQueryPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkQueryPool* pQueryPool) {
	return vkCreateQueryPool(device, pCreateInfo, pAllocator, pQueryPool);
}

VkResult fn_vkCreateRenderPass(VkDevice device, const VkRenderPassCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass) {
	return vkCreateRenderPass(device, pCreateInfo, pAllocator, pRenderPass);
}

VkResult fn_vkCreateRenderPass2(VkDevice device, const VkRenderPassCreateInfo2* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass) {
	return vkCreateRenderPass2(device, pCreateInfo, pAllocator, pRenderPass);
}

VkResult fn_vkCreateSampler(VkDevice device, const VkSamplerCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSampler* pSampler) {
	return vkCreateSampler(device, pCreateInfo, pAllocator, pSampler);
}

VkResult fn_vkCreateSamplerYcbcrConversion(VkDevice device, const VkSamplerYcbcrConversionCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSamplerYcbcrConversion* pYcbcrConversion) {
	return vkCreateSamplerYcbcrConversion(device, pCreateInfo, pAllocator, pYcbcrConversion);
}

VkResult fn_vkCreateSemaphore(VkDevice device, const VkSemaphoreCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSemaphore* pSemaphore) {
	return vkCreateSemaphore(device, pCreateInfo, pAllocator, pSemaphore);
}

VkResult fn_vkCreateShaderModule(VkDevice device, const VkShaderModuleCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkShaderModule* pShaderModule) {
	return vkCreateShaderModule(device, pCreateInfo, pAllocator, pShaderModule);
}

VkResult fn_vkCreateShadersEXT(VkDevice device, uint32_t createInfoCount, const VkShaderCreateInfoEXT* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkShaderEXT* pShaders) {
	return vkCreateShadersEXT(device, createInfoCount, pCreateInfos, pAllocator, pShaders);
}

VkResult fn_vkCreateSwapchainKHR(VkDevice device, const VkSwapchainCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSwapchainKHR* pSwapchain) {
	return vkCreateSwapchainKHR(device, pCreateInfo, pAllocator, pSwapchain);
}

void fn_vkDestroyBuffer(VkDevice device, VkBuffer buffer, const VkAllocationCallbacks* pAllocator) {
	vkDestroyBuffer(device, buffer, pAllocator);
}

void fn_vkDestroyBufferView(VkDevice device, VkBufferView bufferView, const VkAllocationCallbacks* pAllocator) {
	vkDestroyBufferView(device, bufferView, pAllocator);
}

void fn_vkDestroyCommandPool(VkDevice device, VkCommandPool commandPool, const VkAllocationCallbacks* pAllocator) {
	vkDestroyCommandPool(device, commandPool, pAllocator);
}

void fn_vkDestroyDebugUtilsMessengerEXT(VkInstance instance, VkDebugUtilsMessengerEXT messenger, const VkAllocationCallbacks* pAllocator) {
	vkDestroyDebugUtilsMessengerEXT(instance, messenger, pAllocator);
}

void fn_vkDestroyDescriptorPool(VkDevice device, VkDescriptorPool descriptorPool, const VkAllocationCallbacks* pAllocator) {
	vkDestroyDescriptorPool(device, descriptorPool, pAllocator);
}

void fn_vkDestroyDescriptorSetLayout(VkDevice device, VkDescriptorSetLayout descriptorSetLayout, const VkAllocationCallbacks* pAllocator) {
	vkDestroyDescriptorSetLayout(device, descriptorSetLayout, pAllocator);
}

void fn_vkDestroyDescriptorUpdateTemplate(VkDevice device, VkDescriptorUpdateTemplate descriptorUpdateTemplate, const VkAllocationCallbacks* pAllocator) {
	vkDestroyDescriptorUpdateTemplate(device, descriptorUpdateTemplate, pAllocator);
}

void fn_vkDestroyDevice(VkDevice device, const VkAllocationCallbacks* pAllocator) {
	vkDestroyDevice(device, pAllocator);
}

void fn_vkDestroyEvent(VkDevice device, VkEvent event, const VkAllocationCallbacks* pAllocator) {
	vkDestroyEvent(device, event, pAllocator);
}

void fn_vkDestroyFence(VkDevice device, VkFence fence, const VkAllocationCallbacks* pAllocator) {
	vkDestroyFence(device, fence, pAllocator);
}

void fn_vkDestroyFramebuffer(VkDevice device, VkFramebuffer framebuffer, const VkAllocationCallbacks* pAllocator) {
	vkDestroyFramebuffer(device, framebuffer, pAllocator);
}

void fn_vkDestroyImage(VkDevice device, VkImage image, const VkAllocationCallbacks* pAllocator) {
	vkDestroyImage(device, image, pAllocator);
}

void fn_vkDestroyImageView(VkDevice device, VkImageView imageView, const VkAllocationCallbacks* pAllocator) {
	vkDestroyImageView(device, imageView, pAllocator);
}

void fn_vkDestroyInstance(VkInstance instance, const VkAllocationCallbacks* pAllocator) {
	vkDestroyInstance(instance, pAllocator);
}

void fn_vkDestroyPipeline(VkDevice device, VkPipeline pipeline, const VkAllocationCallbacks* pAllocator) {
	vkDestroyPipeline(device, pipeline, pAllocator);
}

void fn_vkDestroyPipelineCache(VkDevice device, VkPipelineCache pipelineCache, const VkAllocationCallbacks* pAllocator) {
	vkDestroyPipelineCache(device, pipelineCache, pAllocator);
}

void fn_vkDestroyPipelineLayout(VkDevice device, VkPipelineLayout pipelineLayout, const VkAllocationCallbacks* pAllocator) {
	vkDestroyPipelineLayout(device, pipelineLayout, pAllocator);
}

void fn_vkDestroyPrivateDataSlot(VkDevice device, VkPrivateDataSlot privateDataSlot, const VkAllocationCallbacks* pAllocator) {
	vkDestroyPrivateDataSlot(device, privateDataSlot, pAllocator);
}

void fn_vkDestroyQueryPool(VkDevice device, VkQueryPool queryPool, const VkAllocationCallbacks* pAllocator) {
	vkDestroyQueryPool(device, queryPool, pAllocator);
}

void fn_vkDestroyRenderPass(VkDevice device, VkRenderPass renderPass, const VkAllocationCallbacks* pAllocator) {
	vkDestroyRenderPass(device, renderPass, pAllocator);
}

void fn_vkDestroySampler(VkDevice device, VkSampler sampler, const VkAllocationCallbacks* pAllocator) {
	vkDestroySampler(device, sampler, pAllocator);
}

void fn_vkDestroySamplerYcbcrConversion(VkDevice device, VkSamplerYcbcrConversion ycbcrConversion, const VkAllocationCallbacks* pAllocator) {
	vkDestroySamplerYcbcrConversion(device, ycbcrConversion, pAllocator);
}

void fn_vkDestroySemaphore(VkDevice device, VkSemaphore semaphore, const VkAllocationCallbacks* pAllocator) {
	vkDestroySemaphore(device, semaphore, pAllocator);
}

void fn_vkDestroyShaderEXT(VkDevice device, VkShaderEXT shader, const VkAllocationCallbacks* pAllocator) {
	vkDestroyShaderEXT(device, shader, pAllocator);
}

void fn_vkDestroyShaderModule(VkDevice device, VkShaderModule shaderModule, const VkAllocationCallbacks* pAllocator) {
	vkDestroyShaderModule(device, shaderModule, pAllocator);
}

void fn_vkDestroySurfaceKHR(VkInstance instance, VkSurfaceKHR surface, const VkAllocationCallbacks* pAllocator) {
	vkDestroySurfaceKHR(instance, surface, pAllocator);
}

void fn_vkDestroySwapchainKHR(VkDevice device, VkSwapchainKHR swapchain, const VkAllocationCallbacks* pAllocator) {
	vkDestroySwapchainKHR(device, swapchain, pAllocator);
}

VkResult fn_vkDeviceWaitIdle(VkDevice device) {
	return vkDeviceWaitIdle(device);
}

VkResult fn_vkEndCommandBuffer(VkCommandBuffer commandBuffer) {
	return vkEndCommandBuffer(commandBuffer);
}

VkResult fn_vkEnumerateDeviceExtensionProperties(VkPhysicalDevice physicalDevice, const char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties) {
	return vkEnumerateDeviceExtensionProperties(physicalDevice, pLayerName, pPropertyCount, pProperties);
}

VkResult fn_vkEnumerateDeviceLayerProperties(VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkLayerProperties* pProperties) {
	return vkEnumerateDeviceLayerProperties(physicalDevice, pPropertyCount, pProperties);
}

VkResult fn_vkEnumerateInstanceExtensionProperties(const char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties) {
	return vkEnumerateInstanceExtensionProperties(pLayerName, pPropertyCount, pProperties);
}

VkResult fn_vkEnumerateInstanceLayerProperties(uint32_t* pPropertyCount, VkLayerProperties* pProperties) {
	return vkEnumerateInstanceLayerProperties(pPropertyCount, pProperties);
}

VkResult fn_vkEnumerateInstanceVersion(uint32_t* pApiVersion) {
	return vkEnumerateInstanceVersion(pApiVersion);
}

VkResult fn_vkEnumeratePhysicalDeviceGroups(VkInstance instance, uint32_t* pPhysicalDeviceGroupCount, VkPhysicalDeviceGroupProperties* pPhysicalDeviceGroupProperties) {
	return vkEnumeratePhysicalDeviceGroups(instance, pPhysicalDeviceGroupCount, pPhysicalDeviceGroupProperties);
}

VkResult fn_vkEnumeratePhysicalDevices(VkInstance instance, uint32_t* pPhysicalDeviceCount, VkPhysicalDevice* pPhysicalDevices) {
	return vkEnumeratePhysicalDevices(instance, pPhysicalDeviceCount, pPhysicalDevices);
}

VkResult fn_vkFlushMappedMemoryRanges(VkDevice device, uint32_t memoryRangeCount, const VkMappedMemoryRange* pMemoryRanges) {
	return vkFlushMappedMemoryRanges(device, memoryRangeCount, pMemoryRanges);
}

void fn_vkFreeCommandBuffers(VkDevice device, VkCommandPool commandPool, uint32_t commandBufferCount, const VkCommandBuffer* pCommandBuffers) {
	vkFreeCommandBuffers(device, commandPool, commandBufferCount, pCommandBuffers);
}

VkResult fn_vkFreeDescriptorSets(VkDevice device, VkDescriptorPool descriptorPool, uint32_t descriptorSetCount, const VkDescriptorSet* pDescriptorSets) {
	return vkFreeDescriptorSets(device, descriptorPool, descriptorSetCount, pDescriptorSets);
}

void fn_vkFreeMemory(VkDevice device, VkDeviceMemory memory, const VkAllocationCallbacks* pAllocator) {
	vkFreeMemory(device, memory, pAllocator);
}

VkDeviceAddress fn_vkGetBufferDeviceAddress(VkDevice device, const VkBufferDeviceAddressInfo* pInfo) {
	return vkGetBufferDeviceAddress(device, pInfo);
}

void fn_vkGetBufferMemoryRequirements(VkDevice device, VkBuffer buffer, VkMemoryRequirements* pMemoryRequirements) {
	vkGetBufferMemoryRequirements(device, buffer, pMemoryRequirements);
}

void fn_vkGetBufferMemoryRequirements2(VkDevice device, const VkBufferMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	vkGetBufferMemoryRequirements2(device, pInfo, pMemoryRequirements);
}

uint64_t fn_vkGetBufferOpaqueCaptureAddress(VkDevice device, const VkBufferDeviceAddressInfo* pInfo) {
	return vkGetBufferOpaqueCaptureAddress(device, pInfo);
}

void fn_vkGetDescriptorSetLayoutSupport(VkDevice device, const VkDescriptorSetLayoutCreateInfo* pCreateInfo, VkDescriptorSetLayoutSupport* pSupport) {
	vkGetDescriptorSetLayoutSupport(device, pCreateInfo, pSupport);
}

void fn_vkGetDeviceBufferMemoryRequirements(VkDevice device, const VkDeviceBufferMemoryRequirements* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	vkGetDeviceBufferMemoryRequirements(device, pInfo, pMemoryRequirements);
}

void fn_vkGetDeviceGroupPeerMemoryFeatures(VkDevice device, uint32_t heapIndex, uint32_t localDeviceIndex, uint32_t remoteDeviceIndex, VkPeerMemoryFeatureFlags* pPeerMemoryFeatures) {
	vkGetDeviceGroupPeerMemoryFeatures(device, heapIndex, localDeviceIndex, remoteDeviceIndex, pPeerMemoryFeatures);
}

VkResult fn_vkGetDeviceGroupPresentCapabilitiesKHR(VkDevice device, VkDeviceGroupPresentCapabilitiesKHR* pDeviceGroupPresentCapabilities) {
	return vkGetDeviceGroupPresentCapabilitiesKHR(device, pDeviceGroupPresentCapabilities);
}

VkResult fn_vkGetDeviceGroupSurfacePresentModesKHR(VkDevice device, VkSurfaceKHR surface, VkDeviceGroupPresentModeFlagsKHR* pModes) {
	return vkGetDeviceGroupSurfacePresentModesKHR(device, surface, pModes);
}

void fn_vkGetDeviceImageMemoryRequirements(VkDevice device, const VkDeviceImageMemoryRequirements* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	vkGetDeviceImageMemoryRequirements(device, pInfo, pMemoryRequirements);
}

void fn_vkGetDeviceImageSparseMemoryRequirements(VkDevice device, const VkDeviceImageMemoryRequirements* pInfo, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2* pSparseMemoryRequirements) {
	vkGetDeviceImageSparseMemoryRequirements(device, pInfo, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}

void fn_vkGetDeviceImageSubresourceLayout(VkDevice device, const VkDeviceImageSubresourceInfo* pInfo, VkSubresourceLayout2* pLayout) {
	vkGetDeviceImageSubresourceLayout(device, pInfo, pLayout);
}

void fn_vkGetDeviceMemoryCommitment(VkDevice device, VkDeviceMemory memory, VkDeviceSize* pCommittedMemoryInBytes) {
	vkGetDeviceMemoryCommitment(device, memory, pCommittedMemoryInBytes);
}

uint64_t fn_vkGetDeviceMemoryOpaqueCaptureAddress(VkDevice device, const VkDeviceMemoryOpaqueCaptureAddressInfo* pInfo) {
	return vkGetDeviceMemoryOpaqueCaptureAddress(device, pInfo);
}

PFN_vkVoidFunction fn_vkGetDeviceProcAddr(VkDevice device, const char* pName) {
	return vkGetDeviceProcAddr(device, pName);
}

void fn_vkGetDeviceQueue(VkDevice device, uint32_t queueFamilyIndex, uint32_t queueIndex, VkQueue* pQueue) {
	vkGetDeviceQueue(device, queueFamilyIndex, queueIndex, pQueue);
}

void fn_vkGetDeviceQueue2(VkDevice device, const VkDeviceQueueInfo2* pQueueInfo, VkQueue* pQueue) {
	vkGetDeviceQueue2(device, pQueueInfo, pQueue);
}

VkResult fn_vkGetEventStatus(VkDevice device, VkEvent event) {
	return vkGetEventStatus(device, event);
}

VkResult fn_vkGetFenceStatus(VkDevice device, VkFence fence) {
	return vkGetFenceStatus(device, fence);
}

void fn_vkGetImageMemoryRequirements(VkDevice device, VkImage image, VkMemoryRequirements* pMemoryRequirements) {
	vkGetImageMemoryRequirements(device, image, pMemoryRequirements);
}

void fn_vkGetImageMemoryRequirements2(VkDevice device, const VkImageMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	vkGetImageMemoryRequirements2(device, pInfo, pMemoryRequirements);
}

void fn_vkGetImageSparseMemoryRequirements(VkDevice device, VkImage image, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements* pSparseMemoryRequirements) {
	vkGetImageSparseMemoryRequirements(device, image, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}

void fn_vkGetImageSparseMemoryRequirements2(VkDevice device, const VkImageSparseMemoryRequirementsInfo2* pInfo, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2* pSparseMemoryRequirements) {
	vkGetImageSparseMemoryRequirements2(device, pInfo, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}

void fn_vkGetImageSubresourceLayout(VkDevice device, VkImage image, const VkImageSubresource* pSubresource, VkSubresourceLayout* pLayout) {
	vkGetImageSubresourceLayout(device, image, pSubresource, pLayout);
}

void fn_vkGetImageSubresourceLayout2(VkDevice device, VkImage image, const VkImageSubresource2* pSubresource, VkSubresourceLayout2* pLayout) {
	vkGetImageSubresourceLayout2(device, image, pSubresource, pLayout);
}

PFN_vkVoidFunction fn_vkGetInstanceProcAddr(VkInstance instance, const char* pName) {
	return vkGetInstanceProcAddr(instance, pName);
}

void fn_vkGetPhysicalDeviceExternalBufferProperties(VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalBufferInfo* pExternalBufferInfo, VkExternalBufferProperties* pExternalBufferProperties) {
	vkGetPhysicalDeviceExternalBufferProperties(physicalDevice, pExternalBufferInfo, pExternalBufferProperties);
}

void fn_vkGetPhysicalDeviceExternalFenceProperties(VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalFenceInfo* pExternalFenceInfo, VkExternalFenceProperties* pExternalFenceProperties) {
	vkGetPhysicalDeviceExternalFenceProperties(physicalDevice, pExternalFenceInfo, pExternalFenceProperties);
}

void fn_vkGetPhysicalDeviceExternalSemaphoreProperties(VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalSemaphoreInfo* pExternalSemaphoreInfo, VkExternalSemaphoreProperties* pExternalSemaphoreProperties) {
	vkGetPhysicalDeviceExternalSemaphoreProperties(physicalDevice, pExternalSemaphoreInfo, pExternalSemaphoreProperties);
}

void fn_vkGetPhysicalDeviceFeatures(VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures* pFeatures) {
	vkGetPhysicalDeviceFeatures(physicalDevice, pFeatures);
}

void fn_vkGetPhysicalDeviceFeatures2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures2* pFeatures) {
	vkGetPhysicalDeviceFeatures2(physicalDevice, pFeatures);
}

void fn_vkGetPhysicalDeviceFormatProperties(VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties* pFormatProperties) {
	vkGetPhysicalDeviceFormatProperties(physicalDevice, format, pFormatProperties);
}

void fn_vkGetPhysicalDeviceFormatProperties2(VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties2* pFormatProperties) {
	vkGetPhysicalDeviceFormatProperties2(physicalDevice, format, pFormatProperties);
}

VkResult fn_vkGetPhysicalDeviceImageFormatProperties(VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkImageTiling tiling, VkImageUsageFlags usage, VkImageCreateFlags flags, VkImageFormatProperties* pImageFormatProperties) {
	return vkGetPhysicalDeviceImageFormatProperties(physicalDevice, format, type, tiling, usage, flags, pImageFormatProperties);
}

VkResult fn_vkGetPhysicalDeviceImageFormatProperties2(VkPhysicalDevice physicalDevice, const VkPhysicalDeviceImageFormatInfo2* pImageFormatInfo, VkImageFormatProperties2* pImageFormatProperties) {
	return vkGetPhysicalDeviceImageFormatProperties2(physicalDevice, pImageFormatInfo, pImageFormatProperties);
}

void fn_vkGetPhysicalDeviceMemoryProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties* pMemoryProperties) {
	vkGetPhysicalDeviceMemoryProperties(physicalDevice, pMemoryProperties);
}

void fn_vkGetPhysicalDeviceMemoryProperties2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties2* pMemoryProperties) {
	vkGetPhysicalDeviceMemoryProperties2(physicalDevice, pMemoryProperties);
}

VkResult fn_vkGetPhysicalDevicePresentRectanglesKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pRectCount, VkRect2D* pRects) {
	return vkGetPhysicalDevicePresentRectanglesKHR(physicalDevice, surface, pRectCount, pRects);
}

void fn_vkGetPhysicalDeviceProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties* pProperties) {
	vkGetPhysicalDeviceProperties(physicalDevice, pProperties);
}

void fn_vkGetPhysicalDeviceProperties2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties2* pProperties) {
	vkGetPhysicalDeviceProperties2(physicalDevice, pProperties);
}

void fn_vkGetPhysicalDeviceQueueFamilyProperties(VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties* pQueueFamilyProperties) {
	vkGetPhysicalDeviceQueueFamilyProperties(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}

void fn_vkGetPhysicalDeviceQueueFamilyProperties2(VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties2* pQueueFamilyProperties) {
	vkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}

void fn_vkGetPhysicalDeviceSparseImageFormatProperties(VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkSampleCountFlagBits samples, VkImageUsageFlags usage, VkImageTiling tiling, uint32_t* pPropertyCount, VkSparseImageFormatProperties* pProperties) {
	vkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice, format, type, samples, usage, tiling, pPropertyCount, pProperties);
}

void fn_vkGetPhysicalDeviceSparseImageFormatProperties2(VkPhysicalDevice physicalDevice, const VkPhysicalDeviceSparseImageFormatInfo2* pFormatInfo, uint32_t* pPropertyCount, VkSparseImageFormatProperties2* pProperties) {
	vkGetPhysicalDeviceSparseImageFormatProperties2(physicalDevice, pFormatInfo, pPropertyCount, pProperties);
}

VkResult fn_vkGetPhysicalDeviceSurfaceCapabilitiesKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, VkSurfaceCapabilitiesKHR* pSurfaceCapabilities) {
	return vkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice, surface, pSurfaceCapabilities);
}

VkResult fn_vkGetPhysicalDeviceSurfaceFormatsKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pSurfaceFormatCount, VkSurfaceFormatKHR* pSurfaceFormats) {
	return vkGetPhysicalDeviceSurfaceFormatsKHR(physicalDevice, surface, pSurfaceFormatCount, pSurfaceFormats);
}

VkResult fn_vkGetPhysicalDeviceSurfacePresentModesKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pPresentModeCount, VkPresentModeKHR* pPresentModes) {
	return vkGetPhysicalDeviceSurfacePresentModesKHR(physicalDevice, surface, pPresentModeCount, pPresentModes);
}

VkResult fn_vkGetPhysicalDeviceSurfaceSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, VkSurfaceKHR surface, VkBool32* pSupported) {
	return vkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice, queueFamilyIndex, surface, pSupported);
}

VkResult fn_vkGetPhysicalDeviceToolProperties(VkPhysicalDevice physicalDevice, uint32_t* pToolCount, VkPhysicalDeviceToolProperties* pToolProperties) {
	return vkGetPhysicalDeviceToolProperties(physicalDevice, pToolCount, pToolProperties);
}

VkResult fn_vkGetPipelineCacheData(VkDevice device, VkPipelineCache pipelineCache, size_t* pDataSize, void* pData) {
	return vkGetPipelineCacheData(device, pipelineCache, pDataSize, pData);
}

void fn_vkGetPrivateData(VkDevice device, VkObjectType objectType, uint64_t objectHandle, VkPrivateDataSlot privateDataSlot, uint64_t* pData) {
	vkGetPrivateData(device, objectType, objectHandle, privateDataSlot, pData);
}

VkResult fn_vkGetQueryPoolResults(VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, size_t dataSize, void* pData, VkDeviceSize stride, VkQueryResultFlags flags) {
	return vkGetQueryPoolResults(device, queryPool, firstQuery, queryCount, dataSize, pData, stride, flags);
}

void fn_vkGetRenderAreaGranularity(VkDevice device, VkRenderPass renderPass, VkExtent2D* pGranularity) {
	vkGetRenderAreaGranularity(device, renderPass, pGranularity);
}

void fn_vkGetRenderingAreaGranularity(VkDevice device, const VkRenderingAreaInfo* pRenderingAreaInfo, VkExtent2D* pGranularity) {
	vkGetRenderingAreaGranularity(device, pRenderingAreaInfo, pGranularity);
}

VkResult fn_vkGetSemaphoreCounterValue(VkDevice device, VkSemaphore semaphore, uint64_t* pValue) {
	return vkGetSemaphoreCounterValue(device, semaphore, pValue);
}

VkResult fn_vkGetShaderBinaryDataEXT(VkDevice device, VkShaderEXT shader, size_t* pDataSize, void* pData) {
	return vkGetShaderBinaryDataEXT(device, shader, pDataSize, pData);
}

VkResult fn_vkGetSwapchainImagesKHR(VkDevice device, VkSwapchainKHR swapchain, uint32_t* pSwapchainImageCount, VkImage* pSwapchainImages) {
	return vkGetSwapchainImagesKHR(device, swapchain, pSwapchainImageCount, pSwapchainImages);
}

VkResult fn_vkInvalidateMappedMemoryRanges(VkDevice device, uint32_t memoryRangeCount, const VkMappedMemoryRange* pMemoryRanges) {
	return vkInvalidateMappedMemoryRanges(device, memoryRangeCount, pMemoryRanges);
}

VkResult fn_vkMapMemory(VkDevice device, VkDeviceMemory memory, VkDeviceSize offset, VkDeviceSize size, VkMemoryMapFlags flags, void** ppData) {
	return vkMapMemory(device, memory, offset, size, flags, ppData);
}

VkResult fn_vkMapMemory2(VkDevice device, const VkMemoryMapInfo* pMemoryMapInfo, void** ppData) {
	return vkMapMemory2(device, pMemoryMapInfo, ppData);
}

VkResult fn_vkMergePipelineCaches(VkDevice device, VkPipelineCache dstCache, uint32_t srcCacheCount, const VkPipelineCache* pSrcCaches) {
	return vkMergePipelineCaches(device, dstCache, srcCacheCount, pSrcCaches);
}

void fn_vkQueueBeginDebugUtilsLabelEXT(VkQueue queue, const VkDebugUtilsLabelEXT* pLabelInfo) {
	vkQueueBeginDebugUtilsLabelEXT(queue, pLabelInfo);
}

VkResult fn_vkQueueBindSparse(VkQueue queue, uint32_t bindInfoCount, const VkBindSparseInfo* pBindInfo, VkFence fence) {
	return vkQueueBindSparse(queue, bindInfoCount, pBindInfo, fence);
}

void fn_vkQueueEndDebugUtilsLabelEXT(VkQueue queue) {
	vkQueueEndDebugUtilsLabelEXT(queue);
}

void fn_vkQueueInsertDebugUtilsLabelEXT(VkQueue queue, const VkDebugUtilsLabelEXT* pLabelInfo) {
	vkQueueInsertDebugUtilsLabelEXT(queue, pLabelInfo);
}

VkResult fn_vkQueuePresentKHR(VkQueue queue, const VkPresentInfoKHR* pPresentInfo) {
	return vkQueuePresentKHR(queue, pPresentInfo);
}

VkResult fn_vkQueueSubmit(VkQueue queue, uint32_t submitCount, const VkSubmitInfo* pSubmits, VkFence fence) {
	return vkQueueSubmit(queue, submitCount, pSubmits, fence);
}

VkResult fn_vkQueueSubmit2(VkQueue queue, uint32_t submitCount, const VkSubmitInfo2* pSubmits, VkFence fence) {
	return vkQueueSubmit2(queue, submitCount, pSubmits, fence);
}

VkResult fn_vkQueueWaitIdle(VkQueue queue) {
	return vkQueueWaitIdle(queue);
}

VkResult fn_vkResetCommandBuffer(VkCommandBuffer commandBuffer, VkCommandBufferResetFlags flags) {
	return vkResetCommandBuffer(commandBuffer, flags);
}

VkResult fn_vkResetCommandPool(VkDevice device, VkCommandPool commandPool, VkCommandPoolResetFlags flags) {
	return vkResetCommandPool(device, commandPool, flags);
}

VkResult fn_vkResetDescriptorPool(VkDevice device, VkDescriptorPool descriptorPool, VkDescriptorPoolResetFlags flags) {
	return vkResetDescriptorPool(device, descriptorPool, flags);
}

VkResult fn_vkResetEvent(VkDevice device, VkEvent event) {
	return vkResetEvent(device, event);
}

VkResult fn_vkResetFences(VkDevice device, uint32_t fenceCount, const VkFence* pFences) {
	return vkResetFences(device, fenceCount, pFences);
}

void fn_vkResetQueryPool(VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount) {
	vkResetQueryPool(device, queryPool, firstQuery, queryCount);
}

VkResult fn_vkSetDebugUtilsObjectNameEXT(VkDevice device, const VkDebugUtilsObjectNameInfoEXT* pNameInfo) {
	return vkSetDebugUtilsObjectNameEXT(device, pNameInfo);
}

VkResult fn_vkSetDebugUtilsObjectTagEXT(VkDevice device, const VkDebugUtilsObjectTagInfoEXT* pTagInfo) {
	return vkSetDebugUtilsObjectTagEXT(device, pTagInfo);
}

VkResult fn_vkSetEvent(VkDevice device, VkEvent event) {
	return vkSetEvent(device, event);
}

VkResult fn_vkSetPrivateData(VkDevice device, VkObjectType objectType, uint64_t objectHandle, VkPrivateDataSlot privateDataSlot, uint64_t data) {
	return vkSetPrivateData(device, objectType, objectHandle, privateDataSlot, data);
}

VkResult fn_vkSignalSemaphore(VkDevice device, const VkSemaphoreSignalInfo* pSignalInfo) {
	return vkSignalSemaphore(device, pSignalInfo);
}

void fn_vkSubmitDebugUtilsMessageEXT(VkInstance instance, VkDebugUtilsMessageSeverityFlagBitsEXT messageSeverity, VkDebugUtilsMessageTypeFlagsEXT messageTypes, const VkDebugUtilsMessengerCallbackDataEXT* pCallbackData) {
	vkSubmitDebugUtilsMessageEXT(instance, messageSeverity, messageTypes, pCallbackData);
}

VkResult fn_vkTransitionImageLayout(VkDevice device, uint32_t transitionCount, const VkHostImageLayoutTransitionInfo* pTransitions) {
	return vkTransitionImageLayout(device, transitionCount, pTransitions);
}

void fn_vkTrimCommandPool(VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlags flags) {
	vkTrimCommandPool(device, commandPool, flags);
}

void fn_vkUnmapMemory(VkDevice device, VkDeviceMemory memory) {
	vkUnmapMemory(device, memory);
}

VkResult fn_vkUnmapMemory2(VkDevice device, const VkMemoryUnmapInfo* pMemoryUnmapInfo) {
	return vkUnmapMemory2(device, pMemoryUnmapInfo);
}

void fn_vkUpdateDescriptorSetWithTemplate(VkDevice device, VkDescriptorSet descriptorSet, VkDescriptorUpdateTemplate descriptorUpdateTemplate, const void* pData) {
	vkUpdateDescriptorSetWithTemplate(device, descriptorSet, descriptorUpdateTemplate, pData);
}

void fn_vkUpdateDescriptorSets(VkDevice device, uint32_t descriptorWriteCount, const VkWriteDescriptorSet* pDescriptorWrites, uint32_t descriptorCopyCount, const VkCopyDescriptorSet* pDescriptorCopies) {
	vkUpdateDescriptorSets(device, descriptorWriteCount, pDescriptorWrites, descriptorCopyCount, pDescriptorCopies);
}

VkResult fn_vkWaitForFences(VkDevice device, uint32_t fenceCount, const VkFence* pFences, VkBool32 waitAll, uint64_t timeout) {
	return vkWaitForFences(device, fenceCount, pFences, waitAll, timeout);
}

VkResult fn_vkWaitSemaphores(VkDevice device, const VkSemaphoreWaitInfo* pWaitInfo, uint64_t timeout) {
	return vkWaitSemaphores(device, pWaitInfo, timeout);
}


#ifdef VK_USE_PLATFORM_WAYLAND_KHR
VkResult fn_vkCreateWaylandSurfaceKHR(VkInstance instance, const VkWaylandSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return vkCreateWaylandSurfaceKHR(instance, pCreateInfo, pAllocator, pSurface);
}

VkBool32 fn_vkGetPhysicalDeviceWaylandPresentationSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, struct wl_display* display) {
	return vkGetPhysicalDeviceWaylandPresentationSupportKHR(physicalDevice, queueFamilyIndex, display);
}

#endif // VK_USE_PLATFORM_WAYLAND_KHR

#ifdef VK_USE_PLATFORM_WIN32_KHR
VkResult fn_vkCreateWin32SurfaceKHR(VkInstance instance, const VkWin32SurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return vkCreateWin32SurfaceKHR(instance, pCreateInfo, pAllocator, pSurface);
}

VkBool32 fn_vkGetPhysicalDeviceWin32PresentationSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex) {
	return vkGetPhysicalDeviceWin32PresentationSupportKHR(physicalDevice, queueFamilyIndex);
}

#endif // VK_USE_PLATFORM_WIN32_KHR

#ifdef VK_USE_PLATFORM_XCB_KHR
VkResult fn_vkCreateXcbSurfaceKHR(VkInstance instance, const VkXcbSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return vkCreateXcbSurfaceKHR(instance, pCreateInfo, pAllocator, pSurface);
}

VkBool32 fn_vkGetPhysicalDeviceXcbPresentationSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, xcb_connection_t* connection, xcb_visualid_t visual_id) {
	return vkGetPhysicalDeviceXcbPresentationSupportKHR(physicalDevice, queueFamilyIndex, connection, visual_id);
}

#endif // VK_USE_PLATFORM_XCB_KHR

#ifdef VK_USE_PLATFORM_XLIB_KHR
VkResult fn_vkCreateXlibSurfaceKHR(VkInstance instance, const VkXlibSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return vkCreateXlibSurfaceKHR(instance, pCreateInfo, pAllocator, pSurface);
}

VkBool32 fn_vkGetPhysicalDeviceXlibPresentationSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, Display* dpy, VisualID visualID) {
	return vkGetPhysicalDeviceXlibPresentationSupportKHR(physicalDevice, queueFamilyIndex, dpy, visualID);
}

#endif // VK_USE_PLATFORM_XLIB_KHR
void* trampoline_AllocationCallbacks_PfnAllocation(void* p0, size_t p1, size_t p2, VkSystemAllocationScope p3) {
	return go_bridge_AllocationCallbacks_PfnAllocation(p0, p1, p2, p3);
}
void* trampoline_AllocationCallbacks_PfnReallocation(void* p0, void* p1, size_t p2, size_t p3, VkSystemAllocationScope p4) {
	return go_bridge_AllocationCallbacks_PfnReallocation(p0, p1, p2, p3, p4);
}
void trampoline_AllocationCallbacks_PfnFree(void* p0, void* p1) {
	go_bridge_AllocationCallbacks_PfnFree(p0, p1);
}
void trampoline_AllocationCallbacks_PfnInternalAllocation(void* p0, size_t p1, VkInternalAllocationType p2, VkSystemAllocationScope p3) {
	go_bridge_AllocationCallbacks_PfnInternalAllocation(p0, p1, p2, p3);
}
void trampoline_AllocationCallbacks_PfnInternalFree(void* p0, size_t p1, VkInternalAllocationType p2, VkSystemAllocationScope p3) {
	go_bridge_AllocationCallbacks_PfnInternalFree(p0, p1, p2, p3);
}
VkBool32 trampoline_DebugUtilsMessengerCreateInfoEXT_PfnUserCallback(VkDebugUtilsMessageSeverityFlagBitsEXT p0, VkDebugUtilsMessageTypeFlagsEXT p1, VkDebugUtilsMessengerCallbackDataEXT* p2, void* p3) {
	return go_bridge_DebugUtilsMessengerCreateInfoEXT_PfnUserCallback(p0, p1, p2, p3);
}
