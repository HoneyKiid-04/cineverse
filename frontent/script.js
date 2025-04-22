
document.addEventListener('DOMContentLoaded', function() {
    // Mobile menu toggle
    const hamburger = document.getElementById('hamburger');
    const navbar = document.querySelector('.navbar');

    hamburger.addEventListener('click', function() {
        navbar.classList.toggle('active');
    });

    // DOM Elements
    const moviesContainer = document.getElementById('movies-container');
    const tvShowsContainer = document.getElementById('tv-shows-container');
    const searchInput = document.getElementById('search-input');
    const searchBtn = document.getElementById('search-btn');

    // API Base URL
    const API_BASE_URL = 'http://localhost:8080/api';

    // Fetch and render movies
    async function fetchAndRenderMovies() {
        try {
            const response = await fetch(`${API_BASE_URL}/movies`);
            if (!response.ok) {
                throw new Error('Failed to fetch movies');
            }
            const movies = await response.json();
            renderMovies(movies);
        } catch (error) {
            console.error('Error fetching movies:', error);
            moviesContainer.innerHTML = '<p class="error">Failed to load movies. Please try again later.</p>';
        }
    }

    // Render movies
    function renderMovies(movies) {
        moviesContainer.innerHTML = '';

        if (movies.length === 0) {
            moviesContainer.innerHTML = '<p class="no-results">No movies found.</p>';
            return;
        }

        movies.forEach(movie => {
            const movieCard = document.createElement('div');
            movieCard.className = 'movie-card';
            movieCard.dataset.id = movie.id;
            movieCard.innerHTML = `
                <img src="${movie.poster}" alt="${movie.title}" class="movie-poster">
                <div class="movie-info">
                    <h3 class="movie-title">${movie.title}</h3>
                    <p class="movie-year-genre">${movie.year} • ${movie.genre}</p>
                    <div class="movie-rating">
                        <i class="fas fa-star"></i>
                        <span>${movie.rating}</span>
                    </div>
                    <div class="movie-actions">
                        <button class="edit-btn" data-id="${movie.id}"><i class="fas fa-edit"></i></button>
                        <button class="delete-btn" data-id="${movie.id}"><i class="fas fa-trash"></i></button>
                    </div>
                </div>
            `;
            moviesContainer.appendChild(movieCard);
        });

        // Add event listeners to action buttons
        document.querySelectorAll('.edit-btn').forEach(btn => {
            btn.addEventListener('click', (e) => {
                e.stopPropagation();
                const movieId = btn.dataset.id;
                showEditMovieForm(movieId);
            });
        });

        document.querySelectorAll('.delete-btn').forEach(btn => {
            btn.addEventListener('click', (e) => {
                e.stopPropagation();
                const movieId = btn.dataset.id;
                deleteMovie(movieId);
            });
        });
    }

    // Show add movie form
    function showAddMovieForm() {
        const formHtml = `
            <div class="modal" id="movie-modal">
                <div class="modal-content">
                    <span class="close-btn">&times;</span>
                    <h2>Add New Movie</h2>
                    <form id="movie-form">
                        <input type="hidden" id="movie-id" value="">
                        <div class="form-group">
                            <label for="title">Title</label>
                            <input type="text" id="title" required>
                        </div>
                        <div class="form-group">
                            <label for="year">Year</label>
                            <input type="number" id="year" required>
                        </div>
                        <div class="form-group">
                            <label for="genre">Genre</label>
                            <input type="text" id="genre" required>
                        </div>
                        <div class="form-group">
                            <label for="poster">Poster URL</label>
                            <input type="url" id="poster" required>
                        </div>
                        <div class="form-group">
                            <label for="rating">Rating</label>
                            <input type="number" id="rating" step="0.1" min="0" max="10" required>
                        </div>
                        <button type="submit" class="submit-btn">Save</button>
                    </form>
                </div>
            </div>
        `;

        document.body.insertAdjacentHTML('beforeend', formHtml);

        const modal = document.getElementById('movie-modal');
        const closeBtn = document.querySelector('.close-btn');
        const form = document.getElementById('movie-form');

        // Close modal
        closeBtn.addEventListener('click', () => {
            modal.remove();
        });

        // Close when clicking outside
        modal.addEventListener('click', (e) => {
            if (e.target === modal) {
                modal.remove();
            }
        });

        // Form submission
        form.addEventListener('submit', async (e) => {
            e.preventDefault();
            await addMovie();
            modal.remove();
        });
    }

    // Show edit movie form
    async function showEditMovieForm(movieId) {
        try {
            const response = await fetch(`${API_BASE_URL}/movies/${movieId}`);
            if (!response.ok) {
                throw new Error('Failed to fetch movie');
            }
            const movie = await response.json();

            const formHtml = `
                <div class="modal" id="movie-modal">
                    <div class="modal-content">
                        <span class="close-btn">&times;</span>
                        <h2>Edit Movie</h2>
                        <form id="movie-form">
                            <input type="hidden" id="movie-id" value="${movie.id}">
                            <div class="form-group">
                                <label for="title">Title</label>
                                <input type="text" id="title" value="${movie.title}" required>
                            </div>
                            <div class="form-group">
                                <label for="year">Year</label>
                                <input type="number" id="year" value="${movie.year}" required>
                            </div>
                            <div class="form-group">
                                <label for="genre">Genre</label>
                                <input type="text" id="genre" value="${movie.genre}" required>
                            </div>
                            <div class="form-group">
                                <label for="poster">Poster URL</label>
                                <input type="url" id="poster" value="${movie.poster}" required>
                            </div>
                            <div class="form-group">
                                <label for="rating">Rating</label>
                                <input type="number" id="rating" step="0.1" min="0" max="10" value="${movie.rating}" required>
                            </div>
                            <button type="submit" class="submit-btn">Update</button>
                        </form>
                    </div>
                </div>
            `;

            document.body.insertAdjacentHTML('beforeend', formHtml);

            const modal = document.getElementById('movie-modal');
            const closeBtn = document.querySelector('.close-btn');
            const form = document.getElementById('movie-form');

            // Close modal
            closeBtn.addEventListener('click', () => {
                modal.remove();
            });

            // Close when clicking outside
            modal.addEventListener('click', (e) => {
                if (e.target === modal) {
                    modal.remove();
                }
            });

            // Form submission
            form.addEventListener('submit', async (e) => {
                e.preventDefault();
                await updateMovie(movieId);
                modal.remove();
            });

        } catch (error) {
            console.error('Error fetching movie:', error);
            alert('Failed to load movie details. Please try again.');
        }
    }

    // Add a new movie
    async function addMovie() {
        const movieData = {
            title: document.getElementById('title').value,
            year: parseInt(document.getElementById('year').value),
            genre: document.getElementById('genre').value,
            poster: document.getElementById('poster').value,
            rating: parseFloat(document.getElementById('rating').value)
        };

        try {
            const response = await fetch(`${API_BASE_URL}/movies`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(movieData)
            });

            if (!response.ok) {
                throw new Error('Failed to add movie');
            }

            const newMovie = await response.json();
            fetchAndRenderMovies();
            alert('Movie added successfully!');

        } catch (error) {
            console.error('Error adding movie:', error);
            alert('Failed to add movie. Please try again.');
        }
    }

    // Update a movie
    async function updateMovie(movieId) {
        const movieData = {
            title: document.getElementById('title').value,
            year: parseInt(document.getElementById('year').value),
            genre: document.getElementById('genre').value,
            poster: document.getElementById('poster').value,
            rating: parseFloat(document.getElementById('rating').value)
        };

        try {
            const response = await fetch(`${API_BASE_URL}/movies/${movieId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(movieData)
            });

            if (!response.ok) {
                throw new Error('Failed to update movie');
            }

            const updatedMovie = await response.json();
            fetchAndRenderMovies();
            alert('Movie updated successfully!');

        } catch (error) {
            console.error('Error updating movie:', error);
            alert('Failed to update movie. Please try again.');
        }
    }

    // Delete a movie
    async function deleteMovie(movieId) {
        if (!confirm('Are you sure you want to delete this movie?')) {
            return;
        }

        try {
            const response = await fetch(`${API_BASE_URL}/movies/${movieId}`, {
                method: 'DELETE'
            });

            if (!response.ok) {
                throw new Error('Failed to delete movie');
            }

            fetchAndRenderMovies();
            alert('Movie deleted successfully!');

        } catch (error) {
            console.error('Error deleting movie:', error);
            alert('Failed to delete movie. Please try again.');
        }
    }

    // Fetch and render TV shows
    async function fetchAndRenderTVShows() {
        try {
            const response = await fetch(`${API_BASE_URL}/tv-shows`);
            if (!response.ok) {
                throw new Error('Failed to fetch TV shows');
            }
            const tvShows = await response.json();
            renderTVShows(tvShows);
        } catch (error) {
            console.error('Error fetching TV shows:', error);
            tvShowsContainer.innerHTML = '<p class="error">Failed to load TV shows. Please try again later.</p>';
        }
    }

    // Render TV shows
    function renderTVShows(tvShows) {
        tvShowsContainer.innerHTML = '';

        if (tvShows.length === 0) {
            tvShowsContainer.innerHTML = '<p class="no-results">No TV shows found.</p>';
            return;
        }

        tvShows.forEach(show => {
            const tvShowCard = document.createElement('div');
            tvShowCard.className = 'tv-show-card';
            tvShowCard.innerHTML = `
                <img src="${show.poster}" alt="${show.title}" class="tv-show-poster">
                <div class="tv-show-info">
                    <h3 class="tv-show-title">${show.title}</h3>
                    <div class="tv-show-rating">
                        <i class="fas fa-star"></i>
                        <span>${show.rating}</span>
                    </div>
                </div>
            `;
            tvShowsContainer.appendChild(tvShowCard);
        });
    }

    // Search functionality
    searchBtn.addEventListener('click', async function() {
        const searchTerm = searchInput.value.trim();
        if (searchTerm) {
            try {
                const response = await fetch(`${API_BASE_URL}/search?q=${encodeURIComponent(searchTerm)}`);
                if (!response.ok) {
                    throw new Error('Search failed');
                }
                const results = await response.json();

                // Display search results
                if (results.movies.length > 0 || results.tv_shows.length > 0) {
                    renderMovies(results.movies);
                    renderTVShows(results.tv_shows);

                    // Scroll to movies section
                    document.querySelector('#movies').scrollIntoView({ behavior: 'smooth' });
                } else {
                    moviesContainer.innerHTML = '<p class="no-results">No movies found matching your search.</p>';
                    tvShowsContainer.innerHTML = '<p class="no-results">No TV shows found matching your search.</p>';
                }
            } catch (error) {
                console.error('Search error:', error);
                alert('Search failed. Please try again.');
            }
        }
    });

    // Add "Add Movie" button functionality
    const addMovieBtn = document.createElement('button');
    addMovieBtn.className = 'add-movie-btn';
    addMovieBtn.innerHTML = '<i class="fas fa-plus"></i> Add Movie';
    addMovieBtn.addEventListener('click', showAddMovieForm);

    const moviesHeader = document.querySelector('.movies-section h2');
    moviesHeader.insertAdjacentElement('afterend', addMovieBtn);

    // Smooth scrolling for anchor links
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function(e) {
            e.preventDefault();

            const targetId = this.getAttribute('href');
            const targetElement = document.querySelector(targetId);

            if (targetElement) {
                window.scrollTo({
                    top: targetElement.offsetTop - 80,
                    behavior: 'smooth'
                });

                // Close mobile menu if open
                if (navbar.classList.contains('active')) {
                    navbar.classList.remove('active');
                }
            }
        });
    });

    // Scroll effect for navbar
    window.addEventListener('scroll', function() {
        if (window.scrollY > 50) {
            navbar.style.backgroundColor = 'rgba(10, 10, 18, 0.95)';
            navbar.style.padding = '1rem 5%';
        } else {
            navbar.style.backgroundColor = 'rgba(15, 15, 26, 0.9)';
            navbar.style.padding = '1.5rem 5%';
        }
    });

    // Initial load
    fetchAndRenderMovies();
    fetchAndRenderTVShows();
});
