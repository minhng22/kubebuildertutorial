package v2

import (
	"fmt"
	_ "fmt"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"strings"
	"tutorial.kubebuilder.io/project/api/v1"
)

// ConvertTo converts this CronJob to the Hub version (v1).
func (src *CronJob) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.CronJob)

	schedule := src.Spec.Schedule
	scheduleParts := []string{"*", "*", "*", "*", "*"}
	if schedule.Minute != nil {
		scheduleParts[0] = string(*schedule.Minute)
	}
	if schedule.Hour != nil {
		scheduleParts[1] = string(*schedule.Hour)
	}
	if schedule.DayOfMonth != nil {
		scheduleParts[2] = string(*schedule.DayOfMonth)
	}
	if schedule.Month != nil {
		scheduleParts[3] = string(*schedule.Month)
	}
	if schedule.DayOfWeek != nil {
		scheduleParts[4] = string(*schedule.DayOfWeek)
	}
	dst.Spec.Schedule = strings.Join(scheduleParts, " ")
	// rote conversion

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.StartingDeadlineSeconds = src.Spec.StartingDeadlineSeconds
	dst.Spec.ConcurrencyPolicy = v1.ConcurrencyPolicy(src.Spec.ConcurrencyPolicy)
	dst.Spec.Suspend = src.Spec.Suspend
	dst.Spec.JobTemplate = src.Spec.JobTemplate
	dst.Spec.SuccessfulJobHistoryLimit = src.Spec.SuccessfulJobHistoryLimit
	dst.Spec.FailedJobsHistoryLimit = src.Spec.FailedJobsHistoryLimit

	// Status
	dst.Status.Active = src.Status.Active
	dst.Status.LastScheduleTime = src.Status.LastScheduleTime
	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *CronJob) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.CronJob)

	schedParts := strings.Split(src.Spec.Schedule, " ")
	if len(schedParts) != 5 {
		return fmt.Errorf("invalid schedule: not a standard 5-field schedule")
	}
	partIfNeeded := func(raw string) *CronField {
		if raw == "*" {
			return nil
		}
		part := CronField(raw)
		return &part
	}
	dst.Spec.Schedule.Minute = partIfNeeded(schedParts[0])
	dst.Spec.Schedule.Hour = partIfNeeded(schedParts[1])
	dst.Spec.Schedule.DayOfMonth = partIfNeeded(schedParts[2])
	dst.Spec.Schedule.Month = partIfNeeded(schedParts[3])
	dst.Spec.Schedule.DayOfWeek = partIfNeeded(schedParts[4])

	// rote conversion

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.StartingDeadlineSeconds = src.Spec.StartingDeadlineSeconds
	dst.Spec.ConcurrencyPolicy = ConcurrencyPolicy(src.Spec.ConcurrencyPolicy)
	dst.Spec.Suspend = src.Spec.Suspend
	dst.Spec.JobTemplate = src.Spec.JobTemplate
	dst.Spec.SuccessfulJobHistoryLimit = src.Spec.SuccessfulJobHistoryLimit
	dst.Spec.FailedJobsHistoryLimit = src.Spec.FailedJobsHistoryLimit

	// Status
	dst.Status.Active = src.Status.Active
	dst.Status.LastScheduleTime = src.Status.LastScheduleTime
	return nil
}
