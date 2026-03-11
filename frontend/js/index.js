// ===================== 1. 加载层逻辑 =====================
const loadingOverlay = document.getElementById('loading-overlay');

// 监听页面加载完成事件
window.addEventListener('load', function() {
    loadingOverlay.classList.add('hidden');
    // 延迟移除元素（让过渡动画完成）
    setTimeout(() => {
        loadingOverlay.remove();
    }, 500);
});

// 容错处理：页面加载超时也隐藏加载层
setTimeout(() => {
    if (!loadingOverlay.classList.contains('hidden')) {
        loadingOverlay.classList.add('hidden');
    }
}, 10000);

// ===================== 2. 登录状态检测 + 用户信息展示（核心优化） =====================
/**
 * 清除本地登录状态
 */
function clearLoginState() {
    localStorage.removeItem('token');
    localStorage.removeItem('refreshToken');
    localStorage.removeItem('userId');
}

/**
 * 初始化用户信息（导航栏替换）
 */
async function initUserInfo() {
    const token = localStorage.getItem('token');
    const userId = localStorage.getItem('userId');
    const userLink = document.getElementById('userLink');
    console.log('uid',userId);
    console.log('token',token);
    // 未登录：保持默认
    if (!token || !userId) {
        userLink.href = 'login.html';
        userLink.innerHTML = '登录/注册';
        // 重置样式（避免头像样式残留）
        userLink.style.display = 'inline';
        userLink.style.alignItems = 'unset';
        return;
    }

    // 已登录：调用用户信息接口
    try {
        const response = await fetch(`http://10.41.189.139:8080/user/info/${userId}`, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        });

        const data = await response.json();
        // 接口请求成功
        if (data.status === 10000) {
            const userInfo = data.date.date; // 适配后端返回的嵌套date结构
            let userHtml = '';

            // 处理头像展示逻辑
            if (userInfo.avatar && userInfo.avatar !== '{:url}') {
                // 有有效头像：显示头像 + 昵称（优先）/我的
                userHtml = `
                    <div style="display: flex; align-items: center; gap: 8px;">
                        <img src="${userInfo.avatar}" class="user-avatar" alt="${userInfo.nickName || '头像'}" 
                             style="width: 28px; height: 28px; border-radius: 50%; object-fit: cover; border: 1px solid #eee;">
                        <span>${userInfo.nickName || '我的'}</span>
                    </div>
                `;
            } else {
                // 无有效头像：显示默认文字
                userHtml = `<span>${userInfo.nickName || '我的'}</span>`;
            }

            // 替换链接和内容
            userLink.href = 'mine.html';
            userLink.innerHTML = userHtml;
            // 调整样式适配
            userLink.style.display = 'flex';
            userLink.style.alignItems = 'center';
        } else {
            // 接口返回失败，清除登录状态
            clearLoginState();
            userLink.href = 'login.html';
            userLink.innerHTML = '登录/注册';
        }
    } catch (error) {
        console.error('获取用户信息失败：', error);
        clearLoginState();
        userLink.href = 'login.html';
        userLink.innerHTML = '登录/注册';
    }
}

// ===================== 3. 商品跳转函数 =====================
function jumpToProductDetail(productId, baseUrl = "goods.html") {
    const targetUrl = `${baseUrl}?id=${productId}`;
    console.log('跳转到商品详情页：', targetUrl);
    window.open(targetUrl);
}

// ===================== 4. 商品列表渲染 =====================
document.addEventListener('DOMContentLoaded', async function() {
    // 优先初始化用户信息（保证导航栏先加载）
    await initUserInfo();

    const productList = document.getElementById('productList');
    try {
        const response = await fetch('http://10.41.189.139:8080/product/list');
        if (!response.ok) throw new Error(`HTTP错误，状态码：${response.status}`);

        const result = await response.json();
        if (result.status !== 10000) {
            throw new Error(`接口返回失败：${result.info || '未知错误'}`);
        }

        const products = result.date || [];
        productList.innerHTML = '';

        if (products.length === 0) {
            productList.innerHTML = "<div style='width: 100%; text-align: center;'>暂无商品</div>";
            return;
        }

        products.forEach(product => {
            const linkUrl = product.link || "#";
            const productName = product.name || "未命名商品";
            const productId = product.productId || 0;

            const shopItem = document.createElement('div');
            shopItem.className = 'shop';
            shopItem.style.cursor = 'pointer';
            shopItem.onclick = () => jumpToProductDetail(productId);

            shopItem.innerHTML = `
                <div class="tu-pian">
                    <img src="${linkUrl}" alt="${productName}">
                </div>
                <div class="wen-zi">${productName}</div>
            `;

            productList.appendChild(shopItem);
        });
    } catch (error) {
        productList.innerHTML = `<div style='width: 100%; text-align: center; color: red;'>加载商品失败：${error.message}</div>`;
        console.error("商品列表加载失败：", error);
    }
});

// ===================== 5. 轮播图逻辑 =====================
let currentIndex = 0;
let dbSlides = [];
const slidesContainer = document.getElementById('carousel-slides');
const indicatorsContainer = document.getElementById('carousel-indicators');
const prevBtn = document.getElementById('prev-btn');
const nextBtn = document.getElementById('next-btn');

async function getSlidesFromDB() {
    try {
        await new Promise(resolve => setTimeout(resolve, 600));
        return [
            { productId: 1, imgSrc: "assets/1728899126126.jpg", imgAlt: "商品1", name: "商品1 - 爆款手机" },
            { productId: 2, imgSrc: "assets/1728899126126.jpg", imgAlt: "商品2", name: "商品2 - 无线耳机" },
            { productId: 3, imgSrc: "assets/1728899126126.jpg", imgAlt: "商品3", name: "商品3 - 智能手表" },
            { productId: 4, imgSrc: "assets/1728899126126.jpg", imgAlt: "商品4", name: "商品4 - 平板电脑" },
            { productId: 5, imgSrc: "assets/1728899126126.jpg", imgAlt: "商品5", name: "商品5 - 充电宝" },
            { productId: 6, imgSrc: "assets/1728899126126.jpg", imgAlt: "商品6", name: "商品6 - 蓝牙耳机" }
        ];
    } catch (err) {
        console.error('获取轮播数据失败:', err);
        return [
            { productId: 0, imgSrc: "https://via.placeholder.com/800x600?text=加载失败&bgcolor=f5f5f5&color=888", imgAlt: "错误", name: "加载失败" }
        ];
    }
}

function initCarousel(slides) {
    dbSlides = slides;
    currentIndex = 0;
    slidesContainer.innerHTML = '';
    indicatorsContainer.innerHTML = '';

    slides.forEach((slide, index) => {
        const slideEl = document.createElement('div');
        slideEl.className = 'carousel-slide';

        const img = document.createElement('img');
        img.className = 'slide-img';
        img.src = slide.imgSrc;
        img.alt = slide.imgAlt;
        img.onerror = function() {
            this.src = 'https://via.placeholder.com/800x600?text=图片加载失败&bgcolor=f5f5f5&color=888';
        };

        img.style.cursor = 'pointer';
        img.onclick = () => jumpToProductDetail(slide.productId);

        slideEl.appendChild(img);
        slidesContainer.appendChild(slideEl);

        const indicator = document.createElement('div');
        indicator.className = `indicator ${index === 0 ? 'active' : ''}`;
        indicator.dataset.index = index;
        indicator.onclick = () => goToSlide(index);
        indicatorsContainer.appendChild(indicator);
    });

    updateCarousel();
}

function updateCarousel() {
    slidesContainer.style.transform = `translateX(-${currentIndex * 100}%)`;
    document.querySelectorAll('.indicator').forEach((el, idx) => {
        el.classList.toggle('active', idx === currentIndex);
    });
}

function goToSlide(index) {
    currentIndex = index < 0
        ? dbSlides.length - 1
        : index >= dbSlides.length
            ? 0
            : index;
    updateCarousel();
}

prevBtn.onclick = () => goToSlide(currentIndex - 1);
nextBtn.onclick = () => goToSlide(currentIndex + 1);

let autoPlay;
function startAutoPlay() {
    autoPlay = setInterval(() => goToSlide(currentIndex + 1), 5000);
}

// ===================== 6. 页面初始化（整合所有逻辑） =====================
window.onload = async () => {
    // 重新初始化用户信息（确保页面加载完成后再次校验）
    await initUserInfo();

    // 初始化轮播图
    const slides = await getSlidesFromDB();
    initCarousel(slides);
    startAutoPlay();

    // 轮播图鼠标悬停暂停/恢复自动播放
    const carousel = document.querySelector('.carousel');
    if (carousel) {
        carousel.onmouseenter = () => clearInterval(autoPlay);
        carousel.onmouseleave = () => startAutoPlay();
    }
};