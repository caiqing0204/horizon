package metrics

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

//nolint:lll
func TestObserve(t *testing.T) {
	wprBody := `{
    "pipelineRun":{
        "metadata":{
            "name":"test-music-docker-q58rp",
            "namespace":"tekton-resources",
            "labels":{
                "app.kubernetes.io/managed-by":"Helm",
                "cloudnative.music.netease.com/application":"testapp-1",
                "cloudnative.music.netease.com/cluster":"testcluster-1",
                "cloudnative.music.netease.com/environment":"env",
                "cloudnative.music.netease.com/region":"hz-test",
                "tekton.dev/pipeline":"default",
                "triggers.tekton.dev/eventlistener":"default-listener",
                "triggers.tekton.dev/trigger":"",
                "triggers.tekton.dev/triggers-eventid":"cttzw"
            }
        },
        "status":{
            "conditions":[
                {
                    "type":"Succeeded",
                    "status":"True",
                    "lastTransitionTime":"2021-06-24T06:38:18Z",
                    "reason":"Succeeded",
                    "message":"Tasks Completed: 2 (Failed: 0, Cancelled 0), Skipped: 0"
                }
            ],
            "startTime":"2021-06-24T06:36:11Z",
            "completionTime":"2021-06-24T06:38:18Z",
            "taskRuns":{
                "test-music-docker-q58rp-build-g8khd":{
                    "pipelineTaskName":"build",
                    "status":{
                        "conditions":[
                            {
                                "type":"Succeeded",
                                "status":"True",
                                "lastTransitionTime":"2021-06-24T06:36:43Z",
                                "reason":"Succeeded",
                                "message":"All Steps have completed executing"
                            }
                        ],
                        "podName":"test-music-docker-q58rp-build-g8khd-pod-mwsld",
                        "startTime":"2021-06-24T06:36:11Z",
                        "completionTime":"2021-06-24T06:36:43Z",
                        "steps":[
                            {
                                "terminated":{
                                    "exitCode":0,
                                    "reason":"Completed",
                                    "startedAt":"2021-06-24T06:36:18Z",
                                    "finishedAt":"2021-06-24T06:36:26Z",
                                    "containerID":"docker://3cccbd086c26e83e41fe8fcd86ef4e0f42e3963371c581e458df223b94da8d1e"
                                },
                                "name":"git",
                                "container":"step-git",
                                "imageID":"docker-pullable://harbor.mock.org/cloudnative/library/tekton-builder@sha256:14194e518981f5d893b85e170a28ba8aa80c2c610f63cfba814b6a460f48dc29"
                            },
                            {
                                "terminated":{
                                    "exitCode":0,
                                    "reason":"Completed",
                                    "startedAt":"2021-06-24T06:36:26Z",
                                    "finishedAt":"2021-06-24T06:36:34Z",
                                    "containerID":"docker://58d06c0a4bfa8212620e5a85a42e9af0768a4adb9ade2219dc75aee4d650ff23"
                                },
                                "name":"compile",
                                "container":"step-compile",
                                "imageID":"docker-pullable://harbor.mock.org/cloudnative/library/tekton-builder@sha256:14194e518981f5d893b85e170a28ba8aa80c2c610f63cfba814b6a460f48dc29"
                            },
                            {
                                "terminated":{
                                    "exitCode":0,
                                    "reason":"Completed",
                                    "message":"[{\"key\":\"properties\",\"value\":\"harbor.mock.org/ndp-gjq/test-music-docker:helloworld-b1f57848-20210624143634 git@github.com:demo/demo.git helloworld b1f578488e3123e97ec00b671db143fb8f0abecf\",\"type\":\"TaskRunResult\"}]",
                                    "startedAt":"2021-06-24T06:36:34Z",
                                    "finishedAt":"2021-06-24T06:36:42Z",
                                    "containerID":"docker://9189624ad3981fd738ec5bf286f1fc5b688d71128b9827820ebc2541b2801dae"
                                },
                                "name":"image",
                                "container":"step-image",
                                "imageID":"docker-pullable://harbor.mock.org/cloudnative/library/kaniko-executor@sha256:473d6dfb011c69f32192e668d86a47c0235791e7e857c870ad70c5e86ec07e8c"
                            }
                        ]
                    }
                },
                "test-music-docker-q58rp-deploy-xzjkg":{
                    "pipelineTaskName":"deploy",
                    "status":{
                        "conditions":[
                            {
                                "type":"Succeeded",
                                "status":"True",
                                "lastTransitionTime":"2021-06-24T06:38:18Z",
                                "reason":"Succeeded",
                                "message":"All Steps have completed executing"
                            }
                        ],
                        "podName":"test-music-docker-q58rp-deploy-xzjkg-pod-zkcc4",
                        "startTime":"2021-06-24T06:36:43Z",
                        "completionTime":"2021-06-24T06:38:18Z",
                        "steps":[
                            {
                                "terminated":{
                                    "exitCode":0,
                                    "reason":"Completed",
                                    "startedAt":"2021-06-24T06:36:48Z",
                                    "finishedAt":"2021-06-24T06:38:18Z",
                                    "containerID":"docker://fb2579fe83579e1918b5dcedc35f3686cad8ac632cc750d6d92f556341b5f7bb"
                                },
                                "name":"deploy",
                                "container":"step-deploy",
                                "imageID":"docker-pullable://harbor.mock.org/cloudnative/library/tekton-builder@sha256:14194e518981f5d893b85e170a28ba8aa80c2c610f63cfba814b6a460f48dc29"
                            }
                        ]
                    }
                }
            }
        }
    }
	}
	`

	prHistogramMetric := `
        # HELP horizon_pipelinerun_duration_seconds PipelineRun duration info
        # TYPE horizon_pipelinerun_duration_seconds histogram
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="0"} 0
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="30"} 0
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="60"} 0
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="90"} 0
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="120"} 0
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="150"} 1
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="180"} 1
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="210"} 1
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="240"} 1
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="270"} 1
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="300"} 1
        horizon_pipelinerun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",le="+Inf"} 1
        horizon_pipelinerun_duration_seconds_sum{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok"} 127
        horizon_pipelinerun_duration_seconds_count{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok"} 1
        `

	trHistogramMetric := `
        # HELP horizon_taskrun_duration_seconds Taskrun duration info
        # TYPE horizon_taskrun_duration_seconds histogram
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="0"} 0
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="15"} 0
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="30"} 0
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="45"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="60"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="75"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="90"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="105"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="120"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="135"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="150"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="165"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="180"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="195"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="210"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="225"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="240"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="255"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="270"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="285"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="300"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build",le="+Inf"} 1
        horizon_taskrun_duration_seconds_sum{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build"} 32
        horizon_taskrun_duration_seconds_count{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="build"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="0"} 0
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="15"} 0
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="30"} 0
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="45"} 0
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="60"} 0
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="75"} 0
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="90"} 0
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="105"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="120"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="135"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="150"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="165"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="180"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="195"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="210"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="225"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="240"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="255"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="270"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="285"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="300"} 1
        horizon_taskrun_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy",le="+Inf"} 1
        horizon_taskrun_duration_seconds_sum{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy"} 95
        horizon_taskrun_duration_seconds_count{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",task="deploy"} 1
        `

	stepHistogramMetric := `
        # HELP horizon_step_duration_seconds Step duration info
        # TYPE horizon_step_duration_seconds histogram
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="0"} 0
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="15"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="30"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="45"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="60"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="75"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="90"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="105"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="120"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="135"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="150"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="165"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="180"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="195"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="210"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="225"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="240"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="255"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="270"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="285"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="300"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build",le="+Inf"} 1
        horizon_step_duration_seconds_sum{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build"} 8
        horizon_step_duration_seconds_count{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="compile",task="build"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="0"} 0
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="15"} 0
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="30"} 0
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="45"} 0
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="60"} 0
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="75"} 0
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="90"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="105"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="120"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="135"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="150"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="165"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="180"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="195"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="210"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="225"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="240"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="255"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="270"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="285"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="300"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy",le="+Inf"} 1
        horizon_step_duration_seconds_sum{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy"} 90
        horizon_step_duration_seconds_count{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="deploy",task="deploy"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="0"} 0
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="15"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="30"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="45"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="60"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="75"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="90"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="105"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="120"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="135"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="150"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="165"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="180"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="195"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="210"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="225"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="240"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="255"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="270"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="285"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="300"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build",le="+Inf"} 1
        horizon_step_duration_seconds_sum{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build"} 8
        horizon_step_duration_seconds_count{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="git",task="build"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="0"} 0
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="15"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="30"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="45"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="60"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="75"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="90"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="105"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="120"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="135"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="150"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="165"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="180"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="195"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="210"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="225"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="240"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="255"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="270"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="285"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="300"} 1
        horizon_step_duration_seconds_bucket{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build",le="+Inf"} 1
        horizon_step_duration_seconds_sum{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build"} 8
        horizon_step_duration_seconds_count{application="ndp-gjq",cluster="test-music-docker",environment="test",pipeline="default",region="hz-test",result="ok",step="image",task="build"} 1
      `

	var wpr1 *WrappedPipelineRun
	_ = json.Unmarshal([]byte(wprBody), &wpr1)
	Observe(FormatPipelineResults(wpr1.PipelineRun))
	if err := testutil.CollectAndCompare(_prHistogram, strings.NewReader(prHistogramMetric)); err != nil {
		t.Fatalf("err: %v", err)
	}

	if err := testutil.CollectAndCompare(_trHistogram, strings.NewReader(trHistogramMetric)); err != nil {
		t.Fatalf("err: %v", err)
	}

	if err := testutil.CollectAndCompare(_stepHistogram, strings.NewReader(stepHistogramMetric)); err != nil {
		t.Fatalf("err: %v", err)
	}
}
