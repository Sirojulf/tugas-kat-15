document.addEventListener("DOMContentLoaded", () => {
    const searchInput = document.getElementById("searchInput");

    // Add placeholder animation
    let placeholderIndex = 0;
    const placeholders = [
        "Search for laptops...",
        "Search for smartphones...",
        "Search for smartwatches..."
    ];

    setInterval(() => {
        searchInput.placeholder = placeholders[placeholderIndex];
        placeholderIndex = (placeholderIndex + 1) % placeholders.length;
    }, 3000);

    // Prevent accidental form submission with empty input
    const form = document.querySelector(".search-form");
    form.addEventListener("submit", (event) => {
        if (!searchInput.value.trim()) {
            event.preventDefault();
            alert("Please enter a search term!");
        }
    });
});
