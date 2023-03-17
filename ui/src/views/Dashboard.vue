<template>
    <v-container  fluid class="fill-height">
      <v-row no-gutters class="fill-height">
        <v-col>
        </v-col>
        <v-col>
        </v-col>
        <v-col order="last">
         <v-card elevation="1" class="fill-height">
           <v-card-title>Latest Results</v-card-title>
            <v-timeline side="end" align="start">
              <v-timeline-item
                dot-color="pink"
                size="small"
              >
                <div class="d-flex">
                  <strong class="me-3">3/14/2023 17:01</strong>
                  <div>
                    <strong>Test</strong>
                    <div class="text-caption">
                      Package github.com/benthosdev/benthos
                    </div>
                  </div>
                </div>
              </v-timeline-item>
          
              <v-timeline-item
                dot-color="teal-lighten-3"
                size="small"
              >
                <div class="d-flex">
                  <strong class="me-4">3/15/2023 11:34</strong>
                  <div>
                    <strong>Benchmark</strong>
                    <div class="text-caption mb-2">
                      Package github.com/benthosdev/benthos
                    </div>
                  </div>
                </div>
              </v-timeline-item>
  
              <v-timeline-item
                dot-color="teal-lighten-3"
                size="small"
              >
                <div class="d-flex">
                  <strong class="me-4">3/15/2023 11:42</strong>
                  <div>
                    <strong>Benchmark</strong>
                    <div class="text-caption mb-2">
                      Package github.com/benthosdev/benthos
                    </div>
                  </div>
                </div>
              </v-timeline-item>
  
              <v-timeline-item
                dot-color="teal-lighten-3"
                size="small"
              >
                <div class="d-flex">
                  <strong class="me-4">3/15/2023 11:46</strong>
                  <div>
                    <strong>Benchmark</strong>
                    <div class="text-caption mb-2">
                      Package github.com/benthosdev/benthos
                    </div>
                  </div>
                </div>
              </v-timeline-item>
          
              <v-timeline-item
                dot-color="pink"
                size="small"
              >
                <div class="d-flex">
                  <strong class="me-4">3/17/2023 19:21</strong>
                  <div>
                    <strong>Test</strong>
                    <div class="text-caption">
                      Package github.com/benthosdev/benthos
                    </div>
                  </div>
                </div>
              </v-timeline-item>
            </v-timeline>
	 </v-card>
        </v-col>
      </v-row>
    </v-container>
</template>

<script>
import { gql } from "apollo-boost";

const FETCH_BENCHMARKS = gql`
        query benches {
          benches {
            edges {
              node {
                id
                os
                arch
                cpu
                package
                benchResult {
                  name
                  allocsPerOp
                  allocedBytesPerOp
                  mbPerS
                  n
                  nsPerOp
                }
              }
            }
          }
        }`;
export default {
  data() {
    return {
      benches: [],
      headers: [
        {
	  title: "Package",
	  key: "node.package",
	},
        {
	  title: "OS",
	  key: "node.os",
	},
        {
	  title: "Arch",
	  key: "node.arch",
	},
        {
	  title: "CPU",
	  key: "node.cpu",
	},
      ]
    }
  },
  apollo: {
    benches: {
      query: FETCH_BENCHMARKS,
    }
  }
};
</script>
