<template>
  <div>
    <b-form @submit="handleRegistration">
      <b-form-group label="Phone Number" label-for="phone-number">
        <b-form-input id="phone-number" v-model="phoneNumber" type="text"></b-form-input>
      </b-form-group>

      <b-form-group label="Password" label-for="password">
        <b-form-input id="password" v-model="password" type="password"></b-form-input>
      </b-form-group>

      <b-button type="submit" variant="primary">Register</b-button>
    </b-form>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'SignUpUser',
  data() {
    return {
      phoneNumber: '',
      password: ''
    }
  },
  methods: {
    handleRegistration() {
      axios.post('/register', {
        phoneNumber: this.phoneNumber,
        password: this.password
      })
      .then(response => {
        console.log(response)
      })
      .catch(error => {
        console.log(error)
      })
    }
  }
}
</script>