<template>
  <div>
    <b-form @submit="submitForm">
      <b-form-group label="Phone Number" label-for="phone-number">
        <b-form-input id="phone-number" v-model="phoneNumber" type="tel" required></b-form-input>
      </b-form-group>
      <b-form-group label="Email" label-for="email">
        <b-form-input id="email" v-model="email" type="email" required></b-form-input>
      </b-form-group>
      <b-button type="submit" variant="primary">Submit</b-button>
    </b-form>
    <b-alert v-if="errorMessage" :variant="'danger'" dismissible>{{ errorMessage }}</b-alert>
    <b-alert v-if="successMessage" :variant="'success'" dismissible>{{ successMessage }}</b-alert>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'RecoveryUser',
  data() {
    return {
      phoneNumber: '',
      email: '',
      errorMessage: '',
      successMessage: ''
    }
  },
  methods: {
    async submitForm() {
      try {
        await axios.post('/api/password-recovery', { phoneNumber: this.phoneNumber, email: this.email })
        this.successMessage = 'A password recovery link has been sent to your email'
        this.errorMessage = ''
      } catch (error) {
        this.errorMessage = 'An error occurred. Please try again later.'
        this.successMessage = ''
      }
    }
  }
}
</script>