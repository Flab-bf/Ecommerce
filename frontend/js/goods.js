// 在 goods.html 的 <script> 标签中执行
function getUrlParam(paramName) {
    // 获取当前页面的URL参数部分（?id=xxx 这部分）
    const searchParams = new URLSearchParams(window.location.search);
    // 返回指定参数名的值
    return searchParams.get(paramName);
}

// 提取id参数并使用
const productId = getUrlParam('id');

// 验证提取结果
if (productId) {
    console.log('提取到的商品ID：', productId); // 输出示例：123
    // 后续逻辑：用productId调用后端接口获取商品详情、渲染页面等
} else {
    console.log('未获取到商品ID，可能是直接访问详情页');
    // 可选：跳转到商品列表页或提示用户
    // window.location.href = 'product_list.html';
}