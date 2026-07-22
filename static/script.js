async function updateStats() {
    try {
        const response = await fetch('/stats')
        const data = await response.json()

        // Update CPU Items
        const cpuContainer = document.getElementById('cpu-container')
        cpuContainer.innerHTML = ''
    
        data.cpus.forEach((cpu) => {
            const cpuDiv = document.createElement('div')

            cpuDiv.className = 'cpu-item'
            cpuDiv.innerHTML = `C${Number(cpu.core) + 1}<br>${cpu.usage}%`
            cpuDiv.style.backgroundColor = cpu.usage > 80 ? '#451a1a' : '#334155'

            cpuContainer.appendChild(cpuDiv)
        })

        // Update RAM Elements
        const ramContainer = document.getElementById('ram-container')
        ramContainer.innerHTML = ''
    
        data.ram.forEach((ram) => {
            const ramDiv = document.createElement('div')
            ramDiv.className = 'memory-item'

            ramDiv.innerHTML = /* html */`
                <p id="ram-text">${ram.used_gb} / ${ram.total_gb} GB</p>
                <div class="progress-bar">
                    <div id="ram-fill" class="progress-fill" style="width: ${ram.percent}%"></div>
                </div>
                <small id="ram-percent" style="color:var(--accent)">${ram.percent}%</small>
            `
            ramContainer.appendChild(ramDiv)
        })
    
        // Update GPU Elements
        const gpuContainer = document.getElementById('gpu-container')
        gpuContainer.innerHTML = ''
    
        data.gpus.forEach((gpu) => {
            const gpuDiv = document.createElement('div')

            gpuDiv.className = 'gpu-card'
            gpuDiv.style.marginBottom = "15px"

            gpuDiv.innerHTML = /* html */`
                <div>
                    <strong style="display:block">${gpu.name}</strong>
                    <small>Load: ${gpu.load_percent}% | VRAM: ${Math.round(gpu.mem_used / 1024)}MB / ${Math.round(gpu.mem_total / 1024)}MB</small>
                </div>
                <div class="progress-bar_small">
                    <div class="progress-fill_small" style="width:${gpu.mem_percent}%"></div>
                </div>
            `
            gpuContainer.appendChild(gpuDiv)
        })

        // Update Disk Elements
        const diskList = document.getElementById('disk-list')
        diskList.innerHTML = ''

        data.disks.forEach((disk) => {
            const disksDiv = document.createElement('div')
            disksDiv.style.marginBottom = '10px'
            disksDiv.innerHTML = /* html */`
                <strong style="font-size: 0.9rem">${disk.mount}</strong> (${disk.used}GB / ${disk.total}GB)<br>
                <div class="progress-bar" style="height:8px width: ${disk.percent}%"></div>
                <small>${disk.percent}% - ${disk.device}</small>
            `

            diskList.appendChild(disksDiv)
        })

    }
    catch (err) { 
        console.log("Error fetching stats", err) 
    }
}

updateStats()
setInterval(updateStats, 1000)
