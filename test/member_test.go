package init
import(
	"testing"
	"github.com/Jibx24/gittestfinal/entity"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
)
func TestPhoneNumber(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`Phone Number length is not 10 digits`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "11111111",
			Password: "1111111",
			Url: "https://www.linkedin.com/company/ilink/",
		}
		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone Number length is not 10 digits."))
	})
	t.Run("Valid Member", func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "1111111111",
			Password: "1111111",
			Url: "https://www.linkedin.com/company/ilink/",			
		}
		ok, err := govalidator.ValidateStruct(member)
        g.Expect(ok).To(BeTrue())
        g.Expect(err).To(BeNil())
	})
}