document.addEventListener('alpine:init', () => {
    Alpine.store('darkMode', {
        on: true,

        toggle() {
            this.on = !this.on
        }
    })
})