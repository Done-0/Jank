/**
 * 文章渲染组件
 */

export class PostRenderer {
  /**
   * 渲染文章列表
   * @param {Array} posts - 文章列表
   * @param {string} containerId - 容器ID
   */
  renderPosts(posts, containerId) {
    const container = document.getElementById(containerId);
    if (!container) return;

    if (!posts || posts.length === 0) {
      container.innerHTML = this.renderEmptyState('暂无文章');
      return;
    }

    const postsHtml = posts.map(post => this.createPostCard(post)).join('');
    container.innerHTML = postsHtml;
  }

  /**
   * 渲染文章详情
   * @param {Object} post - 文章数据
   * @param {string} containerId - 容器ID
   */
  renderPostDetail(post, containerId) {
    const container = document.getElementById(containerId);
    if (!container) return;

    if (!post) {
      container.innerHTML = this.renderEmptyState('文章不存在');
      return;
    }

    container.innerHTML = `
      <article class="max-w-4xl mx-auto">
        <header class="mb-8">
          <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-4">${post.title}</h1>
          <div class="flex items-center text-sm text-gray-600 dark:text-gray-400 space-x-4">
            <span>${post.category_name || '未分类'}</span>
            <time datetime="${post.created_at}">${new Date(post.created_at).toLocaleDateString()}</time>
          </div>
        </header>
        <div class="prose prose-lg dark:prose-invert max-w-none">
          ${post.html || post.markdown || ''}
        </div>
      </article>
    `;
  }

  /**
   * 创建文章卡片
   * @param {Object} post - 文章数据
   * @returns {string} 文章卡片HTML
   */
  createPostCard(post) {

    const truncatedDescription = post.description && post.description.length > 100 
      ? post.description.substring(0, 100) + '...' 
      : post.description || '';

    return `
      <article class="post-card" data-post-id="${post.id}">
        <div class="post-image-container">
          ${post.image ? `
            <img src="${post.image}" alt="${post.title}" class="post-image">
          ` : `
            <div class="post-image-placeholder"></div>
          `}
        </div>
        
        <div class="post-content">
          <div class="post-header">
            <h3 class="post-title">${post.title}</h3>
            ${post.category_name && post.category_name !== '未分类' ? `<span class="post-category">${post.category_name}</span>` : ''}
          </div>
          
          ${truncatedDescription ? `<p class="post-description">${truncatedDescription}</p>` : ''}
          
          <time class="post-date" datetime="${post.created_at}">${new Date(post.created_at).toLocaleDateString()}</time>
        </div>
      </article>
    `;
  }

  /**
   * 渲染空状态
   * @param {string} message - 空状态消息
   * @returns {string} 空状态HTML
   */
  renderEmptyState(message) {
    return `
      <div class="flex items-center justify-center py-12">
        <div class="text-center">
          <div class="text-gray-400 mb-2">📝</div>
          <p class="text-gray-600 dark:text-gray-400">${message}</p>
        </div>
      </div>
    `;
  }
}
