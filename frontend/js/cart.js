// 工具函数：从 localStorage 获取 token
function getToken() {
    return localStorage.getItem('token') || '';
}

// 页面加载完成后获取购物车数据
document.addEventListener('DOMContentLoaded', () => {
    fetchCartData();
    // 绑定结算按钮点击事件
    document.querySelector('.checkout-btn').addEventListener('click', handleCheckout);
});

// 1. 获取购物车列表（调用 /product/cart 接口）
async function fetchCartData() {
    try {
        const res = await fetch('http://10.41.189.139:8080/product/cart', {
            headers: {
                'Authorization': `Bearer ${getToken()}`
            }
        });
        const result = await res.json();

        if (result.status === 10000) {
            // 修复：接口返回字段是 date，不是 data，增加空数组兜底
            const cartItems = result.date || [];
            renderCartList(cartItems);
        } else {
            console.error('获取购物车失败:', result.info);
            renderCartList([]);
        }
    } catch (err) {
        console.error('获取购物车数据失败:', err);
        renderCartList([]);
    }
}


// 3. 渲染购物车列表
function renderCartList(cartItems) {
    // 兜底判断：确保是数组类型
    if (!Array.isArray(cartItems)) {
        cartItems = [];
    }
    console.log(cartItems);
    const cartListEl = document.getElementById('cart-list');
    cartListEl.innerHTML = '';

    let totalPrice = 0;

    // 购物车为空时显示提示
    if (cartItems.length === 0) {
        cartListEl.innerHTML = `
            <div style="text-align: center; padding: 50px 0; color: #999;">
                <p>您的购物车还是空的哦~</p>
                <button style="margin-top: 20px; padding: 8px 16px; background: #6610f2; color: white; border: none; border-radius: 4px; cursor: pointer;" onclick="window.location.href='index.html'">去逛逛</button>
            </div>
        `;
        document.getElementById('total-price').textContent = '0.00';
        return;
    }

    cartItems.forEach(item => {
        const itemTotal = item.price * item.num;
        totalPrice += itemTotal;
        const imageUrl = `..${item.link}`;
        console.log("aaaa",imageUrl);
        const itemHtml = `
            <div class="cart-item" data-product-id="${item.productId}">
                <img src="${imageUrl}" alt="${item.name}" class="cart-item-img">
                <div class="cart-item-info">
                    <div class="cart-item-name">${item.name}</div>
                    <div class="cart-item-type">类型：${item.type}</div>
                    <div class="cart-item-price">单价：¥${item.price.toFixed(2)}</div>
                </div>
                <div class="cart-item-quantity">
                    <button class="quantity-btn minus">-</button>
                    <input type="number" class="quantity-input" value="${item.num}" min="1" readonly>
                    <button class="quantity-btn plus">+</button>
                </div>
                <div class="cart-item-total">¥${itemTotal.toFixed(2)}</div>
                <div class="delete-btn">删除</div>
            </div>
        `;
        cartListEl.insertAdjacentHTML('beforeend', itemHtml);
    });

    // 更新总价
    document.getElementById('total-price').textContent = totalPrice.toFixed(2);

    // 绑定事件
    bindCartEvents();
}

// 4. 绑定购物车交互事件（数量增减、删除）
function bindCartEvents() {
    const cartItems = document.querySelectorAll('.cart-item');

    cartItems.forEach(item => {
        const minusBtn = item.querySelector('.minus');
        const plusBtn = item.querySelector('.plus');
        const deleteBtn = item.querySelector('.delete-btn');
        const quantityInput = item.querySelector('.quantity-input');
        const productId = item.dataset.productId;

        // 数量减
        minusBtn.addEventListener('click', () => {
            let num = parseInt(quantityInput.value);
            if (num > 1) {
                num--;
                quantityInput.value = num;
                updateItemTotal(item, num);
                updateTotalPrice();
                // 这里可调用后端更新数量接口（需后端实现）
                // updateCartQuantity(productId, num);
            }
        });

        // 数量加
        plusBtn.addEventListener('click', () => {
            let num = parseInt(quantityInput.value);
            num++;
            quantityInput.value = num;
            updateItemTotal(item, num);
            updateTotalPrice();
            // 这里可调用后端更新数量接口（需后端实现）
            // updateCartQuantity(productId, num);
        });

        // 删除商品
        deleteBtn.addEventListener('click', () => {
            if (confirm('确定要删除该商品吗？')) {
                // 这里调用后端删除购物车项接口（需后端实现）
                // deleteCartItem(productId);
                item.remove();
                updateTotalPrice();
                // 重新检查购物车是否为空
                checkCartEmpty();
            }
        });
    });
}

// 5. 更新单个商品小计
function updateItemTotal(item, num) {
    const price = parseFloat(item.querySelector('.cart-item-price').textContent.replace('单价：¥', ''));
    const totalEl = item.querySelector('.cart-item-total');
    totalEl.textContent = `¥${(price * num).toFixed(2)}`;
}

// 6. 更新总价
function updateTotalPrice() {
    const totalEls = document.querySelectorAll('.cart-item-total');
    let total = 0;
    totalEls.forEach(el => {
        total += parseFloat(el.textContent.replace('¥', ''));
    });
    document.getElementById('total-price').textContent = total.toFixed(2);
}

// 7. 检查购物车是否为空
function checkCartEmpty() {
    const cartItems = document.querySelectorAll('.cart-item');
    if (cartItems.length === 0) {
        renderCartList([]); // 重新渲染空购物车提示
    }
}

// 8. 结算功能（核心新增）
async function handleCheckout() {
    // 1. 先检查购物车是否为空
    const cartItems = document.querySelectorAll('.cart-item');
    if (cartItems.length === 0) {
        alert('您的购物车为空，无法结算！');
        return;
    }

    // 2. 检查用户是否登录（通过token判断）
    const token = getToken();
    if (!token) {
        if (confirm('您还未登录，是否前往登录页？')) {
            window.location.href = 'login.html';
        }
        return;
    }

    try {
        // 3. 调用下单接口
        const res = await fetch('http://10.41.189.139:8080/operate/order', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json' // 根据后端要求调整，若用form-data则改为multipart/form-data
            }
            // 若后端需要传递购物车商品信息，可添加body
            // body: JSON.stringify({
            //     productIds: Array.from(cartItems).map(item => item.dataset.productId)
            // })
        });

        const result = await res.json();

        // 4. 处理下单结果
        if (result.status === 10000) {
            const orderId = result.date.order_id;
            alert(`下单成功！您的订单号是：${orderId}`);
            // 下单成功后清空购物车（可选，根据业务逻辑调整）
            // clearCart();
            // 跳转到订单详情页
            //window.location.href = `order_detail.html?order_id=${orderId}`;
            window.location.href = 'index.html';
        } else {
            alert('下单失败：' + (result.info || '未知错误'));
        }
    } catch (err) {
        console.error('下单请求失败:', err);
        alert('下单失败，请检查网络或重试！');
    }
}

// 9. 清空购物车（可选）
function clearCart() {
    document.getElementById('cart-list').innerHTML = '';
    document.getElementById('total-price').textContent = '0.00';
    renderCartList([]);
}