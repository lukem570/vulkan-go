#define VOLK_IMPLEMENTATION
#include "volk.h"
#include <vulkan/vulkan.h>
#include "volk_wrappers.h"

void fn_vkCmdBeginRenderPass(VkCommandBuffer commandBuffer, const VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassContents contents) {
	vkCmdBeginRenderPass(commandBuffer, pRenderPassBegin, contents);
}

void fn_vkCmdBeginRenderPass2(VkCommandBuffer commandBuffer, const VkRenderPassBeginInfo* pRenderPassBegin, const VkSubpassBeginInfo* pSubpassBeginInfo) {
	vkCmdBeginRenderPass2(commandBuffer, pRenderPassBegin, pSubpassBeginInfo);
}

void fn_vkCmdBeginRendering(VkCommandBuffer commandBuffer, const VkRenderingInfo* pRenderingInfo) {
	vkCmdBeginRendering(commandBuffer, pRenderingInfo);
}

void fn_vkCmdBindIndexBuffer(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkIndexType indexType) {
	vkCmdBindIndexBuffer(commandBuffer, buffer, offset, indexType);
}

void fn_vkCmdBindIndexBuffer2(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkDeviceSize size, VkIndexType indexType) {
	vkCmdBindIndexBuffer2(commandBuffer, buffer, offset, size, indexType);
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

void fn_vkCmdClearDepthStencilImage(VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, const VkClearDepthStencilValue* pDepthStencil, uint32_t rangeCount, const VkImageSubresourceRange* pRanges) {
	vkCmdClearDepthStencilImage(commandBuffer, image, imageLayout, pDepthStencil, rangeCount, pRanges);
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

void fn_vkCmdEndRenderPass(VkCommandBuffer commandBuffer) {
	vkCmdEndRenderPass(commandBuffer);
}

void fn_vkCmdEndRenderPass2(VkCommandBuffer commandBuffer, const VkSubpassEndInfo* pSubpassEndInfo) {
	vkCmdEndRenderPass2(commandBuffer, pSubpassEndInfo);
}

void fn_vkCmdEndRendering(VkCommandBuffer commandBuffer) {
	vkCmdEndRendering(commandBuffer);
}

void fn_vkCmdNextSubpass(VkCommandBuffer commandBuffer, VkSubpassContents contents) {
	vkCmdNextSubpass(commandBuffer, contents);
}

void fn_vkCmdNextSubpass2(VkCommandBuffer commandBuffer, const VkSubpassBeginInfo* pSubpassBeginInfo, const VkSubpassEndInfo* pSubpassEndInfo) {
	vkCmdNextSubpass2(commandBuffer, pSubpassBeginInfo, pSubpassEndInfo);
}

void fn_vkCmdResolveImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkImageResolve* pRegions) {
	vkCmdResolveImage(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions);
}

void fn_vkCmdResolveImage2(VkCommandBuffer commandBuffer, const VkResolveImageInfo2* pResolveImageInfo) {
	vkCmdResolveImage2(commandBuffer, pResolveImageInfo);
}

void fn_vkCmdSetBlendConstants(VkCommandBuffer commandBuffer, const float* blendConstants) {
	vkCmdSetBlendConstants(commandBuffer, blendConstants);
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

void fn_vkCmdSetDepthCompareOp(VkCommandBuffer commandBuffer, VkCompareOp depthCompareOp) {
	vkCmdSetDepthCompareOp(commandBuffer, depthCompareOp);
}

void fn_vkCmdSetDepthTestEnable(VkCommandBuffer commandBuffer, VkBool32 depthTestEnable) {
	vkCmdSetDepthTestEnable(commandBuffer, depthTestEnable);
}

void fn_vkCmdSetDepthWriteEnable(VkCommandBuffer commandBuffer, VkBool32 depthWriteEnable) {
	vkCmdSetDepthWriteEnable(commandBuffer, depthWriteEnable);
}

void fn_vkCmdSetFrontFace(VkCommandBuffer commandBuffer, VkFrontFace frontFace) {
	vkCmdSetFrontFace(commandBuffer, frontFace);
}

void fn_vkCmdSetLineStipple(VkCommandBuffer commandBuffer, uint32_t lineStippleFactor, uint16_t lineStipplePattern) {
	vkCmdSetLineStipple(commandBuffer, lineStippleFactor, lineStipplePattern);
}

void fn_vkCmdSetLineWidth(VkCommandBuffer commandBuffer, float lineWidth) {
	vkCmdSetLineWidth(commandBuffer, lineWidth);
}

void fn_vkCmdSetPrimitiveRestartEnable(VkCommandBuffer commandBuffer, VkBool32 primitiveRestartEnable) {
	vkCmdSetPrimitiveRestartEnable(commandBuffer, primitiveRestartEnable);
}

void fn_vkCmdSetPrimitiveTopology(VkCommandBuffer commandBuffer, VkPrimitiveTopology primitiveTopology) {
	vkCmdSetPrimitiveTopology(commandBuffer, primitiveTopology);
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

void fn_vkCmdSetScissor(VkCommandBuffer commandBuffer, uint32_t firstScissor, uint32_t scissorCount, const VkRect2D* pScissors) {
	vkCmdSetScissor(commandBuffer, firstScissor, scissorCount, pScissors);
}

void fn_vkCmdSetScissorWithCount(VkCommandBuffer commandBuffer, uint32_t scissorCount, const VkRect2D* pScissors) {
	vkCmdSetScissorWithCount(commandBuffer, scissorCount, pScissors);
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

void fn_vkCmdSetViewport(VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, const VkViewport* pViewports) {
	vkCmdSetViewport(commandBuffer, firstViewport, viewportCount, pViewports);
}

void fn_vkCmdSetViewportWithCount(VkCommandBuffer commandBuffer, uint32_t viewportCount, const VkViewport* pViewports) {
	vkCmdSetViewportWithCount(commandBuffer, viewportCount, pViewports);
}

VkResult fn_vkCreateFramebuffer(VkDevice device, const VkFramebufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFramebuffer* pFramebuffer) {
	return vkCreateFramebuffer(device, pCreateInfo, pAllocator, pFramebuffer);
}

VkResult fn_vkCreateGraphicsPipelines(VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkGraphicsPipelineCreateInfo* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return vkCreateGraphicsPipelines(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}

VkResult fn_vkCreateRenderPass(VkDevice device, const VkRenderPassCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass) {
	return vkCreateRenderPass(device, pCreateInfo, pAllocator, pRenderPass);
}

VkResult fn_vkCreateRenderPass2(VkDevice device, const VkRenderPassCreateInfo2* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass) {
	return vkCreateRenderPass2(device, pCreateInfo, pAllocator, pRenderPass);
}

void fn_vkDestroyFramebuffer(VkDevice device, VkFramebuffer framebuffer, const VkAllocationCallbacks* pAllocator) {
	vkDestroyFramebuffer(device, framebuffer, pAllocator);
}

void fn_vkDestroyRenderPass(VkDevice device, VkRenderPass renderPass, const VkAllocationCallbacks* pAllocator) {
	vkDestroyRenderPass(device, renderPass, pAllocator);
}

void fn_vkGetRenderAreaGranularity(VkDevice device, VkRenderPass renderPass, VkExtent2D* pGranularity) {
	vkGetRenderAreaGranularity(device, renderPass, pGranularity);
}

void fn_vkGetRenderingAreaGranularity(VkDevice device, const VkRenderingAreaInfo* pRenderingAreaInfo, VkExtent2D* pGranularity) {
	vkGetRenderingAreaGranularity(device, pRenderingAreaInfo, pGranularity);
}

