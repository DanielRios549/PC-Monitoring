document.addEventListener('alpine:init', () => {
    const key = 'theme'

    Alpine.store(key, {
        current: localStorage.getItem(key) || 'dark',

        toggle() {
            let newTheme = 'light'
    
            if (this.current === 'light') {
                newTheme = 'dark'
            }

            this.current = newTheme
            localStorage.setItem(key, newTheme)
        }
    })
})