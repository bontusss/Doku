<script setup lang="ts">
import { Button } from "@/components/ui/button";

defineProps<{ msg: string }>();
</script>

<template>
  <h1>{{ msg }}</h1>

  <div>
    <Button> Click me</Button>
  </div>
</template>

<style scoped></style>
