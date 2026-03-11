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

// ===================== 新增：右侧热门推荐商品渲染 =====================
/**
 * 从后端获取热门推荐商品（3个）
 */
async function getHotRecommendFromDB() {
    try {
        // 请求商品列表接口（可替换为专门的热门商品接口）
        const response = await fetch('http://10.41.189.139:8080/product/list', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error(`接口请求失败：${response.status}`);
        }

        const result = await response.json();
        // 校验接口返回格式
        if (result.status !== 10000 || !Array.isArray(result.date)) {
            throw new Error('接口返回数据格式异常');
        }

        const allProducts = result.date;
        // 随机筛选3个商品作为热门推荐（也可根据销量/热度排序）
        const hotProducts = allProducts.length <= 3
            ? allProducts
            : allProducts.sort(() => Math.random() - 0.5).slice(0, 3);

        // 格式化数据结构
        return hotProducts.map(product => ({
            productId: product.productId || product.id || 0,
            imgSrc: product.link || product.imageUrl || 'https://picsum.photos/60/60?random=100',
            name: product.name || '未命名商品',
            price: product.price || '0.00'
        }));

    } catch (err) {
        console.error('获取热门推荐数据失败:', err);
        // 异常兜底：返回3个默认商品
        return [
            { productId: 101, imgSrc: 'https://picsum.photos/60/60?random=10', name: '爆款无线耳机', price: '¥99' },
            { productId: 102, imgSrc: 'https://picsum.photos/60/60?random=11', name: '夏季纯棉T恤', price: '¥59' },
            { productId: 103, imgSrc: 'https://picsum.photos/60/60?random=12', name: '考研英语真题', price: '¥39' }
        ];
    }
}

/**
 * 渲染右侧热门推荐列表
 */
async function renderHotRecommend() {
    const hotRecommendList = document.querySelector('.hot-recommend-list');
    if (!hotRecommendList) return;

    // 清空原有静态内容
    hotRecommendList.innerHTML = '';

    // 获取热门商品数据
    const hotProducts = await getHotRecommendFromDB();

    // 渲染每个热门商品项
    hotProducts.forEach(product => {
        const hotItem = document.createElement('li');
        hotItem.className = 'hot-item';

        const link = document.createElement('a');
        link.href = 'javascript:void(0)';
        link.onclick = () => jumpToProductDetail(product.productId);

        link.innerHTML = `
            <img src="${product.imgSrc}" alt="${product.name}" class="hot-item-img" onerror="this.src='https://picsum.photos/60/60?random=200'">
            <div class="hot-item-info">
                <p class="hot-item-name">${product.name}</p>
                <p class="hot-item-price">¥${product.price}</p>
            </div>
        `;

        hotItem.appendChild(link);
        hotRecommendList.appendChild(hotItem);
    });
}

// ===================== 4. 商品列表渲染 =====================
// ===================== 7. 商品列表加载更多逻辑 =====================
document.addEventListener('DOMContentLoaded', function() {
    const productList = document.getElementById('productList');
    const loadMoreBtn = document.getElementById('loadMoreBtn');
    const itemsPerRow = 5; // 每行商品数量
    const defaultRows = 2; // 默认展示行数
    const defaultShowCount = defaultRows * itemsPerRow; // 默认展示数量

    // 初始化商品列表高度（2行）
    productList.style.maxHeight = `${(200 + 30) * defaultRows}px`; // 200px商品高度 + 30px间距

    // 加载更多按钮点击事件
    loadMoreBtn.addEventListener('click', function() {
        // 展示全部商品（移除高度限制）
        productList.style.maxHeight = 'none';
        // 隐藏加载更多按钮
        loadMoreBtn.classList.add('hidden');
    });

    // 监听商品列表渲染完成后判断是否需要显示加载更多按钮
    function checkLoadMoreButton(products) {
        if (products.length > defaultShowCount) {
            loadMoreBtn.classList.remove('hidden');
        } else {
            loadMoreBtn.classList.add('hidden');
            productList.style.maxHeight = 'none'; // 商品不足时直接展示全部
        }
    }

    // 重写商品列表渲染逻辑（整合原有逻辑）
    async function renderProductList() {
        try {
            const response = await fetch('http://10.41.189.139:8080/product/list');
            if (!response.ok) throw new Error(`HTTP错误，状态码：${response.status}`);

            const result = await response.json();
            if (result.status !== 10000) {
                throw new Error(`接口返回失败：${result.info || '未知错误'}`);
            }

            let products = result.date || [];
            function shuffleArray(array) {
                // 先创建数组副本，避免修改原数组
                const newArray = [...array];
                for (let i = newArray.length - 1; i > 0; i--) {
                    const j = Math.floor(Math.random() * (i + 1));
                    [newArray[i], newArray[j]] = [newArray[j], newArray[i]];
                }
                return newArray;
            }
            // 执行打乱操作
            products = shuffleArray(products);
            productList.innerHTML = '';

            if (products.length === 0) {
                productList.innerHTML = "<div style='width: 100%; text-align: center;'>暂无商品</div>";
                loadMoreBtn.classList.add('hidden');
                return;
            }

            // 渲染商品列表
            products.forEach(product => {
                const linkUrl = product.link || "#";
                const productName = product.name || "未命名商品";
                const productId = product.productId || 0;
                console.log(linkUrl);
                const shopItem = document.createElement('div');
                shopItem.className = 'shop';
                shopItem.style.cursor = 'pointer';
                shopItem.onclick = () => jumpToProductDetail(productId);

                shopItem.innerHTML = `
                    <div class="tu-pian">
                        <img src="${linkUrl}" alt="${productName}">
                    </div>
                    <div class="wen-zi">
                        <p>${productName}</p>
                        <p>¥${product.price || '0.00'}</p>
                    </div>
                `;

                productList.appendChild(shopItem);
            });

            // 判断是否显示加载更多按钮
            checkLoadMoreButton(products);

        } catch (error) {
            productList.innerHTML = `<div style='width: 100%; text-align: center; color: red;'>加载商品失败：${error.message}</div>`;
            loadMoreBtn.classList.add('hidden');
            console.error("商品列表加载失败：", error);
        }
    }

    // 初始化调用商品列表渲染
    renderProductList();
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
        // 请求商品列表接口获取轮播图数据
        const response = await fetch('http://10.41.189.139:8080/product/list', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error(`接口请求失败：${response.status}`);
        }

        const result = await response.json();
        if (result.status !== 10000 || !Array.isArray(result.date)) {
            throw new Error('接口返回数据格式异常');
        }

        const allProducts = result.date;
        // 随机筛选6个商品作为轮播图
        const randomSlides = allProducts.length <= 6
            ? allProducts
            : allProducts.sort(() => Math.random() - 0.5).slice(0, 6);

        return randomSlides.map(product => ({
            productId: product.productId || product.id || 0,
            imgSrc: product.link || product.imageUrl || 'https://picsum.photos/800/400?random=10',
            imgAlt: product.name || `商品${product.productId}`,
            name: product.name || '未命名商品'
        }));

    } catch (err) {
        console.error('获取轮播数据失败:', err);
        // 异常兜底
        return Array.from({ length: 6 }, (_, index) => ({
            productId: index + 1,
            imgSrc: `https://picsum.photos/800/400?random=${index + 10}`,
            imgAlt: `默认轮播图${index + 1}`,
            name: `默认商品${index + 1}`
        }));
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
            this.src = `https://picsum.photos/800/400?random=${index + 20}`;
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

    // 初始化右侧热门推荐
    await renderHotRecommend();

    // 轮播图鼠标悬停暂停/恢复自动播放
    const carousel = document.querySelector('.carousel');
    if (carousel) {
        carousel.onmouseenter = () => clearInterval(autoPlay);
        carousel.onmouseleave = () => startAutoPlay();
    }
};