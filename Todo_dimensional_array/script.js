// Fetch news data from API
fetch('https://newsapi.org/v2/top-headlines?country=us&apiKey=YOUR_API_KEY')
    .then(response => response.json())
    .then(data => {
        const newsContainer = document.getElementById('news-container');
        data.articles.forEach(article => {
            const newsCard = document.createElement('div');
            newsCard.classList.add('news-card');
            newsCard.innerHTML = `
                <img src="${article.urlToImage}" alt="${article.title}">
                <h2>${article.title}</h2>
                <p>${article.description}</p>
                <button>Read More</button>
            `;
            newsContainer.appendChild(newsCard);
        });
    });
    