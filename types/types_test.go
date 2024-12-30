package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSample(t *testing.T) {
	now := time.Now()
	sample := NewSample("prefix", "metric", 123.45, map[string]string{"label1": "value1","label2": "value2"})
	t.Logf("sample: %+v\n", sample.ConvertTimeSeries("ms"))
	t.Logf("sample: %+v\n", sample)
	assert.Equal(t, "prefix_metric", sample.Metric)
	assert.Equal(t, 123.45, sample.Value)
	// assert.Equal(t, "value1", sample.Labels["label1"])
	assert.True(t, sample.Timestamp.IsZero())

	sample.SetTime(now)
	assert.Equal(t, now, sample.Timestamp)
}

func TestSample_ConvertTimeSeries(t *testing.T) {
	sample := NewSample("", "metric", 123.45, map[string]string{"label1": "value1"})
	ts := sample.ConvertTimeSeries("ms")

	assert.NotNil(t, ts)
	assert.Equal(t, "metric", ts.Labels[0].Value)
	assert.Equal(t, "label1", ts.Labels[1].Name)
	assert.Equal(t, "value1", ts.Labels[1].Value)
	assert.Equal(t, 123.45, ts.Samples[0].Value)
}

func TestSafeList(t *testing.T) {
	list := NewSafeList[int]()
	list.PushFront(1)
	list.PushFront(2)

	assert.Equal(t, 2, list.Len())

	val := list.PopBack()
	assert.Equal(t, 1, *val)
	assert.Equal(t, 1, list.Len())
}

func TestSampleList(t *testing.T) {
	list := NewSampleList()
	list.PushSample("", "metric1", 1.0)
	list.PushSample("", "metric2", 2.0)

	assert.Equal(t, 2, list.Len())

	samples := list.PopBackN(2)
	assert.Equal(t, 2, len(samples))
	assert.Equal(t, "metric1", samples[0].Metric)
	assert.Equal(t, "metric2", samples[1].Metric)
}

func TestSafeListLimited(t *testing.T) {
	list := NewSafeListLimited[int](2)
	list.PushFront(1)
	list.PushFront(2)
	success := list.PushFront(3) // should fail

	assert.Equal(t, 2, list.Len())
	assert.False(t, success)
}
