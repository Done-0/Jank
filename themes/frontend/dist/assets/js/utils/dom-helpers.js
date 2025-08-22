/**
 * DOM 操作辅助工具
 */

export class DomHelpers {
  /**
   * 安全获取元素
   * @param {string} selector - 选择器
   * @returns {Element|null} DOM 元素
   */
  static getElement(selector) {
    return document.querySelector(selector);
  }

  /**
   * 安全获取元素列表
   * @param {string} selector - 选择器
   * @returns {NodeList} DOM 元素列表
   */
  static getElements(selector) {
    return document.querySelectorAll(selector);
  }

  /**
   * 切换类名
   * @param {Element} element - DOM 元素
   * @param {string} className - 类名
   */
  static toggleClass(element, className) {
    if (element) {
      element.classList.toggle(className);
    }
  }

  /**
   * 添加类名
   * @param {Element} element - DOM 元素
   * @param {string} className - 类名
   */
  static addClass(element, className) {
    if (element) {
      element.classList.add(className);
    }
  }

  /**
   * 移除类名
   * @param {Element} element - DOM 元素
   * @param {string} className - 类名
   */
  static removeClass(element, className) {
    if (element) {
      element.classList.remove(className);
    }
  }

  /**
   * 批量移除类名
   * @param {NodeList|Array} elements - DOM 元素列表
   * @param {string} className - 类名
   */
  static removeClassFromAll(elements, className) {
    elements.forEach(element => {
      this.removeClass(element, className);
    });
  }

  /**
   * 设置元素内容
   * @param {string} selector - 选择器
   * @param {string} content - 内容
   */
  static setContent(selector, content) {
    const element = this.getElement(selector);
    if (element) {
      element.innerHTML = content;
    }
  }

  /**
   * 显示元素
   * @param {Element} element - DOM 元素
   */
  static show(element) {
    if (element) {
      element.style.display = '';
    }
  }

  /**
   * 隐藏元素
   * @param {Element} element - DOM 元素
   */
  static hide(element) {
    if (element) {
      element.style.display = 'none';
    }
  }
}
