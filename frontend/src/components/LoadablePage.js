
export default {
  data() {
    return {
      loaded: false
    }
  },
  methods: {
    fetchData: async function(url) {
      this.loaded = false;
      const response = await fetch(url);
      const result = await response.json();
      this.loaded = true;
      return result;
    }
  }

}