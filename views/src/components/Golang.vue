<template>
    <div>
        <img alt="Golang + Vue.js" src="statics/gopher_vue.png" height="200">
        <h1>Golang + Vue.js</h1>
        <div>
            <h2>Golang Frameworks</h2>
            <ul style="display: inline-block">
                <li v-for="framework in frameworks" :key="framework.ID">{{ framework.Name }}</li>
            </ul>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    data () {
        return {
            frameworks: []
        }
    },
    mounted () {
        let self = this
        axios
            .get('http://localhost:1323/frameworks')
                .then(response => {
                console.log(response.data)
                response.data.forEach(element => {
                    self.frameworks.push(element)
                })
                console.log(self.frameworks)
            })
            .catch(error => {
                console.log(error)   
            })
    }
}
</script>
