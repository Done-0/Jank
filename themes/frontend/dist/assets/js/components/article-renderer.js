/**
 * 文章详情渲染器
 * 负责渲染文章详情页面和目录导航
 */

export class ArticleRenderer {
  constructor() {
    this.currentArticle = null;
  }

  /**
   * 渲染文章详情
   * @param {Object} article - 文章对象
   */
  renderArticle(article) {
    this.currentArticle = article;
    
    // 隐藏左侧分类侧边栏和右侧作者卡片
    this.hideSidebars();
    
    // 更新文章标题
    const titleElement = document.getElementById('article-title');
    if (titleElement) {
      titleElement.textContent = article.title;
    }

    // 更新文章分类
    const categoryElement = document.getElementById('article-category');
    if (categoryElement && article.category_name && article.category_name !== '未分类') {
      categoryElement.textContent = article.category_name;
      categoryElement.style.display = 'inline-block';
    } else if (categoryElement) {
      categoryElement.style.display = 'none';
    }

    // 更新文章创建日期
    const createdDateElement = document.getElementById('article-created-date');
    if (createdDateElement) {
      const createdDate = new Date(article.created_at);
      createdDateElement.textContent = `发布于 ${createdDate.toLocaleDateString('zh-CN')}`;
      createdDateElement.setAttribute('datetime', article.created_at);
    }

    // 更新文章更新日期（如果存在且与创建日期不同）
    const updatedDateElement = document.getElementById('article-updated-date');
    const dateSeparator = document.querySelector('.date-separator');
    
    if (updatedDateElement && article.updated_at) {
      const createdDate = new Date(article.created_at);
      const updatedDate = new Date(article.updated_at);
      
      // 只有当更新时间与创建时间不同时才显示更新日期
      if (updatedDate.getTime() !== createdDate.getTime()) {
        updatedDateElement.textContent = `更新于 ${updatedDate.toLocaleDateString('zh-CN')}`;
        updatedDateElement.setAttribute('datetime', article.updated_at);
        updatedDateElement.style.display = 'inline-block';
        if (dateSeparator) dateSeparator.style.display = 'inline-block';
      } else {
        updatedDateElement.style.display = 'none';
        if (dateSeparator) dateSeparator.style.display = 'none';
      }
    } else {
      if (dateSeparator) dateSeparator.style.display = 'none';
    }

    // 渲染文章HTML内容
    const contentElement = document.getElementById('article-html-content');
    if (contentElement && article.html) {
      // 使用主流HTML渲染库处理内容
      this.renderHTMLContent(contentElement, article.html);
      
      // 为标题添加ID，用于目录导航
      this.addHeadingIds(contentElement);
      
      // 生成目录
      this.generateTOC(contentElement);
    }
  }

  /**
   * 隐藏侧边栏（仅在文章详情页）
   */
  hideSidebars() {
    const mainContainer = document.querySelector('.main-container');
    const leftSidebar = document.querySelector('.sidebar-left');
    const rightSidebar = document.querySelector('.sidebar-right');
    
    // 切换到文章详情布局模式
    if (mainContainer) {
      mainContainer.classList.add('article-detail-mode');
    }
    
    if (leftSidebar) {
      leftSidebar.style.display = 'none';
    }
    
    if (rightSidebar) {
      rightSidebar.style.display = 'none';
    }
  }

  /**
   * 显示侧边栏（返回首页时）
   */
  showSidebars() {
    const mainContainer = document.querySelector('.main-container');
    const leftSidebar = document.querySelector('.sidebar-left');
    const rightSidebar = document.querySelector('.sidebar-right');
    
    // 恢复正常布局模式
    if (mainContainer) {
      mainContainer.classList.remove('article-detail-mode');
    }
    
    if (leftSidebar) {
      leftSidebar.style.display = 'block';
    }
    
    if (rightSidebar) {
      rightSidebar.style.display = 'block';
    }
  }

  /**
   * 使用主流HTML渲染库渲染内容
   * @param {HTMLElement} container - 容器元素
   * @param {string} htmlContent - HTML内容
   */
  renderHTMLContent(container, htmlContent) {
    // 直接设置HTML内容
    container.innerHTML = htmlContent;
    
    // 使用highlight.js高亮代码块
    if (typeof hljs !== 'undefined') {
      const codeBlocks = container.querySelectorAll('pre code');
      codeBlocks.forEach(block => {
        hljs.highlightElement(block);
      });
    }
    
    // 为链接添加target="_blank"
    const links = container.querySelectorAll('a[href^="http"]');
    links.forEach(link => {
      link.setAttribute('target', '_blank');
      link.setAttribute('rel', 'noopener noreferrer');
    });
    
    // 为表格添加响应式样式
    const tables = container.querySelectorAll('table');
    tables.forEach(table => {
      const wrapper = document.createElement('div');
      wrapper.className = 'table-wrapper';
      table.parentNode.insertBefore(wrapper, table);
      wrapper.appendChild(table);
    });
  }

  /**
   * 为标题添加ID
   * @param {HTMLElement} container - 容器元素
   */
  addHeadingIds(container) {
    const headings = container.querySelectorAll('h1, h2, h3, h4, h5, h6');
    headings.forEach((heading, index) => {
      if (!heading.id) {
        const text = heading.textContent.trim();
        const id = `heading-${index}-${text.toLowerCase().replace(/[^a-z0-9\u4e00-\u9fa5]/g, '-').replace(/-+/g, '-').replace(/^-|-$/g, '')}`;
        heading.id = id;
      }
    });
  }

  /**
   * 生成目录导航
   * @param {HTMLElement} contentElement - 文章内容容器
   */
  generateTOC(contentElement) {
    const tocNav = document.getElementById('article-toc-nav');
    if (!tocNav) return;

    const headings = contentElement.querySelectorAll('h1, h2, h3, h4, h5, h6');
    if (headings.length === 0) {
      tocNav.innerHTML = '<p class="toc-empty">暂无目录</p>';
      return;
    }

    let tocHTML = '';
    headings.forEach(heading => {
      const level = heading.tagName.toLowerCase();
      const text = heading.textContent.trim();
      const id = heading.id;
      
      tocHTML += `<a href="#${id}" class="toc-link toc-${level}" data-target="${id}">${text}</a>`;
    });

    tocNav.innerHTML = tocHTML;

    // 绑定目录点击事件
    this.bindTOCEvents();
    
    // 初始化滚动监听
    this.initScrollSpy();
  }

  /**
   * 绑定目录点击事件
   */
  bindTOCEvents() {
    const tocLinks = document.querySelectorAll('.toc-link');
    tocLinks.forEach(link => {
      link.addEventListener('click', (e) => {
        e.preventDefault();
        const targetId = link.getAttribute('data-target');
        const targetElement = document.getElementById(targetId);
        
        if (targetElement) {
          // 平滑滚动到目标位置
          targetElement.scrollIntoView({
            behavior: 'smooth',
            block: 'start'
          });
          
          // 更新活动状态
          this.updateActiveTOCLink(link);
        }
      });
    });
  }

  /**
   * 初始化滚动监听，高亮当前可见的标题
   */
  initScrollSpy() {
    const headings = document.querySelectorAll('#article-html-content h1, #article-html-content h2, #article-html-content h3, #article-html-content h4, #article-html-content h5, #article-html-content h6');
    const tocLinks = document.querySelectorAll('.toc-link');
    
    if (headings.length === 0 || tocLinks.length === 0) return;

    const observer = new IntersectionObserver((entries) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          const id = entry.target.id;
          const activeLink = document.querySelector(`.toc-link[data-target="${id}"]`);
          if (activeLink) {
            this.updateActiveTOCLink(activeLink);
          }
        }
      });
    }, {
      rootMargin: '-20% 0px -70% 0px',
      threshold: 0
    });

    headings.forEach(heading => {
      observer.observe(heading);
    });
  }

  /**
   * 更新目录中的活动链接
   * @param {HTMLElement} activeLink - 当前活动的链接
   */
  updateActiveTOCLink(activeLink) {
    // 移除所有活动状态
    document.querySelectorAll('.toc-link.active').forEach(link => {
      link.classList.remove('active');
    });
    
    // 添加当前活动状态
    activeLink.classList.add('active');
  }

  /**
   * 显示文章详情页
   */
  showArticleView() {
    const homeView = document.getElementById('home-view');
    const articleView = document.getElementById('article-view');
    
    if (homeView) homeView.style.display = 'none';
    if (articleView) articleView.style.display = 'block';
  }

  /**
   * 隐藏文章详情页，返回首页
   */
  hideArticleView() {
    const homeView = document.getElementById('home-view');
    const articleView = document.getElementById('article-view');
    
    if (homeView) homeView.style.display = 'block';
    if (articleView) articleView.style.display = 'none';
    
    this.currentArticle = null;
  }
}
