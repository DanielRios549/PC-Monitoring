document.addEventListener('alpine:init', () => {
    const themeKey = 'theme'
    const refreshesKey = 'refreshes'

    const refreshes = {
        cpu: 'every 1s',
        memory: 'every 2s',
        disk: 'load',
        gpu: 'every 1s'
    }

    Alpine.store(themeKey, {
        current: localStorage.getItem(themeKey) || 'dark',

        toggle() {
            let newTheme = 'light'
    
            if (this.current === 'light') {
                newTheme = 'dark'
            }

            this.current = newTheme
            localStorage.setItem(themeKey, newTheme)
        }
    })

    Alpine.store(refreshesKey, {
        values: (() => {
            try {
                const storedValues = localStorage.getItem(refreshesKey)

                if (storedValues) {
                    return JSON.parse(storedValues)
                }

                return refreshes
            }
            catch(err) {
                return refreshes
            }
        })(),
        getTrigger(refresh) {
            let trigger = `load, ${this.values[refresh]}`

            if (this.values[refresh] === 'load') {
                trigger = 'load'
            }

            return trigger
        },
        get(refresh) {
            const interval = this.values[refresh]

            if (interval === 'load') {
                return 0
            }

            return interval.replace('every', '').replace('s', '').trim()
        },
        set(refresh, value) {
            let newValue = `every ${value}s`

            if (value == '0') {
                newValue = 'load'
            }

            const newValues = {
                ...this.values,
                [refresh]: newValue
            }

            this.values = newValues
            localStorage.setItem(refreshesKey, JSON.stringify(newValues))

            // TODO: Find a way to HTMX updates without reload
            location.reload()
        }
    })
})