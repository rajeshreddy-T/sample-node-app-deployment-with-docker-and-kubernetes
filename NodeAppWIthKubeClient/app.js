const express = require('express');
const bodyParser = require('body-parser');
const k8s = require('@kubernetes/client-node');

const app = express();

// Set up the Kubernetes API client and connection pool
const kubeconfig = new k8s.KubeConfig();
kubeconfig.loadFromDefault();



const k8sApi = kubeconfig.makeApiClient(k8s.BatchV1Api);
// const k8sPool =  k8s.KubeConfigPool(kubeconfig);
const k8sPool = new k8s.KubeConfigPool(kubeconfig);
k8sPool.init();


// Parse JSON request bodies
app.use(bodyParser.json());

// Handle incoming POST requests to launch a container
app.post('/launch', async (req, res) => {
  try {
    // Extract the container image and command from the request body
    const { image, command } = req.body;

    // Define the Job specification
    const jobSpec = {
      apiVersion: 'batch/v1',
      kind: 'Job',
      metadata: {
        name: `my-job-${Date.now()}`,
      },
      spec: {
        template: {
          spec: {
            containers: [{
              name: 'my-container',
              image,
              command,
            }],
            restartPolicy: 'Never',
          },
        },
      },
    };

    // Use a connection from the connection pool to create the Job
    const k8sApiClient = await k8sPool.acquire();
    const response = await k8sApiClient.createNamespacedJob('default', jobSpec);
    const jobName = response.body.metadata.name;
    console.log(`Job ${jobName} created`);

    // Watch the status of the Job
    const jobWatcher = k8sApi.watchNamespacedJob('default', jobName);
    jobWatcher.on('data', async (event) => {
      const jobStatus = event.object.status;
      console.log(`Job ${jobName} status: ${jobStatus.conditions[0].type}`);

      // If the Job has completed, return the pod name and status to the client
      if (jobStatus.conditions[0].type === 'Complete') {
        const podName = jobStatus.podName;
        const podWatcher = k8sApi.watchNamespacedPod('default', podName);
        podWatcher.on('data', (event) => {
          const podStatus = event.object.status;
          console.log(`Pod ${podName} status: ${podStatus.phase}`);

          if (podStatus.phase === 'Succeeded') {
            res.json({ podName, status: podStatus.phase });
            podWatcher.abort();
            jobWatcher.abort();
            k8sPool.release(k8sApiClient);
          }
          else if (podStatus.phase === 'Failed') {
            res.status(500).json({ error: 'Container failed to start' });
            podWatcher.abort();
            jobWatcher.abort();
            k8sPool.release(k8sApiClient);
          }
        });
      }
    });
  }
  catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Failed to launch container' });
  }
});

// Start the web server
app.listen(3000, () => {
  console.log('Web server listening on port 3000');
});

// docker tag local-image:tagname new-repo:tagname
// docker push new-repo:tagname
// 8341265379@kubernetes